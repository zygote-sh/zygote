package zygote

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config interface {
	Set(key string, val any)
	IsSet(key string) bool
	GetString(key string) (string, error)
	GetBool(key string) (bool, error)
	GetBoolPtr(key string) (*bool, error)
	GetInt(key string) (int, error)
	GetIntPtr(key string) (*int, error)
	GetStringSlice(key string) ([]string, error)
	GetStringSliceIsFlagSet(key string) ([]string, bool, error)
	GetStringMapString(key string) (map[string]string, error)
	GetDuration(key string) (time.Duration, error)
}

// LiveConfig is an implementation of Config for live values.
type LiveConfig struct {
	cliArgs map[string]bool
}

var _ Config = &LiveConfig{}

func (c *LiveConfig) Set(key string, val any) {
	viper.Set(key, val)
}

func (c *LiveConfig) IsSet(key string) bool {
	matches := regexp.MustCompile("\b*--([a-z-_]+)").
		FindAllStringSubmatch(strings.Join(os.Args, " "), -1)
	if len(matches) == 0 {
		return false
	}

	if len(c.cliArgs) == 0 {
		args := make(map[string]bool)
		for _, match := range matches {
			args[match[1]] = true
		}
		c.cliArgs = args
	}
	return c.cliArgs[key]
}

// GetString returns a config value as a string.
func (c *LiveConfig) GetString(key string) (string, error) {
	str := viper.GetString(key)

	if isRequired(key) && strings.TrimSpace(str) == "" {
		return "", NewMissingArgsErr(key)
	}
	return str, nil
}

// GetBool returns a config value as a bool.
func (c *LiveConfig) GetBool(key string) (bool, error) {
	return viper.GetBool(key), nil
}

// GetBoolPtr returns a config value as a bool pointer.
func (c *LiveConfig) GetBoolPtr(key string) (*bool, error) {
	if !c.IsSet(key) {
		return nil, nil
	}
	val := viper.GetBool(key)
	return &val, nil
}

// GetInt returns a config value as an int.
func (c *LiveConfig) GetInt(key string) (int, error) {
	val := viper.GetInt(key)

	if isRequired(key) && val == 0 {
		return 0, NewMissingArgsErr(key)
	}
	return val, nil
}

// GetIntPtr returns a config value as an int pointer.
func (c *LiveConfig) GetIntPtr(key string) (*int, error) {

	if !c.IsSet(key) {
		if isRequired(key) {
			return nil, NewMissingArgsErr(key)
		}
		return nil, nil
	}
	val := viper.GetInt(key)
	return &val, nil
}

// GetStringSlice returns a config value as a string slice.
func (c *LiveConfig) GetStringSlice(key string) ([]string, error) {
	val := viper.GetStringSlice(key)

	if isRequired(key) && emptyStringSlice(val) {
		return nil, NewMissingArgsErr(key)
	}

	out := []string{}
	for _, item := range viper.GetStringSlice(key) {
		item = strings.TrimPrefix(item, "[")
		item = strings.TrimSuffix(item, "]")

		list := strings.Split(item, ",")
		for _, str := range list {
			if str == "" {
				continue
			}
			out = append(out, str)
		}
	}
	return out, nil
}

// GetStringSliceIsFlagSet returns a config value as a string slice and a bool representing the existence of the flag.
func (c *LiveConfig) GetStringSliceIsFlagSet(key string) ([]string, bool, error) {
	if !c.IsSet(key) {
		return nil, false, nil
	}
	strSlice, err := c.GetStringSlice(key)
	return strSlice, true, err
}

// GetStringMapString returns a config value as a string to string map.
func (c *LiveConfig) GetStringMapString(key string) (map[string]string, error) {
	if isRequired(key) && !c.IsSet(key) {
		return nil, NewMissingArgsErr(key)
	}

	// We cannot call viper.GetStringMapString because it does not handle
	// pflag's StringToStringP properly:
	// https://github.com/spf13/viper/issues/608
	// Re-implement the necessary pieces on our own instead.

	vals := map[string]string{}
	items := viper.GetStringSlice(key)
	for _, item := range items {
		parts := strings.SplitN(item, "=", 2)
		if len(parts) < 2 {
			return nil, fmt.Errorf("item %q does not adhere to form: key=value", item)
		}
		labelKey := parts[0]
		labelValue := parts[1]
		vals[labelKey] = labelValue
	}

	return vals, nil
}

// GetDuration returns a config value as a duration.
func (c *LiveConfig) GetDuration(key string) (time.Duration, error) {
	return viper.GetDuration(key), nil
}

func isRequired(key string) bool {
	return viper.GetBool(fmt.Sprintf("required.%s", key))
}

// This is needed because an empty StringSlice flag returns `"[]"`
func emptyStringSlice(s []string) bool {
	return (len(s) == 1 && s[0] == "[]") || len(s) == 0
}

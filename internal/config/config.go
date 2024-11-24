package config

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var v = viper.New()

type ConfigOption struct {
	Field        interface{}
	FlagName     string
	EnvVar       string
	DefaultValue string
	Description  string
}

// Interface for configuration structs that can be initialized
type ConfigInitializer interface {
	Init(cmd *cobra.Command)
}

// Global registry of all ConfigInitializers
var configInitializers []ConfigInitializer

// Register a config initializer
func RegisterConfig(initializer ConfigInitializer) {
	configInitializers = append(configInitializers, initializer)
}

// Initialize all registered configurations
func InitAllConfigs(cmd *cobra.Command) {
	for _, initializer := range configInitializers {
		initializer.Init(cmd)
	}
}

func InitializeConfig(cmd *cobra.Command, configStruct interface{}, options []ConfigOption) {
	if err := v.Unmarshal(configStruct); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	for _, opt := range options {
		switch field := opt.Field.(type) {
		case *string:
			cmd.PersistentFlags().StringVar(field, opt.FlagName, opt.DefaultValue, opt.Description)
		case *bool:
			defaultValue := opt.DefaultValue == "true"
			cmd.PersistentFlags().BoolVar(field, opt.FlagName, defaultValue, opt.Description)
		default:
			log.Fatalf("Unsupported field type for flag: %s", opt.FlagName)
		}
		v.BindPFlag(opt.EnvVar, cmd.PersistentFlags().Lookup(opt.FlagName))
	}
}

func InitConfig() {
	v.SetDefault("genome-file", "${ZDOTDIR}/zygote.genome")
	v.SetDefault("static-file", "${ZDOTDIR}/genome.zsh")
	v.SetDefault("cache-dir", "${HOME}/.cache/zygote")
	v.SetDefault("database-name", "zygote.bdb")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
}

func SetConfigFile(path string) {
	v.SetConfigFile(path)
}

func WriteConfig() {
	if err := v.WriteConfig(); err != nil {
		log.Fatalf("Error writing config file, %s", err)
	}
}

func ReadConfig() {
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func init() {
	v.SetEnvPrefix("ZYGOTE")
}

func GetString(configName string) string {
	return os.ExpandEnv(v.GetString(configName))
}

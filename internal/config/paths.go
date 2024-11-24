package config

import "github.com/spf13/cobra"

type PathsConfig struct {
	PluginDir string
	ConfigDir string
	CacheDir  string
}

var Paths = PathsConfig{}

var pathOptions = []ConfigOption{
	{Field: &Paths.PluginDir,
		FlagName:     "plugin-dir",
		EnvVar:       "PLUGIN_DIR",
		DefaultValue: "",
		Description:  "Directory where plugins are stored"},

	{Field: &Paths.ConfigDir,
		FlagName:     "config-dir",
		EnvVar:       "CONFIG_DIR",
		DefaultValue: "",
		Description:  "Directory where the configuration files are stored"},

	{Field: &Paths.CacheDir,
		FlagName:     "cache-dir",
		EnvVar:       "CACHE_DIR",
		DefaultValue: "${HOME}/.cache/zygote",
		Description:  "Directory where the cache files are stored"},
}

// Automatically register the Paths config
func init() {
	RegisterConfig(&Paths)
}

func (p *PathsConfig) Init(cmd *cobra.Command) {
	InitializeConfig(cmd, &Paths, pathOptions)
}

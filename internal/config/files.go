package config

import "github.com/spf13/cobra"

type FilesConfig struct {
	GenomeFile   string
	DatabaseName string
}

var Files = FilesConfig{}

var fileOptions = []ConfigOption{
	{Field: &Files.GenomeFile,
		FlagName:     "genome-file",
		EnvVar:       "GENOME_FILE",
		DefaultValue: "${ZDOTDIR}/genome",
		Description:  "Genome file to load",
	},
	{Field: &Files.DatabaseName,
		FlagName:     "database-name",
		EnvVar:       "DATABASE_NAME",
		DefaultValue: "zygote.bdb",
		Description:  "Name of the database to use",
	},
}

// Automatically register the Paths config
func init() {
	RegisterConfig(&Files)
}

func (p *FilesConfig) Init(cmd *cobra.Command) {
	InitializeConfig(cmd, &Files, fileOptions)
}

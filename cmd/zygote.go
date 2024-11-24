package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zygote-sh/zygote/internal/config"
	"github.com/zygote-sh/zygote/internal/db"
	"github.com/zygote-sh/zygote/internal/filesystem"
)

// var rootCmd = &cobra.Command{
// 	Use:   "zygote",
// 	Short: "Zygote Plugin Manager",
// 	Long:  ``,
// 	PersistentPreRun: func(cmd *cobra.Command, args []string) {
// 		preRun()
// 	},
// 	PersistentPostRun: func(cmd *cobra.Command, args []string) {
// 		db.Close()
// 	},
// 	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
// 		db.Close()
// 		return nil
// 	},
// 	Run: func(cmd *cobra.Command, args []string) {
// 		cmd.Help()
// 	},
// }

var rootCmd = &Cmd{
	Command: &cobra.Command{
		Use:   "zygote",
		Short: "Zygote Plugin Manager",
		Long:  ``,
	},
}

func preRun() {
	config.InitConfig()
	filesystem.CreateDir(config.GetString("cache-dir"))
	db.Init(config.GetString("cache-dir") + "/" + config.GetString("database-name"))
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// config.InitAllConfigs(rootCmd)

	rootCmd.AddCommand(CmdPrint())
}

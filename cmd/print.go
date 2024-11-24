package cmd

// Remove this, this is a test command

import (
	"github.com/spf13/cobra"
)

func CmdPrint() *Cmd {
	cc := &cobra.Command{
		Use:   "print",
		Short: "Debug",
		Long:  ``,
	}

	c := &Cmd{Command: cc}

	return c
}

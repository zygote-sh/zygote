package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// Wrapper around cobra.Command to add child commands.
type Cmd struct {
	*cobra.Command
	*cobra.Group
	childCommands []*Cmd
}

// AddCommand adds child commands and adds child commands for cobra as well.
func (c *Cmd) AddCommand(commands ...*Cmd) {
	c.childCommands = append(c.childCommands, commands...)
	for _, cmd := range commands {
		c.Command.AddCommand(cmd.Command)
	}
}

// ChildCommands returns the child commands.
func (c *Cmd) ChildCommands() []*Cmd {
	return c.childCommands
}

type ValidArgsFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

// AddValidArgsFunc sets the function to run for dynamic completions
// ValidArgsFunc and ValidArgs are mutually exclusive. This function will
// return an error if ValidArgs is already set.
func (c *Cmd) AddValidArgsFunc(fn ValidArgsFunc) error {
	if len(c.Command.ValidArgs) == 0 {
		c.Command.ValidArgsFunction = fn
		return nil
	}
	return errors.New("unable to add ValidArgsFunction when ValidArgs is already set")
}

// CmdBuilder builds a new command.
func CmdBuilder(
	parent *Cmd,
	cliText, shortdesc, longdesc string,
	options ...cmdOption,
) *Cmd {
	cc := &cobra.Command{
		Use:   cliText,
		Short: shortdesc,
		Long:  longdesc,
	}

	c := &Cmd{Command: cc}

	if parent != nil {
		parent.AddCommand(c)
	}

	for _, co := range options {
		co(c)
	}
	return c
}

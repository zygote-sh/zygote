package cmd

// cmdOption allow configuration of a command.
type cmdOption func(*Cmd)

// aliasOpt adds aliases for a command.
func aliasOpt(aliases ...string) cmdOption {
	return func(c *Cmd) {
		if c.Aliases == nil {
			c.Aliases = []string{}
		}

		c.Aliases = append(c.Aliases, aliases...)
	}
}

// hiddenCmd make a command hidden.
func hiddenCmd() cmdOption {
	return func(c *Cmd) {
		c.Hidden = true
	}
}

package cmd

import (
	"github.com/zygote-sh/zygote/internal/zygote"
)

type CmdConfig struct {
	Zygote zygote.Config
	Args   []string
}

func NewCmdConfig(zc zygote.Config, args []string) *CmdConfig {
	return &CmdConfig{
		Zygote: zc,
		Args:   args,
	}
}

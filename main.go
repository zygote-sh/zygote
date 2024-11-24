package main

//go:generate go run . completion bash -o completions/zygote-completion.bash
//go:generate go run . completion fish -o completions/zygote.fish
//go:generate go run . completion powershell -o completions/zygote.ps1
//go:generate go run . completion zsh -o completions/zygote.zsh

import (
	"log"

	"github.com/zygote-sh/zygote/cmd"
)

func main() {
	log.SetPrefix("zygote: ")
	cmd.Execute()
}

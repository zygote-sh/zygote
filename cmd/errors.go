package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var (
	errOperationAborted = fmt.Errorf("Operation aborted.")

	colorErr    = color.RedString("Error")
	colorWarn   = color.YellowString("Warning")
	colorNotice = color.GreenString("Notice")

	// errAction specifies what should happen when an error occurs
	errAction = func() {
		os.Exit(1)
	}

	// ErrExitSilently instructs zygote to exit silently with a bad status code. This can be used to fail a command
	// without printing an error message to the screen.
	//
	// IMPORTANT! Make sure to print your own error message if you use this! It is important for users to know
	// what caused the failure.
	ErrExitSilently = fmt.Errorf("")
)

func checkErr(err error) {
	if err == nil {
		return
	}

	if errors.Is(err, ErrExitSilently) {
		errAction()
		return
	}

	output := viper.GetString("output")

	switch output {
	default:
		fmt.Fprintf(color.Output, "%s: %v\n", colorErr, err)
	case "json":
		es := outputErrors{
			Errors: []outputError{
				{Detail: err.Error()},
			},
		}

		b, _ := json.Marshal(&es)
		fmt.Println(string(b))
	}

	errAction()
}

type outputErrors struct {
	Errors []outputError `json:"errors"`
}

type outputError struct {
	Detail string `json:"detail"`
}

package zygote

import "fmt"

// MissingArgsErr is returned when there are too few arguments for a command.
type MissingArgsErr struct {
	Command string
}

var _ error = &MissingArgsErr{}

// NewMissingArgsErr creates a MissingArgsErr instance.
func NewMissingArgsErr(cmd string) *MissingArgsErr {
	return &MissingArgsErr{Command: cmd}
}

func (e *MissingArgsErr) Error() string {
	return fmt.Sprintf("(%s) command is missing required arguments", e.Command)
}

// TooManyArgsErr is returned when there are too many arguments for a command.
type TooManyArgsErr struct {
	Command string
}

var _ error = &TooManyArgsErr{}

// NewTooManyArgsErr creates a TooManyArgsErr instance.
func NewTooManyArgsErr(cmd string) *TooManyArgsErr {
	return &TooManyArgsErr{Command: cmd}
}

func (e *TooManyArgsErr) Error() string {
	return fmt.Sprintf("(%s) command contains unsupported arguments", e.Command)
}

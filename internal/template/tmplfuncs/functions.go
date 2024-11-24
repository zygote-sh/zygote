package tmplfuncs

import "os/exec"

var vars = map[string]interface{}{}

func (or *TmplFuncsRegistry) hasCommand(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}

func (or *TmplFuncsRegistry) setVar(key string, value interface{}) any {
	vars[key] = value
	return nil
}

func (or *TmplFuncsRegistry) getVar(key string) interface{} {
	return vars[key]
}

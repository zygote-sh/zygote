package context

import "runtime"

var Current *Runtime

type Runtime struct {
	Shell string
	OS    string
	Home  string
	Arch  string
	Path  *Path
	Fpath *Fpath
}

func Init(shell string) {
	Current = &Runtime{
		Shell: shell,
		OS:    runtime.GOOS,
		Arch:  runtime.GOARCH,
		Home:  Home(),
		Path:  getPath(),
		Fpath: getFpath(),
	}
}

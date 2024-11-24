package zygote

// !! Primary entry point lives in /cmd/zygote/main.go !!

import (
	"os"
	"strconv"
)

const (
	LatestReleaseURL = "https://api.github.com/repos/zygote-sh/zygote/releases/latest"
)

func init() {
	if Build != "" {
		ZygoteVersion.Build = Build
	}
	if Major != "" {
		i, _ := strconv.Atoi(Major)
		ZygoteVersion.Major = i
	}
	if Minor != "" {
		i, _ := strconv.Atoi(Minor)
		ZygoteVersion.Minor = i
	}
	if Patch != "" {
		i, _ := strconv.Atoi(Patch)
		ZygoteVersion.Patch = i
	}
	if Label == "" {
		ZygoteVersion.Label = "dev"
	} else {
		ZygoteVersion.Label = Label
	}
}

// CommandName returns the name by which zygote was invoked
func CommandName() string {
	name, ok := os.LookupEnv("SNAP_NAME")
	if !ok || name != "zygote" {
		return os.Args[0]
	}
	return name
}

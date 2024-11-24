package zygote

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/blang/semver"
)

var (
	// Set at build time. It defines the git SHA for the current build.
	Build string

	// Set at build time. Major semantic version of zygote.
	Major string

	// Set at build time. Minor semantic version of zygote.
	Minor string

	// Set at build time. Patch semantic version of zygote.
	Patch string

	// Set at build time. It defines the string that comes after the
	// version, ie, the "dev" in v1.0.0-dev.
	Label string

	// Full version of zygote.
	ZygoteVersion Version
)

type Version struct {
	Major, Minor, Patch int
	Name, Build, Label  string
}

func (v Version) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch))
	if v.Label != "" {
		buffer.WriteString("-" + v.Label)
	}

	return buffer.String()
}

// Returns the whole version string with the build hash.
func (v Version) Complete(lv LatestVersioner) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("doctl version %s", v.String()))

	if v.Build != "" {
		buffer.WriteString(fmt.Sprintf("\nGit commit hash: %s", v.Build))
	}

	if tagName, err := lv.LatestVersion(); err == nil {
		v0, err1 := semver.Make(tagName)
		v1, err2 := semver.Make(v.String())

		if len(v0.Build) == 0 {
			v0, err1 = semver.Make(tagName + "-release")
		}

		if err1 == nil && err2 == nil && v0.GT(v1) {
			buffer.WriteString(fmt.Sprintf("\nrelease %s is available, check it out! ", tagName))
		}
	}

	return buffer.String()
}

// Returns the version in JSON format.
func (v Version) CompleteJSON(lv LatestVersioner) string {
	versionInfo := &struct {
		Version       string `json:"version,omitempty"`
		Commit        string `json:"commit,omitempty"`
		LatestRelease string `json:"latestRelease"`
		Notification  string `json:"notification,omitempty"`
	}{
		Version: v.String(),
		Commit:  v.Build,
	}

	if tagName, err := lv.LatestVersion(); err == nil {
		versionInfo.LatestRelease = tagName

		v0, err1 := semver.Make(tagName)
		v1, err2 := semver.Make(v.String())

		if len(v0.Build) == 0 {
			v0, err1 = semver.Make(tagName + "-release")
		}

		if err1 == nil && err2 == nil && v0.GT(v1) {
			versionInfo.Notification = fmt.Sprintf(
				"release %s is available, check it out!",
				tagName,
			)
		}
	}

	data, _ := json.MarshalIndent(versionInfo, "", "  ")
	return string(data)
}

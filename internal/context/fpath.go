package context

import (
	"os"
	"strings"
)

type Fpath Path

func getFpath() *Fpath {
	if Current != nil {
		return Current.Fpath
	}

	fpath := &Fpath{}
	paths := os.Getenv("fpath")

	for _, p := range strings.Split(paths, pathDelimiter()) {
		fpath.Append(cleanPath(p))
	}

	return fpath
}

func (f *Fpath) Append(path string) {
	f.toPath().Append(path)
}

func (f *Fpath) toPath() *Path {
	return (*Path)(f)
}

func (f *Fpath) Contains(path string) bool {
	return f.toPath().Contains(path)
}

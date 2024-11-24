package path

import (
	"io/fs"
	"os"
	"path/filepath"
)

type Path string

func (p Path) IsExecutable() bool {
	pathInfo, err := os.Stat(string(p))
	if err != nil {
		return false
	}
	// Check if path is a file
	if pathInfo.IsDir() {
		return false
	}
	return pathInfo.Mode().Perm()&0111 != 0
}

func (p Path) IsDir() bool {
	pathInfo, err := os.Stat(string(p))
	if err != nil {
		return false
	}
	return pathInfo.IsDir()
}

func (p Path) Exists() bool {
	_, err := os.Stat(string(p))
	return err == nil
}

func (p Path) IsFile() bool {
	pathInfo, err := os.Stat(string(p))
	if err != nil {
		return false
	}
	return !pathInfo.IsDir()
}

func (path Path) Chmod(mode os.FileMode) error {
	return os.Chmod(string(path), mode)
}

func (path Path) Mkdir(perm os.FileMode) error {
	return os.Mkdir(string(path), perm)
}

func (path Path) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(string(path), perm)
}

func (path Path) Open() (*os.File, error) {
	return os.Open(string(path))
}

func (path Path) OpenFile(flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(string(path), flag, perm)
}

// Returns the absolute representation of path
func (path Path) Abs() (Path, error) {
	res, err := filepath.Abs(string(path))
	return Path(res), err
}

// Checks if path is absolute
func (path Path) IsAbs() bool {
	return filepath.IsAbs(string(path))
}

func (path Path) Base() Path {
	return Path(filepath.Base(string(path)))
}

func (path Path) Clean() Path {
	return Path(filepath.Clean(string(path)))
}

func (path Path) Dir() Path {
	return Path(filepath.Dir(string(path)))
}

func (path Path) Join(elem ...Path) Path {
	var e1 []string
	e1 = append(e1, string(path))
	for _, e := range elem {
		e1 = append(e1, string(e))
	}
	return Path(filepath.Join(e1...))
}

func (path Path) ValidPath() bool {
	return fs.ValidPath(string(path))
}

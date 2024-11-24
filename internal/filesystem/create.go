package filesystem

import "os"

func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

package db

import "time"

func Close() {
	db.Close()
}

func Set(key, value []byte) error {
	return set(key, value)
}

func Get(key []byte) ([]byte, error) {
	return get(key)
}

func Iterate(fn func(key, value []byte) error) error {
	return Iterate(fn)
}

func Count() (int, error) {
	return count()
}

func Delete(key []byte) error {
	return delete(key)
}

func LastReadOfFile(filePath string) (time.Time, error) {
	return lastReadOfFile(db, filePath)
}

func Sync() error {
	return sync()
}

func Compact() error {
	return compact()
}

func DropAll() error {
	return dropAll()
}

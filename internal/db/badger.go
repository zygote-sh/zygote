package db

import (
	"fmt"
	"os"
	"time"

	"github.com/dgraph-io/badger/v4"
)

func get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return valCopy, err
}

func set(key, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func delete(key []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func iterate(fn func(key, value []byte) error) error {
	return db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			v, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			if err := fn(k, v); err != nil {
				return err
			}
		}
		return nil
	})
}

func count() (int, error) {
	var count int
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			count++
		}
		return nil
	})
	return count, err
}

func dropAll() error {
	return db.DropAll()
}

func sync() error {
	return db.Sync()
}

func compact() error {
	return db.Flatten(3)
}

func lastReadOfFile(db *badger.DB, filePath string) (time.Time, error) {
	var lastReadTime time.Time

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(filePath))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				// Return zero time if no entry found
				return nil
			}
			return fmt.Errorf("failed to get last read time: %w", err)
		}

		err = item.Value(func(val []byte) error {
			lastReadTime, err = time.Parse(time.RFC3339, string(val))
			return err
		})
		if err != nil {
			return fmt.Errorf("failed to parse last read time: %w", err)
		}

		return nil
	})

	return lastReadTime, err
}

func compareFileTimes(db *badger.DB, filePath string) error {
	lastReadTime, err := lastReadOfFile(db, filePath)
	if err != nil {
		return fmt.Errorf("error retrieving last read time: %w", err)
	}

	// Get the file info
	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to get file info for %s: %w", filePath, err)
	}

	lastModifiedTime := info.ModTime()

	// Log comparison
	fmt.Printf(
		"File: %s\nLast Read Time: %v\nLast Modified Time: %v\n",
		filePath,
		lastReadTime,
		lastModifiedTime,
	)

	if lastReadTime.IsZero() {
		fmt.Println("File has never been read before.")
	} else if lastReadTime.After(lastModifiedTime) {
		fmt.Println("File has been read since it was last modified.")
	} else {
		fmt.Println("File has not been read since it was last modified.")
	}

	return nil
}

func logFileRead(db *badger.DB, filePath string) error {
	return db.Update(func(txn *badger.Txn) error {
		// Get the current time
		now := time.Now()

		// Store the time as bytes
		err := txn.Set([]byte(filePath), []byte(now.Format(time.RFC3339)))
		if err != nil {
			return fmt.Errorf("failed to set last read time: %w", err)
		}

		return nil
	})
}

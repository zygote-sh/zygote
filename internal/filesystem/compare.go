package filesystem

import (
	"fmt"
	"os"
)

func IsNewer(file1, file2 string) (bool, error) {
	info1, err := os.Stat(file1)
	if err != nil {
		return false, fmt.Errorf("failed to get file info for %s: %w", file1, err)
	}

	info2, err := os.Stat(file2)
	if err != nil {
		return false, fmt.Errorf("failed to get file info for %s: %w", file2, err)
	}

	// Compare the modification times
	return info1.ModTime().After(info2.ModTime()), nil
}

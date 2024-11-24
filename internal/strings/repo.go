package strings

import "strings"

func CheckRepo(repo string) bool {
	// TODO: Add more checks
	// Check if string contains a forward slash
	if strings.Contains(repo, "/") {
		return true
	}
	return false
}

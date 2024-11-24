package git

import "net/url"

type Repo url.URL

func Parse(s string) (Repo, error) {
	u, err := url.Parse(s)
	return Repo(*u), err
}

func (u *Repo) IsGitRepo() bool {
	if u.Scheme == "git" {
		return true
	}
	if u.Scheme == "https" || u.Scheme == "http" {
		return true
	}
	// Check if the scheme is in the host map
	if _, ok := hostMap[u.Scheme]; ok {
		return true
	}

	return false
}

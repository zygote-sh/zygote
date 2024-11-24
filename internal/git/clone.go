package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func clone(dir, repo, branch string) error {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           repo,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  true,
	})
	return err
}

func Clone(dir, repo, branch string) error {
	return clone(dir, repo, branch)
}

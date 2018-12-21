package git

import (
	"fmt"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Release Git release
func Release(dir string, version string, user *User) error {
	_, err := Tag(dir, version, user)
	if err != nil {
		return fmt.Errorf("[Release] tag: %v", err)
	}

	err = Push(dir, version)
	if err != nil {
		return fmt.Errorf("[Release] push: %v", err)
	}

	return nil
}

// Commit commit file
func Commit(dir string, file string, message string, user *User) error {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf("[Git] open repo: %v", err)
	}

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("[Git] worktree: %v", err)
	}

	_, err = w.Add(file)
	if err != nil {
		return fmt.Errorf("[Git] worktree add (%s): %v", file, err)
	}

	_, err = w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  user.Name,
			Email: user.Email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("[Git] commit: %v", err)
	}

	return nil
}

// Tag tag last commit
func Tag(dir string, version string, user *User) (*plumbing.Reference, error) {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return nil, fmt.Errorf("[Tag] open repo: %v", err)
	}

	head, err := r.Head()
	if err != nil {
		return nil, fmt.Errorf("[Tag] head: %v", err)
	}

	ref, err := r.CreateTag(version, head.Hash(), &git.CreateTagOptions{
		Tagger: &object.Signature{
			Name:  user.Name,
			Email: user.Email,
			When:  time.Now(),
		},
		Message: version,
	})
	if err != nil {
		return nil, fmt.Errorf("[Tag] create git tag: %v", err)
	}

	return ref, nil
}

// Push push to remote
func Push(dir string, version string) error {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf("[Push] open repo: %v", err)
	}

	// push using default options
	tagRef := fmt.Sprintf("refs/tags/%s:refs/tags/%s", version, version)
	err = r.Push(&git.PushOptions{
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/heads/master:refs/heads/master"),
			config.RefSpec(tagRef),
		},
	})
	if err != nil {
		switch err {
		case git.ErrRemoteNotFound:
			return nil
		default:
			return fmt.Errorf("[Push] push: %v", err)
		}
	}

	return nil
}

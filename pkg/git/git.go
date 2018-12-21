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
func Release(dir string, version string) error {
	_, err := Tag(dir, version)
	if err != nil {
		return fmt.Errorf("[Release] tag: %v", err)
	}

	err = Push(dir)
	if err != nil {
		return fmt.Errorf("[Release] push: %v", err)
	}

	return nil
}

// Tag tag last commit
func Tag(dir string, version string) (*plumbing.Reference, error) {
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
			Name:  "Release",
			Email: "release",
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
func Push(dir string) error {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf("[Push] open repo: %v", err)
	}

	// push using default options
	err = r.Push(&git.PushOptions{
		RefSpecs: []config.RefSpec{
			config.RefSpec("refs/heads/master:refs/heads/master"),
			config.RefSpec("refs/tags/*:refs/tags/*"),
		},
	})
	if err != nil {
		return fmt.Errorf("[Push] push: %v", err)
	}

	return nil
}

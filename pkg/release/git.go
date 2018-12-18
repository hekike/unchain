package release

import (
	"fmt"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GitTag tag last commit
func GitTag(dir string, version string) (*plumbing.Reference, error) {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return nil, fmt.Errorf("[GitTag] open repo: %v", err)
	}

	head, err := r.Head()
	if err != nil {
		return nil, fmt.Errorf("[GitTag] head: %v", err)
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
		return nil, fmt.Errorf("[GitTag] create git tag: %v", err)
	}

	return ref, nil
}

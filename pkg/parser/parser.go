package parser

import (
	"regexp"
	"strings"

	"github.com/hekike/conventional-commits/pkg/model"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var pattern = regexp.MustCompile(`^(?:(\w+)\(?(\w+)\)?: (.+))(?:(?:\r?\n|$){0,2}(.+))?(?:(?:\r?\n|$){0,2}(.+))?(?:\r?\n|$){0,2}`)
var breakingChange = "BREAKING CHANGE"

// ParseCommits parses commits
func ParseCommits(path string) ([]model.ConventionalCommit, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	tagrefs, err := r.Tags()
	if err != nil {
		return nil, err
	}

	var recentTag *plumbing.Reference
	err = tagrefs.ForEach(func(tag *plumbing.Reference) error {
		recentTag = tag
		return nil
	})
	if err != nil {
		return nil, err
	}

	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	var found = false
	var commits []model.ConventionalCommit

	err = cIter.ForEach(func(c *object.Commit) error {
		if recentTag != nil && c.Hash == recentTag.Hash() {
			found = true
		}
		if found == true {
			return nil
		}
		tmp := pattern.FindStringSubmatch(c.Message)

		commit := model.ConventionalCommit{
			Hash:         c.Hash.String(),
			Type:         tmp[1],
			Component:    tmp[2],
			Description:  tmp[3],
			Body:         tmp[4],
			Footer:       tmp[5],
			SemVerChange: model.Patch,
		}

		if commit.Type == "feat" {
			commit.SemVerChange = model.Minor
		}

		if strings.Contains(commit.Body, breakingChange) {
			commit.SemVerChange = model.Major
		}
		if strings.Contains(commit.Footer, breakingChange) {
			commit.SemVerChange = model.Major
		}

		commits = append(commits, commit)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return commits, nil
}

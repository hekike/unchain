package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hekike/conventional-commits/pkg/model"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var pattern = regexp.MustCompile(`^(?:(\w+)\(?(\w+)\)?: (.+))(?:(?:\r?\n|$){0,2}(.+))?(?:(?:\r?\n|$){0,2}(.+))?(?:\r?\n|$){0,2}`)
var versionPattern = regexp.MustCompile(`^update for version ((([0-9]+)\.([0-9]+)\.([0-9]+)(?:-([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)$`)
var breakingChange = "BREAKING CHANGE: "

// ParseCommits parses commits
func ParseCommits(dir string) ([]model.ConventionalCommit, error) {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return nil, fmt.Errorf("[ParseCommits] open repo: %v", err)
	}

	ref, err := r.Head()
	if err != nil {
		return nil, fmt.Errorf("[ParseCommits] head: %v", err)
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, fmt.Errorf("[ParseCommits] git log: %v", err)
	}

	var found = false
	var commits []model.ConventionalCommit

	err = cIter.ForEach(func(c *object.Commit) error {
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

		// Detect last semver bump
		tmp = versionPattern.FindStringSubmatch(commit.Description)
		if commit.Type == "chore" && commit.Component == "changelog" &&
			len(tmp) > 0 {
			found = true
			commit.SemVer = tmp[1]
		}

		if commit.Type == "feat" {
			commit.SemVerChange = model.Minor
		}

		if strings.Contains(commit.Body, breakingChange) {
			commit.SemVerChange = model.Major
			commit.Breaking = commit.Body[len(breakingChange):]
		}
		if strings.Contains(commit.Footer, breakingChange) {
			commit.SemVerChange = model.Major
			commit.Breaking = commit.Footer[len(breakingChange):]
		}

		commits = append(commits, commit)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ParseCommits] parse: %v", err)
	}

	return commits, nil
}

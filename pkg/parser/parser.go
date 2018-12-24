package parser

import (
	"fmt"
	"regexp"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// SemVerChange describes the semver change type
type SemVerChange string

const (
	// Patch semver change
	Patch SemVerChange = "patch"
	// Minor semver change
	Minor SemVerChange = "minor"
	// Major semver change
	Major SemVerChange = "major"
)

// ConventionalCommit parsed commit
type ConventionalCommit struct {
	Hash         string
	Type         string
	Component    string
	Description  string
	Body         string
	Footer       string
	Breaking     string
	SemVerChange SemVerChange
	SemVer       string
}

var pattern = regexp.MustCompile(`^(?:(\w+)\(?(\w+)\)?: (.+))(?:(?:\r?\n|$){0,2}(.+))?(?:(?:\r?\n|$){0,2}(.+))?(?:\r?\n|$){0,2}`)
var versionPattern = regexp.MustCompile(`^update for version ((([0-9]+)\.([0-9]+)\.([0-9]+)(?:-([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?)$`)
var breakingChange = "BREAKING CHANGE: "

// ParseCommits parses commits
func ParseCommits(dir string) ([]ConventionalCommit, error) {
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
	var commits []ConventionalCommit

	err = cIter.ForEach(func(c *object.Commit) error {
		if found == true {
			return nil
		}
		tmp := pattern.FindStringSubmatch(c.Message)

		// Skip commit that doesn't follow the conventional format
		if len(tmp) < 6 {
			return nil
		}

		commit := ConventionalCommit{
			Hash:         c.Hash.String(),
			Type:         tmp[1],
			Component:    tmp[2],
			Description:  tmp[3],
			Body:         tmp[4],
			Footer:       tmp[5],
			SemVerChange: Patch,
		}

		// Detect last semver bump
		tmp = versionPattern.FindStringSubmatch(commit.Description)
		if commit.Type == "chore" && commit.Component == "changelog" &&
			len(tmp) > 0 {
			found = true
			commit.SemVer = tmp[1]
		}

		if commit.Type == "feat" {
			commit.SemVerChange = Minor
		}

		if strings.Contains(commit.Body, breakingChange) {
			commit.SemVerChange = Major
			commit.Breaking = commit.Body[len(breakingChange):]
		}
		if strings.Contains(commit.Footer, breakingChange) {
			commit.SemVerChange = Major
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

// ToSemVerChange converts string to SemVerChange
func ToSemVerChange(changeFlag string) (change SemVerChange) {
	switch changeFlag {
	case "patch":
		change = Patch
	case "minor":
		change = Minor
	case "major":
		change = Major
	}
	return change
}

package changelog

import (
	"fmt"

	"github.com/hekike/unchain/pkg/git"
	"github.com/hekike/unchain/pkg/parser"
)

// Save generates and adds changelog.md to Git
func Save(
	dir string,
	version string,
	commits []parser.ConventionalCommit,
	user *git.User,
) (
	string,
	string,
	error,
) {
	// Generate changelog
	markdown := Generate(version, commits)

	// Write changelog
	file, err := Prepend(dir, markdown)
	if err != nil {
		return file, markdown, fmt.Errorf("[Save] prepend: %v", err)
	}

	// Add to Git
	err = GitCommit(dir, version, user)
	if err != nil {
		return file, markdown, fmt.Errorf("[Save] git commit: %v", err)
	}

	return file, markdown, nil
}

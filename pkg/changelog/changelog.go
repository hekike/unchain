package changelog

import (
	"fmt"

	"github.com/hekike/conventional-commits/pkg/parser"
)

// Save generates and adds changelog.md to Git
func Save(dir string, version string, commits []parser.ConventionalCommit) (
	string,
	error,
) {
	// Generate changelog
	markdown := Generate(version, commits)

	// Write changelog
	err := Prepend(dir, markdown)
	if err != nil {
		return markdown, fmt.Errorf("[Save] prepend: %v", err)
	}

	// Add to Git
	err = GitCommit(dir, version)
	if err != nil {
		return markdown, fmt.Errorf("[Save] git commit: %v", err)
	}

	return markdown, nil
}

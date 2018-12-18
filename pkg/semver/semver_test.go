package semver

import (
	"testing"

	"github.com/hekike/conventional-commits/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestGetChange(t *testing.T) {
	commits := []parser.ConventionalCommit{}

	commits = append(commits, parser.ConventionalCommit{SemVerChange: parser.Patch})
	res := GetChange(commits)
	assert.Equal(t, parser.Patch, res)

	commits = append(commits, parser.ConventionalCommit{SemVerChange: parser.Minor})
	res = GetChange(commits)
	assert.Equal(t, parser.Minor, res)

	commits = append(commits, parser.ConventionalCommit{SemVerChange: parser.Major})
	res = GetChange(commits)
	assert.Equal(t, parser.Major, res)
}

func TestGetVersion(t *testing.T) {
	res, err := GetVersion("", parser.Patch)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "1.0.0", res)

	res, err = GetVersion("1.0.0", parser.Patch)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "1.0.1", res)

	res, err = GetVersion("1.0.0", parser.Minor)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "1.1.0", res)

	res, err = GetVersion("1.0.0", parser.Major)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "2.0.0", res)
}

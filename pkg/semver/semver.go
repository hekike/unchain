package semver

import (
	"fmt"

	sv "github.com/coreos/go-semver/semver"
	"github.com/hekike/conventional-commits/pkg/model"
)

// GetChange determinate semver changes (patch, minor, major)
func GetChange(commits []model.ConventionalCommit) model.SemVerChange {
	var change model.SemVerChange = model.Patch
	for _, commit := range commits {
		if change != model.Major && commit.SemVerChange == model.Minor {
			change = model.Minor
		}
		if commit.SemVerChange == model.Major {
			change = model.Major
		}
	}
	return change
}

// GetVersion calculate version
func GetVersion(version string, change model.SemVerChange) (string, error) {
	if version == "" {
		return "1.0.0", nil
	}

	v, err := sv.NewVersion(version)
	if err != nil {
		return "", fmt.Errorf(
			"[semver.GetVersion] parse version (%s): %v",
			version,
			err,
		)
	}

	switch change {
	case model.Patch:
		v.BumpPatch()
	case model.Minor:
		v.BumpMinor()
	case model.Major:
		v.BumpMajor()
	default:
		return "", fmt.Errorf("Invalid change type %s", change)
	}

	return v.String(), nil
}

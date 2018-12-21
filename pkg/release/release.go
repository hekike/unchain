package release

import (
	"fmt"

	"github.com/hekike/unchain/pkg/changelog"
	"github.com/hekike/unchain/pkg/git"
	"github.com/hekike/unchain/pkg/npm"
	"github.com/hekike/unchain/pkg/parser"
	"github.com/hekike/unchain/pkg/semver"
)

// Result result of release
type Result struct {
	Message string
	Error   error
}

// Release generate changelog and tag release
func Release(path string, ch chan Result) {
	defer close(ch)

	// Read Git User
	user, err := git.GetUser(path)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Git] get user: %v", err),
		}
		return
	}

	// Parse commits
	commits, err := parser.ParseCommits(path)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Release] parse commits: %v", err),
		}
		return
	}
	if len(commits) == 0 {
		ch <- Result{
			Message: "No new commit found",
		}
	}

	// Read version from last bump commit if exist
	var version string
	if len(commits) > 0 {
		lastCommit := commits[len(commits)-1]
		if lastCommit.SemVer != "" {
			version = lastCommit.SemVer
		}
	}

	// Read version from npm (package.json) if exist
	var npmVersion string
	isNpm := npm.HasPackage(path)
	if isNpm {
		pkg, err := npm.ParsePackage(path)
		if err != nil {
			ch <- Result{
				Error: fmt.Errorf(
					"[Release] parse npm package: %v",
					err,
				),
			}
			return
		}
		npmVersion = pkg.Version
	}

	// Inconsistency between commit history and package.json version
	if npmVersion != "" && npmVersion != version {
		ch <- Result{
			Message: fmt.Sprintf(
				"Inconsistency between package.json's version field %s and version found in git history %s\n",
				npmVersion,
				version,
			),
		}
		ch <- Result{
			Message: "Will use the version from the package.json",
		}
		version = npmVersion
	}

	// Calculate new version
	change := semver.GetChange(commits)
	newVersion, err := semver.GetVersion(version, change)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Release] get semver version: %v", err),
		}
		return
	}

	// Generate changelog
	_, err = changelog.Save(path, newVersion, commits, user)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Release] save changelog: %v", err),
		}
		return
	}

	// Version: npm
	if isNpm {
		_, err = npm.Version(path, newVersion, string(change))
		if err != nil {
			ch <- Result{
				Error: fmt.Errorf("[npm] version: %v", err),
			}
			return
		}
	}

	// Release: Git
	err = git.Release(path, newVersion, user)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Release] git: %v", err),
		}
		return
	}

	// Publish: npm
	if isNpm {
		_, err = npm.Publish(path)
		if err != nil {
			ch <- Result{
				Error: fmt.Errorf("[npm] publish: %v", err),
			}
			return
		}
	}
}

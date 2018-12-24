package release

import (
	"fmt"
	"strconv"

	"github.com/hekike/unchain/pkg/changelog"
	"github.com/hekike/unchain/pkg/git"
	"github.com/hekike/unchain/pkg/npm"
	"github.com/hekike/unchain/pkg/parser"
	"github.com/hekike/unchain/pkg/semver"
)

// Phase describes the release phase
type Phase string

const (
	// PhaseGetGitUser phase
	PhaseGetGitUser Phase = "PhaseGetGitUser"
	// PhaseParseCommits phase
	PhaseParseCommits Phase = "PhaseParseCommits"
	// PhaseLastVersionFromCommit phase
	PhaseLastVersionFromCommit Phase = "PhaseLastVersionFromCommit"
	// PhaseLastVersionFromPackage phase
	PhaseLastVersionFromPackage Phase = "PhaseLastVersionFromPackage"
	// PhaseLastVersionInconsistency phase
	PhaseLastVersionInconsistency Phase = "PhaseLastVersionInconsistency"
	// PhaseChangeFound phase
	PhaseChangeFound Phase = "PhaseChangeFound"
	// PhaseNextVersion phase
	PhaseNextVersion Phase = "PhaseNextVersion"
	// PhaseChangelogUpdated phase
	PhaseChangelogUpdated Phase = "PhaseChangelogUpdated"
	// PhasePackageVersion phase
	PhasePackageVersion Phase = "PhasePackageVersion"
	// PhaseGitRelease phase
	PhaseGitRelease Phase = "PhaseGitRelease"
	// PhasePackagePublish phase
	PhasePackagePublish Phase = "PhasePackagePublish"
)

// Result result of release
type Result struct {
	Phase Phase
	Message string
	Error   error
}

// Release generate changelog and tag release
func Release(path string, change parser.SemVerChange, ch chan Result) {
	defer close(ch)

	// Get Git User
	user, err := git.GetUser(path)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Git] get user: %v", err),
		}
		return
	}
	ch <- Result{
		Phase: PhaseGetGitUser,
		Message: user.String(),
	}

	// Parse Commits
	commits, err := parser.ParseCommits(path)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Release] parse commits: %v", err),
		}
		return
	}
	ch <- Result{
		Phase: PhaseParseCommits,
		Message: strconv.Itoa(len(commits)),
	}

	// Read version from last bump commit if exist
	var version string
	if len(commits) > 0 {
		lastCommit := commits[len(commits)-1]
		if lastCommit.SemVer != "" {
			version = lastCommit.SemVer
			ch <- Result{
				Phase: PhaseLastVersionFromCommit,
				Message: version,
			}
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
		ch <- Result{
			Phase: PhaseLastVersionFromPackage,
			Message: npmVersion,
		}
	}

	// Inconsistency between commit history and package.json version
	if npmVersion != "" && npmVersion != version {
		ch <- Result{
			Phase: PhaseLastVersionInconsistency,
			Message: fmt.Sprintf(
				"package.json: %s, git: %s",
				npmVersion,
				version,
			),
		}
		version = npmVersion
	}

	// Find Change
	if (change == "") {
		change = semver.GetChange(commits)
		ch <- Result{
			Phase: PhaseChangeFound,
			Message: string(change),
		}
	}

	// Calculate new version
	newVersion, err := semver.GetVersion(version, change)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf(
				"[Release] get semver version: %v",
				err,
			),
		}
		return
	}
	ch <- Result{
		Phase: PhaseNextVersion,
		Message: newVersion,
	}

	// Generate changelog
	cf, _, err := changelog.Save(path, newVersion, commits, user)
	if err != nil {
		ch <- Result{
			Error: fmt.Errorf("[Release] save changelog: %v", err),
		}
		return
	}
	ch <- Result{
		Phase: PhaseChangelogUpdated,
		Message: cf,
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
		ch <- Result{
			Phase: PhasePackageVersion,
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
	ch <- Result{
		Phase: PhaseGitRelease,
		Message: newVersion,
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
		ch <- Result{
			Phase: PhasePackagePublish,
		}
	}
}

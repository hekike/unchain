package main

import (
	"fmt"
	"os"

	"github.com/hekike/conventional-commits/pkg/changelog"
	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/npm"
	"github.com/hekike/conventional-commits/pkg/parser"
	"github.com/hekike/conventional-commits/pkg/semver"
	"github.com/hekike/conventional-commits/pkg/release"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	// Parse commits
	commits, err := parser.ParseCommits(path)
	if err != nil {
		panic(err)
	}
	if len(commits) == 0 {
		fmt.Println("No new commit found")
	}

	// Read version from last bump commit if exist
	var version string
	if len(commits) > 0 {
		lastCommit := commits[len(commits) - 1]
		if (lastCommit.SemVer != "") {
			version = lastCommit.SemVer
		}
	}

	// Read version from npm (package.json) if exist
	var npmVersion string
	isNpm := npm.HasPackage(path)
	if isNpm {
		pkg, err := npm.ParsePackage(path)
		if err != nil {
			panic(err)
		}
		npmVersion = pkg.Version
	}

	// Inconsistency between commit history and package.json version
	if (npmVersion != "" && npmVersion != version) {
		fmt.Printf(
			"Inconsistency between package.json's version field %s and version found in git history %s\n",
			npmVersion,
			version,
		)
		fmt.Println("Will use the version from the package.json")
		version = npmVersion
	}

	// Calculate new version
	change := semver.GetChange(commits)
	newVersion, err := semver.GetVersion(version, change)
	if err != nil {
		panic(err)
	}

	// Generate changelog
	_, err = changelog.Save(path, newVersion, commits)
	if err != nil {
		panic(err)
	}

	// Release: npm
	if isNpm {
		_, err = npm.Bump(path, newVersion, string(change))
		if err != nil {
			panic(err)
		}
	} else {
		// Release: git
		_, err = release.GitTag(path, newVersion)
		if err != nil {
			panic(err)
		}
	}
}

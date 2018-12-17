package main

import (
	"fmt"
	"os"

	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/markdown"
	"github.com/hekike/conventional-commits/pkg/npm"
	"github.com/hekike/conventional-commits/pkg/parser"
	"github.com/hekike/conventional-commits/pkg/semver"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	// Parse commits
	commits, err := parser.ParseCommits(path)
	if err != nil {
		panic(err)
	}

	var version string

	// Read version from npm (package.json) if exist
	if npm.HasPackage(path) {
		pkg, err := npm.ParsePackage(path)
		if err != nil {
			panic(err)
		}
		version = pkg.Version
	}

	// Read version from Git Tag tag if exist
	if version == "" {
		tag, err := parser.GetLastTag(path)
		if err != nil {
			panic(err)
		}
		if tag != nil {
			version = tag.Name
		}
	}

	// Calculate new version
	change := semver.GetChange(commits)
	newVersion, err := semver.GetVersion(version, change)
	if err != nil {
		panic(err)
	}

	// Generate output
	output := markdown.Generate(newVersion, commits)
	fmt.Println(output)

}

package main

import (
	"fmt"
	"os"

	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/markdown"
	"github.com/hekike/conventional-commits/pkg/parser"
	"github.com/hekike/conventional-commits/pkg/semver"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	tag, err := parser.GetLastTag(path)
	if err != nil {
		panic(err)
	}

	commits, err := parser.ParseCommits(path)
	if err != nil {
		panic(err)
	}

	change := semver.GetChange(commits)
	version, err := semver.GetVersion(tag, change)
	if err != nil {
		panic(err)
	}

	output := markdown.Generate(version, commits)
	fmt.Println(output)

}

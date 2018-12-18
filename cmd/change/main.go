package main

import (
	"fmt"
	"os"

	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/parser"
	"github.com/hekike/conventional-commits/pkg/semver"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	commits, err := parser.ParseCommits(path)

	if err != nil {
		panic(err)
	}

	change := semver.GetChange(commits)
	fmt.Println(change)

}

package main

import (
	"fmt"
	"os"

	"github.com/hekike/unchain/pkg/cli-tools"
	"github.com/hekike/unchain/pkg/parser"
	"github.com/hekike/unchain/pkg/semver"
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

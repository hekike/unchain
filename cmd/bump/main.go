package main

import (
	"fmt"
	"os"

	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/model"
	"github.com/hekike/conventional-commits/pkg/parser"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	commits, err := parser.ParseCommits(path)

	if err != nil {
		panic(err)
	}

	var change model.SemVerChange = model.Patch
	for _, commit := range commits {
		if change != model.Major && commit.SemVerChange == model.Minor {
			change = model.Minor
		}
		if commit.SemVerChange == model.Major {
			change = model.Major
		}
	}
	fmt.Println(change)

}

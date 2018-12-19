package main

import (
	"fmt"
	"os"

	"github.com/hekike/unchain/pkg/cli-tools"
	"github.com/hekike/unchain/pkg/parser"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	commits, err := parser.ParseCommits(path)

	if err != nil {
		panic(err)
	}

	fmt.Println("hash,semver,type,component,description,body,footer")
	for _, commit := range commits {
		o := fmt.Sprintf(
			"%s,%s,%s,%s,%s,%s,%s",
			commit.Hash,
			commit.SemVerChange,
			commit.Type,
			commit.Component,
			commit.Description,
			commit.Body,
			commit.Footer,
		)
		fmt.Println(o)
	}

}

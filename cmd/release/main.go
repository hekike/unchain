package main

import (
	"fmt"
	"os"

	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/release"
)

func main() {
	clitools.CheckArgs("<path>")
	path := os.Args[1]

	// Start release
	results := make(chan release.Result)
	go release.Release(path, results)

	// Results
	for res := range results {
		if res.Error != nil {
			panic(fmt.Errorf("[cmd] release: %v", res.Error))
		}

		fmt.Println(res.Message)
	}
}

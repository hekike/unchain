package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hekike/conventional-commits/pkg/cli-tools"
	"github.com/hekike/conventional-commits/pkg/release"
)

func main() {
	var path string

	if (len(os.Args)) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		path = dir
	} else {
		clitools.CheckArgs("<path>")
		path = os.Args[1]
	}

	// Start release
	results := make(chan release.Result)
	go release.Release(path, results)

	// Results
	for res := range results {
		if res.Error != nil {
			log.Fatal(res.Error)
			os.Exit(1)
		}

		fmt.Println(res.Message)
	}
}

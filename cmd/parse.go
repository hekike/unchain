package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hekike/unchain/pkg/parser"
	"github.com/spf13/cobra"
)

// GetParseCmd returns the parse cmd
func GetParseCmd() *cobra.Command {
	// Default dir is the working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var cmdParse = &cobra.Command{
		Use:   "parse",
		Short: "Parses commits",
		Long:  `parsing conventional commits.`,
		Run: func(cmd *cobra.Command, args []string) {
			commits, err := parser.ParseCommits(dir)

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
		},
	}

	cmdParse.Flags().StringVarP(
		&dir,
		"repository",
		"r",
		"",
		"Repository directory",
	)

	return cmdParse
}

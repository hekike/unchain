package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hekike/unchain/pkg/release"
	"github.com/spf13/cobra"
)

// GetReleaseCmd returns the release cmd
func GetReleaseCmd() *cobra.Command {
	// Default dir is the working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var cmdRelease = &cobra.Command{
		Use:   "release",
		Short: "Creates a new release",
		Long: `release is generating a changelog based on the
conventional commits and determinates the next version. It also
tags the Git repository with the new version and pushed the change to
remote. For npm libraries it also bumps the package.json file.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Start release
			results := make(chan release.Result)
			go release.Release(dir, results)

			// Results
			for res := range results {
				if res.Error != nil {
					log.Fatal(res.Error)
					os.Exit(1)
				}

				fmt.Println(res.Message)
			}
		},
	}

	cmdRelease.Flags().StringVarP(
		&dir,
		"repository",
		"r",
		"",
		"Repository directory",
	)

	return cmdRelease
}

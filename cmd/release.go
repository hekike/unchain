package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hekike/unchain/pkg/parser"
	"github.com/hekike/unchain/pkg/release"
	"github.com/spf13/cobra"
)

// GetReleaseCmd returns the release cmd
func GetReleaseCmd() *cobra.Command {
	var changeFlag string

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
			// Parse optional semver change flag
			change, err := parseChangeFlag(changeFlag)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			// Start release
			results := make(chan release.Result)
			go release.Release(dir, change, results)

			// Results
			statusCounter := 0
			for res := range results {
				// Error
				if res.Error != nil {
					log.Fatal(res.Error)
					os.Exit(1)
				}

				// Result
				statusCounter++
				fmt.Printf("%d. ", statusCounter)

				handleResult(res)
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

	cmdRelease.Flags().StringVarP(
		&changeFlag,
		"change",
		"c",
		"",
		"SemVer change (patch|minor|major)",
	)

	return cmdRelease
}

func parseChangeFlag(changeFlag string) (change parser.SemVerChange, err error) {
	if changeFlag == "" {
		return change, err
	}

	change = parser.ToSemVerChange(changeFlag)
	if change == "" {
		err = fmt.Errorf("Invalid semver change input: %s", changeFlag)
	}

	return change, err
}

func handleResult(res release.Result) {
	switch res.Phase {
	case release.PhaseGetGitUser:
		fmt.Printf(
			"git user found (%s)\n",
			res.Message,
		)
	case release.PhaseParseCommits:
		fmt.Printf(
			"commits parsed (%s)\n",
			res.Message,
		)
	case release.PhaseLastVersionFromCommit:
		fmt.Printf(
			"last version found in commits (%s)\n",
			res.Message,
		)
	case release.PhaseLastVersionFromPackage:
		fmt.Printf(
			"last version found in package.json (%s)\n",
			res.Message,
		)
	case release.PhaseLastVersionInconsistency:
		fmt.Printf("last version inconsistency (%s)\n", res.Message)
		log.Fatal("update your package.json file")
		os.Exit(1)

	case release.PhaseChangeFound:
		fmt.Printf(
			"change found (%s)\n",
			res.Message,
		)
	case release.PhaseNextVersion:
		fmt.Printf(
			"next version calculated (%s)\n",
			res.Message,
		)
	case release.PhaseChangelogUpdated:
		fmt.Printf(
			"changelog updated (%s)\n",
			res.Message,
		)
	case release.PhasePackageVersion:
		fmt.Print(
			"package version bumped\n",
		)
	case release.PhaseGitRelease:
		fmt.Printf(
			"git tagged and pushed (%s)\n",
			res.Message,
		)
	case release.PhasePackagePublish:
		fmt.Print(
			"package published\n",
		)
	default:
		fmt.Println(res.Message)
	}
}

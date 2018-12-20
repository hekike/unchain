package main

import (
	"github.com/hekike/unchain/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "unchain"}
	rootCmd.AddCommand(cmd.GetParseCmd())
	rootCmd.AddCommand(cmd.GetReleaseCmd())
	rootCmd.AddCommand(cmd.GetSemverCmd())
	rootCmd.Execute()
}

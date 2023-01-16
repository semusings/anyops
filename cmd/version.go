package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ReleaseVersion can be replaced by build tooling
var ReleaseVersion string

// GitVersion should be replaced by the makefile
var GitVersion string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CLI",
	Long:  `Print the version number of CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (git-%s)\n", ReleaseVersion, GitVersion)
	},
}

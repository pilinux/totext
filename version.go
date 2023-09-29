// Package totext uses https://github.com/sajari/docconv
// to extract text from different file types.
package totext

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Build information
var (
	AppVersion string
	BuildDate  string
	CommitHash string
	Author     string
)

// Version returns the version number, build date and commit hash
func Version(appName string) *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number, build date and commit hash",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(appName, "version:", AppVersion)
			fmt.Println("Build date:", BuildDate)
			fmt.Println("Commit hash:", CommitHash)
			fmt.Println("Author:", Author)
		},
	}

	return versionCmd
}

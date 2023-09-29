package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
	"github.com/pilinux/totext/cli"
)

func main() {
	appName := "totextcli"

	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   appName,
		Short: "Extract text from different document types",
	}

	// Define the subcommands
	var docCmd = cli.DocCmd(appName)
	var docxCmd = cli.DocxCmd(appName)
	var odtCmd = cli.OdtCmd(appName)
	var pdfCmd = cli.PdfCmd(appName)
	var versionCmd = totext.Version(appName)

	// Add the commands to the root command
	rootCmd.AddCommand(
		docCmd,
		docxCmd,
		odtCmd,
		pdfCmd,
		versionCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

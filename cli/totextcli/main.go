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
	var fileCmd = cli.FileCmd(appName)
	var htmlCmd = cli.HTMLCmd(appName)
	var odtCmd = cli.OdtCmd(appName)
	var pdfCmd = cli.PdfCmd(appName)
	var rtfCmd = cli.RtfCmd(appName)
	var urlCmd = cli.URLCmd(appName)
	var versionCmd = totext.Version(appName)

	// Add the commands to the root command
	rootCmd.AddCommand(
		docCmd,
		docxCmd,
		fileCmd,
		htmlCmd,
		odtCmd,
		pdfCmd,
		rtfCmd,
		urlCmd,
		versionCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

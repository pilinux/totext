package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-rod/rod"
	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertURLToText receives url as an argument and writes
// its text content and metadata into two separate files
func ConvertURLToText(inputURL string, skipPrettifyError bool, delayInSec int) error {
	inputURL = strings.TrimSpace(inputURL)

	// Create a new browser instance
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Fetch the HTML page and convert to text
	htmlFilename, content, metadata, err := totext.ConvertURLToText(browser, inputURL, skipPrettifyError, delayInSec)
	if err != nil {
		return err
	}

	// Get filename without extension from htmlFile
	filenameWithoutExtension := strings.TrimSuffix(htmlFilename, ".html")

	// Write content to a txt file
	err = totext.WriteText(filenameWithoutExtension+".txt", content)
	if err != nil {
		return err
	}

	// Write metadata to a txt file
	err = totext.WriteText(
		filenameWithoutExtension+"_metadata.txt",
		fmt.Sprintf("%v", metadata),
	)
	if err != nil {
		return err
	}

	return nil
}

// URLCmd defines the "url" command
func URLCmd(appName string) *cobra.Command {
	var urlCmd = &cobra.Command{
		Use:   "url",
		Short: "Fetch HTML page from the URL and write the extracted text to a txt file",
		Args:  cobra.MinimumNArgs(1), // full URL
		Run: func(cmd *cobra.Command, args []string) {
			// Get the value of the skipPrettifyError flag
			skipPrettifyError, err := cmd.Flags().GetBool("skipPrettifyError")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if skipPrettifyError {
				fmt.Println("Skipping prettify error")
			}

			// Get the value of the delayInSec flag
			delayInSec, err := cmd.Flags().GetInt("delayInSec")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Convert HTML page from the given URL to text
			err = ConvertURLToText(args[0], skipPrettifyError, delayInSec)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	// Add the skipPrettifyError flag as an optional argument
	urlCmd.Flags().BoolP(
		"skipPrettifyError",
		"s",
		false,
		"skip prettify error",
	)
	// Add the delayInSec flag as an optional argument
	urlCmd.Flags().IntP(
		"delayInSec",
		"d",
		0,
		"additional delay in seconds for the web page to load",
	)
	urlCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, urlCmd.Use, "['https://example.com/path/to/webpage'] [--skipPrettifyError or -s] [--delayInSec=<seconds> or -d <seconds>]")
		return nil
	})

	return urlCmd
}

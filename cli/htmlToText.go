package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertHTMLToText receives HTML filepath as an argument and
// writes its text content and metadata into two separate files
func ConvertHTMLToText(filepath string, skipPrettifyError bool) error {
	filepath = strings.TrimSpace(filepath)

	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)
	if fileExt != totext.HTML {
		return fmt.Errorf("file type not supported")
	}

	// Convert HTML to text
	content, metadata, err := totext.ConvertHTMLToText(filepath, skipPrettifyError)
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, ".html")

	// Set current working directory
	err = totext.SetCwd(filepath)
	if err != nil {
		return err
	}

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

// HTMLCmd defines the "html" command
func HTMLCmd(appName string) *cobra.Command {
	var htmlCmd = &cobra.Command{
		Use:   "html",
		Short: "Extract text content from an HTML file and write it to a txt file",
		Args:  cobra.MinimumNArgs(1), // html filepath
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

			// Convert HTML to text
			err = ConvertHTMLToText(args[0], skipPrettifyError)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	// Add the skipPrettifyError flag as an optional argument
	htmlCmd.Flags().BoolP(
		"skipPrettifyError",
		"s",
		false,
		"skip prettify error",
	)
	htmlCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, htmlCmd.Use, "[file.html or /path/to/file.html] [--skipPrettifyError or -s]")
		return nil
	})

	return htmlCmd
}

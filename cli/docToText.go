package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertDocToText receives MS word doc filepath as an argument
// writes its text content and metadata into two separate files
func ConvertDocToText(filepath string) error {
	filepath = strings.TrimSpace(filepath)

	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)
	if fileExt != totext.DOC {
		return fmt.Errorf("file type not supported")
	}

	// Convert doc to text
	content, metadata, err := totext.ConvertDocToText(filepath)
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, ".doc")

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

// DocCmd defines the "doc" command
func DocCmd(appName string) *cobra.Command {
	var docCmd = &cobra.Command{
		Use:   "doc",
		Short: "Extract text from a MS word doc file and write it to a txt file",
		Args:  cobra.ExactArgs(1), // doc filepath
		Run: func(cmd *cobra.Command, args []string) {
			// Convert doc to text
			err := ConvertDocToText(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	docCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, docCmd.Use, "[file.doc or /path/to/file.doc]")
		return nil
	})

	return docCmd
}

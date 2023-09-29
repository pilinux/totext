package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertDocxToText receives MS word docx filepath as an argument
// writes its text content and metadata into two separate files
func ConvertDocxToText(filepath string) error {
	filepath = strings.TrimSpace(filepath)

	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)
	if fileExt != totext.DOCX {
		return fmt.Errorf("file type not supported")
	}

	// Convert docx to text
	content, metadata, err := totext.ConvertDocxToText(filepath)
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, ".docx")

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

// DocxCmd defines the "docx" command
func DocxCmd(appName string) *cobra.Command {
	var docxCmd = &cobra.Command{
		Use:   "docx",
		Short: "Extract text from a MS word docx file and write it to a txt file",
		Args:  cobra.ExactArgs(1), // docx filepath
		Run: func(cmd *cobra.Command, args []string) {
			// Convert docx to text
			err := ConvertDocxToText(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	docxCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, docxCmd.Use, "[file.docx or /path/to/file.docx]")
		return nil
	})

	return docxCmd
}

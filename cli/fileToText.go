package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertFileToText receives filepath as an argument and writes
// its text content and metadata into two separate files
func ConvertFileToText(filepath string) error {
	filepath = strings.TrimSpace(filepath)

	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)

	var content string
	var metadata map[string]string
	var err error

	switch fileExt {
	case totext.DOC:
		// Convert doc to text
		content, metadata, err = totext.ConvertDocToText(filepath)
	case totext.DOCX:
		// Convert docx to text
		content, metadata, err = totext.ConvertDocxToText(filepath)
	case totext.ODT:
		// Convert odt to text
		content, metadata, err = totext.ConvertOdtToText(filepath)
	case totext.PDF:
		// Convert PDF to text
		content, metadata, err = totext.ConvertPDFToText(filepath)
	case totext.RTF:
		// Convert rtf to text
		content, metadata, err = totext.ConvertRTFToText(filepath)
	default:
		return fmt.Errorf("file type not supported")
	}
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, "."+string(fileExt))

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

// FileCmd defines the "file" command
func FileCmd(appName string) *cobra.Command {
	var fileCmd = &cobra.Command{
		Use:   "file",
		Short: "Extract text from a file and write it to a txt file",
		Args:  cobra.ExactArgs(1), // filepath
		Run: func(cmd *cobra.Command, args []string) {
			// Convert file to text
			err := ConvertFileToText(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	fileCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, fileCmd.Use, "[file.extension or /path/to/file.extension]")
		return nil
	})

	return fileCmd
}

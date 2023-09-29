package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertOdtToText receives odt filepath as an argument and writes
// its text content and metadata into two separate files
func ConvertOdtToText(filepath string) error {
	filepath = strings.TrimSpace(filepath)

	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)
	if fileExt != totext.ODT {
		return fmt.Errorf("file type not supported")
	}

	// Convert odt to text
	content, metadata, err := totext.ConvertOdtToText(filepath)
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, ".odt")

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

// OdtCmd defines the "odt" command
func OdtCmd(appName string) *cobra.Command {
	var odtCmd = &cobra.Command{
		Use:   "odt",
		Short: "Extract text from an odt file and write it to a txt file",
		Args:  cobra.ExactArgs(1), // odt filepath
		Run: func(cmd *cobra.Command, args []string) {
			// Convert odt to text
			err := ConvertOdtToText(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	odtCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, odtCmd.Use, "[file.odt or /path/to/file.odt]")
		return nil
	})

	return odtCmd
}

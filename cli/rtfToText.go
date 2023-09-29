package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertRTFToText receives rtf filepath as an argument and writes
// its text content and metadata into two separate files
func ConvertRTFToText(filepath string) error {
	filepath = strings.TrimSpace(filepath)

	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)
	if fileExt != totext.RTF {
		return fmt.Errorf("file type not supported")
	}

	// Convert rtf to text
	content, metadata, err := totext.ConvertRTFToText(filepath)
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, ".rtf")

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

// RtfCmd defines the "rtf" command
func RtfCmd(appName string) *cobra.Command {
	var rtfCmd = &cobra.Command{
		Use:   "rtf",
		Short: "Extract text from an RTF file and write it to a txt file",
		Args:  cobra.ExactArgs(1), // rtf filepath
		Run: func(cmd *cobra.Command, args []string) {
			// Convert rtf to text
			err := ConvertRTFToText(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	rtfCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, rtfCmd.Use, "[file.rtf or /path/to/file.rtf]")
		return nil
	})

	return rtfCmd
}

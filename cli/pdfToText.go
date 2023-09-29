package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pilinux/totext"
)

// ConvertPDFToText receives pdf filepath as an argument and writes
// its text content and metadata into two separate files
func ConvertPDFToText(filepath string) error {
	// Get file extension from filepath
	fileExt := totext.GetFileExtension(filepath)
	if fileExt != totext.PDF {
		return fmt.Errorf("file type not supported")
	}

	// Convert PDF to text
	content, metadata, err := totext.ConvertPDFToText(filepath)
	if err != nil {
		return err
	}

	// Get filename from filepath
	filename := totext.GetFilename(filepath)
	filenameWithoutExtension := strings.TrimSuffix(filename, ".pdf")

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

// PdfCmd defines the "pdf" command
func PdfCmd(appName string) *cobra.Command {
	var pdfCmd = &cobra.Command{
		Use:   "pdf",
		Short: "Extract text from a PDF file and write it to a txt file",
		Args:  cobra.ExactArgs(1), // pdf filepath
		Run: func(cmd *cobra.Command, args []string) {
			// Convert PDF to text
			err := ConvertPDFToText(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	pdfCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("Usage:", appName, pdfCmd.Use, "[file.pdf or /path/to/file.pdf]")
		return nil
	})

	return pdfCmd
}

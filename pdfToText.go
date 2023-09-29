package totext

import (
	"os"

	"code.sajari.com/docconv"
)

// ConvertPDFToText receives pdf filepath as an argument and returns its text content and metadata
//
// Dependencies:
//
// Debian/Ubuntu: sudo apt install poppler-utils
//
// MacOS: brew install poppler
func ConvertPDFToText(filepath string) (content string, metadata map[string]string, err error) {
	// Get the PDF file
	pdfFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer pdfFile.Close()

	// Convert PDF to text
	content, metadata, err = docconv.ConvertPDF(pdfFile)
	return
}

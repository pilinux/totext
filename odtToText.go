package totext

import (
	"os"

	"code.sajari.com/docconv"
)

// ConvertOdtToText receives odt filepath as an argument and returns its text content and metadata
func ConvertOdtToText(filepath string) (content string, metadata map[string]string, err error) {
	// Get the odt file
	odtFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer func() {
		_ = odtFile.Close()
	}()

	// Convert odt to text
	content, metadata, err = docconv.ConvertODT(odtFile)
	if err != nil {
		return "", nil, err
	}

	// Filter out non-readable characters
	content = FilterNonReadableCharacter(content)

	return
}

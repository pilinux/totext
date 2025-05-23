package totext

import (
	"os"

	"code.sajari.com/docconv"
)

// ConvertDocxToText receives MS word docx filepath as an argument
// and returns its text content and metadata
func ConvertDocxToText(filepath string) (content string, metadata map[string]string, err error) {
	// Get the docx file
	docxFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer func() {
		_ = docxFile.Close()
	}()

	// Convert docx to text
	content, metadata, err = docconv.ConvertDocx(docxFile)
	if err != nil {
		return "", nil, err
	}

	// Filter out non-readable characters
	content = FilterNonReadableCharacter(content)

	return
}

package totext

import (
	"os"

	"code.sajari.com/docconv"
)

// ConvertDocToText receives MS word doc filepath as an argument
// and returns its text content and metadata
//
// Dependencies:
//
// Debian/Ubuntu: sudo apt install wv
//
// MacOS: brew install wv
func ConvertDocToText(filepath string) (content string, metadata map[string]string, err error) {
	// Get the doc file
	docFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer docFile.Close()

	// Convert doc to text
	content, metadata, err = docconv.ConvertDoc(docFile)
	return
}

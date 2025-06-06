package totext

import (
	"os"

	"code.sajari.com/docconv"
)

// ConvertRTFToText receives rtf filepath as an argument and returns its text content and metadata
//
// Dependencies:
//
// Debian/Ubuntu: sudo apt install unrtf
//
// MacOS: brew install unrtf
func ConvertRTFToText(filepath string) (content string, metadata map[string]string, err error) {
	// Get the rtf file
	rtfFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer func() {
		_ = rtfFile.Close()
	}()

	// Convert rtf to text
	content, metadata, err = docconv.ConvertRTF(rtfFile)
	if err != nil {
		return "", nil, err
	}

	// Filter out non-readable characters
	content = FilterNonReadableCharacter(content)

	return
}

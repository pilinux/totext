package totext

import (
	"os"

	"code.sajari.com/docconv"
)

// ConvertPagesToText receives pages filepath as an argument and returns its text content and metadata
func ConvertPagesToText(filepath string) (content string, metadata map[string]string, err error) {
	// Get the pages file
	pagesFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}

	// Convert pages to text
	content, metadata, err = docconv.ConvertPages(pagesFile)
	if err != nil {
		return "", nil, err
	}

	// Filter out non-readable characters
	content = FilterNonReadableCharacter(content)

	return
}

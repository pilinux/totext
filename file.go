package totext

import (
	"os"
	"strings"
)

// FileExtension is the file extension type
type FileExtension string

// File types
const (
	PDF  FileExtension = "pdf"
	DOC  FileExtension = "doc"
	DOCX FileExtension = "docx"
	ODT  FileExtension = "odt"
	TXT  FileExtension = "txt"
	MD   FileExtension = "md"
	RTF  FileExtension = "rtf"
	JSON FileExtension = "json"
)

// GetFileExtension returns the file extension of a file
func GetFileExtension(filepath string) FileExtension {
	// Get file extension
	fileExt := filepath[strings.LastIndex(filepath, ".")+1:]
	fileExt = strings.TrimSpace(fileExt)
	fileExt = strings.ToLower(fileExt)

	switch fileExt {
	case string(PDF):
		return PDF
	case string(DOC):
		return DOC
	case string(DOCX):
		return DOCX
	case string(ODT):
		return ODT
	case string(TXT):
		return TXT
	case string(MD):
		return MD
	case string(RTF):
		return RTF
	case string(JSON):
		return JSON

	default:
		return ""
	}
}

// GetFilename returns the filename of a file
func GetFilename(filepath string) string {
	// Get filename
	filename := filepath[strings.LastIndex(filepath, "/")+1:]
	filename = strings.TrimSpace(filename)
	filename = strings.ToLower(filename)

	return filename
}

// WriteText writes text content to a file
func WriteText(filepath string, content string) error {
	// Create file
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write content to the file
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// ReadText reads text content from a text file
func ReadText(filepath string) (string, error) {
	// Read the entire file into a string
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// DeleteFile deletes a file
func DeleteFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

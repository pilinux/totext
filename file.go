package totext

import (
	"os"
	"strings"
)

// FileExtension is the file extension type
type FileExtension string

// MIME types
type MIME string

// File types
const (
	DOC   FileExtension = "doc"
	DOCX  FileExtension = "docx"
	HTML  FileExtension = "html"
	JSON  FileExtension = "json"
	MD    FileExtension = "md"
	ODT   FileExtension = "odt"
	PAGES FileExtension = "pages"
	PDF   FileExtension = "pdf"
	RTF   FileExtension = "rtf"
	TXT   FileExtension = "txt"
)

// MIME types
const (
	MimeDOC   MIME = "application/msword"
	MimeDOCX  MIME = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MimeHTML  MIME = "text/html"
	MimeJSON  MIME = "application/json"
	MimeMD    MIME = "text/markdown"
	MimeODT   MIME = "application/vnd.oasis.opendocument.text"
	MimePAGES MIME = "application/vnd.apple.pages"
	MimePDF   MIME = "application/pdf"
	MimeRTF   MIME = "application/rtf"
	MimeTXT   MIME = "text/plain"
)

// GetFileExtension returns the file extension of a file
func GetFileExtension(filepath string) FileExtension {
	// Get file extension
	fileExt := filepath[strings.LastIndex(filepath, ".")+1:]
	fileExt = strings.TrimSpace(fileExt)
	fileExt = strings.ToLower(fileExt)

	switch fileExt {
	case string(DOC):
		return DOC
	case string(DOCX):
		return DOCX
	case string(HTML):
		return HTML
	case string(JSON):
		return JSON
	case string(MD):
		return MD
	case string(ODT):
		return ODT
	case string(PAGES):
		return PAGES
	case string(PDF):
		return PDF
	case string(RTF):
		return RTF
	case string(TXT):
		return TXT

	default:
		return ""
	}
}

// IsMIMETypeMatched compares *multipart.FileHeader MIME type with file extension
func IsMIMETypeMatched(fileExt FileExtension, mime MIME) bool {
	switch fileExt {
	case DOC:
		return mime == MimeDOC
	case DOCX:
		return mime == MimeDOCX
	case HTML:
		return mime == MimeHTML
	case JSON:
		return mime == MimeJSON
	case MD:
		return mime == MimeMD
	case ODT:
		return mime == MimeODT
	case PAGES:
		return mime == MimePAGES
	case PDF:
		return mime == MimePDF
	case RTF:
		return mime == MimeRTF
	case TXT:
		return mime == MimeTXT

	default:
		return false
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
func WriteText(filepath string, content string) (err error) {
	// Create file
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if e := f.Close(); e != nil && err == nil {
			err = e
		}
	}()

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

// IsAbsPath checks if the filepath is an absolute path
func IsAbsPath(filepath string) bool {
	return filepath != "" && filepath[0] == '/'
}

// SetCwd sets current working directory
func SetCwd(filepath string) error {
	filepath = strings.TrimSpace(filepath)
	filepath = strings.TrimSuffix(
		filepath,
		strings.TrimSpace(filepath[strings.LastIndex(filepath, "/")+1:]),
	)

	if !IsAbsPath(filepath) {
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}
		filepath = currentDir + "/" + filepath
	}

	err := os.Chdir(filepath)
	return err
}

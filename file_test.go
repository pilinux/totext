package totext

import "testing"

// TestGetFileExtension tests GetFileExtension function
func TestGetFileExtension(t *testing.T) {
	// Test data
	testData := []struct {
		filepath string
		expected FileExtension
	}{
		{"test.pdf", PDF},
		{"/path/to.here/test.pdf", PDF},

		{"test.doc", DOC},
		{"/path/to.here/test.pdf.DOC", DOC},

		{"test.docx", DOCX},
		{"test.odt", ODT},
		{"test.txt", TXT},
		{"test.md", MD},
		{"test.rtf", RTF},
		{"test.json", JSON},

		{"test", ""},
	}

	// Iterate over test data
	for _, data := range testData {
		// Get file extension
		fileExt := GetFileExtension(data.filepath)

		// Compare file extension
		if fileExt != data.expected {
			t.Errorf("Expected file extension %s, got %s", data.expected, fileExt)
		}
	}
}

// TestIsMIMETypeMatched tests IsMIMETypeMatched function
func TestIsMIMETypeMatched(t *testing.T) {
	// Test data
	testData := []struct {
		fileExt  FileExtension
		mime     MIME
		expected bool
	}{
		{DOC, MimeDOC, true},
		{DOC, MimeDOCX, false},
		{DOCX, MimeDOCX, true},
		{JSON, MimeJSON, true},
		{MD, MimeMD, true},
		{ODT, MimeODT, true},
		{PAGES, MimePAGES, true},
		{PDF, MimePDF, true},
		{RTF, MimeRTF, true},
		{TXT, MimeTXT, true},
	}

	// Iterate over test data
	for _, data := range testData {
		// Check if MIME type is matched
		isMIMETypeMatched := IsMIMETypeMatched(FileExtension(data.fileExt), data.mime)

		// Compare isMIMETypeMatched
		if isMIMETypeMatched != data.expected {
			t.Errorf(
				"Expected isMIMETypeMatched %t, got %t for file extension %s",
				data.expected, isMIMETypeMatched, data.fileExt,
			)
		}
	}
}

// TestGetFilename tests GetFilename function
func TestGetFilename(t *testing.T) {
	// Test data
	testData := []struct {
		filepath string
		expected string
	}{
		{"test.pdf", "test.pdf"},
		{"/path/to.here/test.pdf", "test.pdf"},

		{"test.doc", "test.doc"},
		{"/path/to.here/test.pdf.DOC", "test.pdf.doc"},

		{"test.docx", "test.docx"},
		{"/path/to/test.odt", "test.odt"},
		{"test.txt", "test.txt"},
		{"/path/to/test.md", "test.md"},
		{"test.rtf", "test.rtf"},
		{"/path/to/test.json", "test.json"},

		{"/path/to/test", "test"},
	}

	// Iterate over test data
	for _, data := range testData {
		// Get filename
		filename := GetFilename(data.filepath)

		// Compare filename
		if filename != data.expected {
			t.Errorf("Expected filename %s, got %s", data.expected, filename)
		}
	}
}

// TestWriteText tests WriteText function
func TestWriteText(t *testing.T) {
	// Test data
	testData := []struct {
		filepath string
		content  string
	}{
		{"test.txt", "Hello World"},
		{"test.md", "# Hello World"},
	}

	// Iterate over test data
	for _, data := range testData {
		// Write text
		err := WriteText(data.filepath, data.content)
		if err != nil {
			t.Errorf("Error writing text to file %s: %s", data.filepath, err)
		}
	}
}

// TestReadText tests ReadText function
func TestReadText(t *testing.T) {
	// Test data
	testData := []struct {
		filepath string
		expected string
	}{
		{"test.txt", "Hello World"},
		{"test.md", "# Hello World"},
	}

	// Iterate over test data
	for _, data := range testData {
		// Read text
		content, err := ReadText(data.filepath)
		if err != nil {
			t.Errorf("Error reading text from file %s: %s", data.filepath, err)
		}

		// Compare content
		if content != data.expected {
			t.Errorf("Expected content %s, got %s", data.expected, content)
		}
	}
}

// TestDeleteFile tests DeleteFile function
func TestDeleteFile(t *testing.T) {
	// Test data
	testData := []struct {
		filepath string
	}{
		{"test.txt"},
		{"test.md"},
	}

	// Iterate over test data
	for _, data := range testData {
		// Delete file
		err := DeleteFile(data.filepath)
		if err != nil {
			t.Errorf("Error deleting file %s: %s", data.filepath, err)
		}
	}
}

// TestIsAbsPath tests IsAbsPath function
func TestIsAbsPath(t *testing.T) {
	// Test data
	testData := []struct {
		filepath string
		expected bool
	}{
		{"test.txt", false},
		{"test.md", false},
		{"test", false},
		{"to/test.txt", false},
		{"./path/to/test.txt", false},
		{"/path/to/test.txt", true},
		{"/path/to/test.txt", true},
		{"/path/to/test.md", true},
	}

	// Iterate over test data
	for _, data := range testData {
		// Check if filepath is an absolute path
		isAbsPath := IsAbsPath(data.filepath)

		// Compare isAbsPath
		if isAbsPath != data.expected {
			t.Errorf("Expected isAbsPath %t, got %t for path %s", data.expected, isAbsPath, data.filepath)
		}
	}
}

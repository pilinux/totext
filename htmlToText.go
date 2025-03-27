package totext

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// ConvertHTMLToText receives HTML filepath as an argument and returns its text content and metadata
func ConvertHTMLToText(filepath string, skipPrettifyError bool) (content string, metadata map[string]string, err error) {
	// Prettify the HTML file
	err = PrettifyHTML(filepath)
	if !skipPrettifyError && err != nil {
		return "", nil, err
	}

	// Get the HTML file
	htmlFile, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer func() {
		if e := htmlFile.Close(); e != nil && err == nil {
			err = e
		}
	}()

	// Copy the HTML content into a buffer
	var htmlBuffer bytes.Buffer
	_, err = io.Copy(&htmlBuffer, htmlFile)
	if err != nil {
		return "", nil, err
	}

	// Initialize metadata map
	metadata = make(map[string]string)

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(&htmlBuffer)
	if err != nil {
		return "", nil, err
	}

	// Find and extract the page title if any
	pageTitle := strings.TrimSpace(doc.Find("title").Text())
	if pageTitle != "" {
		metadata["title"] = pageTitle
	}

	// Find and extract the meta description if any
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		name = strings.TrimSpace(name)
		content, _ := s.Attr("content")
		content = strings.TrimSpace(content)
		if name == "description" {
			if content != "" {
				metadata["description"] = content
			}
		}
	})

	// Reset the file reader to the beginning of the file
	_, err = htmlFile.Seek(0, io.SeekStart)
	if err != nil {
		return "", nil, err
	}

	// Initialize a buffer to collect text content
	var textBuffer bytes.Buffer

	// Parse the HTML content
	tokenizer := html.NewTokenizer(htmlFile)

	// Keep track of whether we're inside a <style> or <script> or <footer> element
	insideStyle := false
	insideScript := false
	insideFooter := false

	// Iterate through the HTML tokens
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			// End of the document
			if tokenizer.Err() == io.EOF {
				// Convert the buffer to a string and return it
				content = textBuffer.String()

				// Clean up the HTML content
				content = CleanUpHTML(content)

				// Return the text content and metadata
				return
			}
			// Error occurred while parsing the HTML
			err = fmt.Errorf("error parsing HTML: %v", tokenizer.Err())
			return

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "style" {
				// If we encounter a <style> tag, set insideStyle to true
				insideStyle = true
			}
			if token.Data == "script" {
				// If we encounter a <script> tag, set insideScript to true
				insideScript = true
			}
			if token.Data == "footer" {
				// If we encounter a <footer> tag, set insideFooter to true
				insideFooter = true
			}

		case html.TextToken:
			if !insideStyle && !insideScript && !insideFooter {
				// Extract the text content from the HTML and write it to the buffer
				_, err = textBuffer.WriteString(tokenizer.Token().Data)
				if err != nil {
					return "", nil, err
				}
			}

		case html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "style" {
				// If we encounter the closing tag </style>, set insideStyle to false
				insideStyle = false
			}
			if token.Data == "script" {
				// If we encounter the closing tag </script>, set insideScript to false
				insideScript = false
			}
			if token.Data == "footer" {
				// If we encounter the closing tag </footer>, set insideFooter to false
				insideFooter = false
			}
		}
	}
}

// PrettifyHTML prettifies the HTML content using the prettier library
//
// Dependencies:
//
// npm init
//
// npm install --save-dev --save-exact prettier
func PrettifyHTML(filepath string) (err error) {
	// Check if the file exists
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}

	var cmd *exec.Cmd

	// run cmd based on OS
	switch os := runtime.GOOS; os {
	case "windows":
		// Command to execute in PowerShell
		command := "npx prettier --write " + filepath

		// Create a new PowerShell session
		// and execute the command
		cmd = exec.Command("powershell.exe", "-Command", command)

	default:
		// Command to execute in Bash
		command := "source $HOME/.bashrc && npx prettier --write " + filepath
		cmd = exec.Command("/bin/bash", "-c", command)
	}

	// Prettify the HTML file using prettier command
	err = cmd.Run()

	return
}

// CleanUpHTML cleans up the HTML content and extracts the text content
func CleanUpHTML(content string) string {
	// Define a regular expression pattern to match HTML tags
	re := regexp.MustCompile(`</?(?:div|a|img|picture|svg|video|audio|track|source|canvas|map|noscript|iframe)[^>]*>`)
	// Replace all HTML tags with a newline character
	content = re.ReplaceAllString(content, "\n")

	// Define a regular expression pattern to match HTML comments
	re = regexp.MustCompile(`<!--.*?-->`)
	// Replace all HTML comments with a newline character
	content = re.ReplaceAllString(content, "\n")

	// Create a scanner to read the input string line by line
	scanner := bufio.NewScanner(strings.NewReader(content))

	// Initialize a buffer to collect the cleaned lines
	var buffer bytes.Buffer

	// Iterate through the lines
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// If the line is not empty, append it to the slice
		if line != "" {
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return ""
	}

	// Get the cleaned content
	content = buffer.String()

	return content
}

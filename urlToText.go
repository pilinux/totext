package totext

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// ConvertURLToText fetches the HTML page at the URL given and returns its text content and metadata
func ConvertURLToText(inputURL string) (htmlFilename, content string, metadata map[string]string, err error) {
	inputURL = strings.TrimSpace(inputURL)

	// Parse the URL and validate it
	u, err := ParseURLAndValidate(inputURL)
	if err != nil {
		return
	}

	// Capture the HTML page
	htmlContent, err := CaptureHTML(inputURL)
	if err != nil {
		return
	}

	// Create a filename for the HTML file
	htmlFilename = CreateHTMLFilename(u)

	// Write the HTML content to a file
	err = WriteText(htmlFilename, htmlContent)
	if err != nil {
		return
	}

	// Convert the HTML file to text
	content, metadata, err = ConvertHTMLToText(htmlFilename)

	return
}

// IsHostnameValid validates the hostname
func IsHostnameValid(hostname string) bool {
	// Perform a DNS lookup
	_, err := net.LookupHost(hostname)
	return err == nil
}

// ParseURLAndValidate parses the URL and validates it
func ParseURLAndValidate(inputURL string) (*url.URL, error) {
	// Parse the URL
	u, err := url.Parse(inputURL)
	if err != nil {
		return nil, err
	}

	// Check if the URL has a valid scheme (http or https)
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("invalid scheme")
	}

	// Check if the URL has a valid hostname
	if !IsHostnameValid(u.Hostname()) {
		return nil, fmt.Errorf("invalid hostname")
	}

	return u, nil
}

// CaptureHTML fetches the HTML page at the URL given and
// returns the complete HTML content
func CaptureHTML(inputURL string) (content string, err error) {
	// Create a new browser instance
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Create a new page and navigate to the URL
	page := browser.MustPage(inputURL)
	defer page.MustClose()

	// Set a timeout of 30 seconds and wait for the page to load
	page.Timeout(30 * time.Second).MustWaitLoad()

	// Start to analyze request events
	wait := page.MustWaitRequestIdle()

	// Wait until the page is idle
	wait()

	// Get the HTML content
	content, err = page.HTML()

	return
}

// CreateHTMLFilename generates a filename for the HTML file from the URL
func CreateHTMLFilename(u *url.URL) string {
	// Get the hostname
	hostname := u.Hostname()

	// Get the path
	path := u.Path

	// Set the filename
	filename := hostname + path
	filename = strings.ToLower(filename)

	// Replace the slashes with underscores
	filename = strings.ReplaceAll(filename, "/", "_")

	// Add timestamp to the filename
	filename = fmt.Sprintf("%s_%d", filename, time.Now().UTC().Unix())

	// Add the .html extension
	filename = filename + ".html"

	return filename
}

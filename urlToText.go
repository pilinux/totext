package totext

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

// ConvertURLToText fetches the HTML page at the URL given and returns its text content and metadata
func ConvertURLToText(browser *rod.Browser, inputURL string, skipPrettifyError bool) (htmlFilename, content string, metadata map[string]string, err error) {
	inputURL = strings.TrimSpace(inputURL)

	// Parse the URL and validate it
	u, err := ParseURLAndValidate(inputURL)
	if err != nil {
		return
	}

	// Capture the HTML page
	htmlContent, err := CaptureHTML(browser, inputURL)
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
	content, metadata, err = ConvertHTMLToText(htmlFilename, skipPrettifyError)
	if err != nil {
		return "", "", nil, err
	}

	// Filter out non-readable characters
	content = FilterNonReadableCharacter(content)

	return
}

// IsHostnameValid validates the hostname
func IsHostnameValid(hostname string) bool {
	// Perform a DNS lookup
	_, err := net.LookupHost(hostname)
	return err == nil
}

// IsContentTypeHTML checks if the content type is HTML
func IsContentTypeHTML(contentType string) bool {
	contentType = strings.ToLower(contentType)
	return strings.Contains(contentType, "text/html")
}

// ParseURLAndValidate parses the URL and validates
// the scheme, hostname and content type
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

	// Create an HTTP client with a timeout of 15 seconds
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// Make an HTTP HEAD request to check the content type
	resp, err := client.Head(inputURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the content type is HTML
	if !IsContentTypeHTML(resp.Header.Get("Content-Type")) {
		return nil, fmt.Errorf("invalid content type")
	}

	return u, nil
}

// CaptureHTML fetches the HTML page at the URL given and
// returns the complete HTML content
func CaptureHTML(browser *rod.Browser, inputURL string) (content string, err error) {
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

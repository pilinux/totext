package totext

import "testing"

// TestIsHostnameValid tests IsHostnameValid function
func TestIsHostnameValid(t *testing.T) {
	// Test data
	testData := []struct {
		domain   string
		expected bool
	}{
		{"cloudflare.com", true},
		{"google.com", true},
		{"goapi.pilinux.me", true},
		{"invalid.com", false},
	}

	// Iterate over test data
	for _, data := range testData {
		// Check if domain is valid
		isValid := IsHostnameValid(data.domain)

		// Compare isValid
		if isValid != data.expected {
			t.Errorf("Expected isValid %t, got %t for domain %s", data.expected, isValid, data.domain)
		}
	}
}

package totext

import "unicode"

// FilterNonReadableCharacter - filter out non-readable characters
func FilterNonReadableCharacter(input string) string {
	var cleanedContent []rune
	var lastChar rune

	for _, char := range input {
		if unicode.IsPrint(char) || char == '\n' {
			if char == '\n' && lastChar == '\n' {
				// skip consecutive newlines
				continue
			}
			cleanedContent = append(cleanedContent, char)
			lastChar = char
		}
	}

	return string(cleanedContent)
}

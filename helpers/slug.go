package helpers

import (
	"strings"
	"unicode"
)

func GenerateSlug(name string) string {
	var slugBuilder strings.Builder
	for _, r := range strings.ToLower(name) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			slugBuilder.WriteRune(r)
		} else if r == ' ' || r == '-' {
			slugBuilder.WriteRune('-')
		}
	}
	return strings.Trim(slugBuilder.String(), "-")
}

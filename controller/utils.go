package controller

import "regexp"

const (
	// Regex to remove all 'x-' or 'X-' HTTP headers (Ex: 'X-Forwarded-For')
	headersToRemoveRegex = "(^((forwarded|Forwarded).*)|((x|X)-(.*))$)"
)

func removeCustomeHeaders(h map[string]string) map[string]string {
	for header := range h {
		matched, _ := regexp.MatchString(headersToRemoveRegex, header)
		if matched {
			delete(h, header)
		}
	}

	return h
}
package controller

import "regexp"

var (
	// Regex to remove all 'x-' or 'X-' HTTP headers (Ex: 'X-Forwarded-For' or 'X-Forwarded-Proto')
	headersToRemoveRegex = regexp.MustCompile("(^((forwarded|Forwarded).*)|((x|X)-(.*))$)") 
)

func removeCustomHeaders(h map[string]string) map[string]string {

	for header := range h {
		matched := headersToRemoveRegex.MatchString(header)
		if matched {
			delete(h, header)
		}
	}

	return h
}
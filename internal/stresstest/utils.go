package stresstest

import "regexp"

func ValidateHeaders(headers []string) bool {
	rx, _ := regexp.Compile("^[^:]+:[^:]+$")
	for _, header := range headers {
		if !rx.Match([]byte(header)) {
			return false
		}
	}
	return true
}

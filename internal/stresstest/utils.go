package stresstest

import "regexp"

// ValidateHeaders verifica se os cabeçalhos estão no formato NAME:VALUE
func ValidateHeaders(headers []string) bool {
	rx, _ := regexp.Compile("^[^:]+:[^:]+$")
	for _, header := range headers {
		if !rx.Match([]byte(header)) {
			return false
		}
	}
	return true
}

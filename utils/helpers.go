package utils

import "strings"

func SplitAndTrim(line, d string) []string {
	tokens := strings.Split(line, d)
	for i := range tokens {
		tokens[i] = strings.TrimSuffix(tokens[i], "|")
	}

	return tokens
}

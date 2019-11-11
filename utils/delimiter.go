package utils

import (
	"strings"
)

type EncodingChars struct {
	encodingChars []string
}

func NewEncodingChars(delimiter string) *EncodingChars {
	delimiters := strings.SplitAfter(delimiter, "")
	return &EncodingChars{delimiters}
}

func (e *EncodingChars) GetDelimiters() []string {
	return e.encodingChars
}

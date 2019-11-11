package parser

import (
	"fmt"
	"hl7/utils"
	"io/ioutil"
	"strings"
)

type mainParser struct {
	fileContents string
}

func (m *mainParser) ParseFile() {
	lines := strings.SplitAfter(m.fileContents, "\n")

	var delimiter *utils.EncodingChars

	for _, v := range lines {
		subStr := v[:3]

		switch subStr {
		case "MSH":
			delimiter = utils.NewEncodingChars(v[4:8])
			ParseMSH(v, delimiter)
		}
	}
}

func NewMainParser(filePath string) *mainParser {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	return &mainParser{string(data)}
}

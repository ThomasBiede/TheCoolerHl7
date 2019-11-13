package parser

import (
	"fmt"
	"hl7/segments"
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
			msh := segments.ParseMSH(v, delimiter)
			fmt.Println(msh)

		case "EVN":
			evn := segments.ParseEVN(v, delimiter)
			fmt.Print(evn)

		case "DG1":
			dg1 := segments.ParseDG1(v, delimiter)
			fmt.Println(dg1)

		case "NK1":
			nk1 := segments.ParseNK1(v, delimiter)
			fmt.Println(nk1)

		case "PD1":
			pd1 := segments.ParsePD1(v, delimiter)
			fmt.Println(pd1)

		case "PID":
			pid := segments.ParsePID(v, delimiter)
			fmt.Println(pid)

		case "PV1":
			pv1 := segments.ParsePV1(v, delimiter)
			fmt.Println(pv1)
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

package segments

import "hl7/utils"

type PD1 struct {
	Line string
}

func ParsePD1(line string, encodingChars *utils.EncodingChars) *PD1 {
	return &PD1{line}
}

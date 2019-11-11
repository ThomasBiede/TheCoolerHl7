package parser

import (
	"hl7/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type MSH struct {
	fieldSeparator                      string
	encodingChars                       string
	sendingApplication                  string
	sendingFacility                     string
	receivingApplication                string
	receivingFacility                   string
	dateTimeOfMessage                   time.Time
	security                            string
	messageType                         string
	messageControlID                    string
	processingID                        string
	versionID                           string
	sequenceNumber                      string
	continuationPointer                 string
	acceptAcknowledgmentType            string
	applicationAcknowledgmentType       string
	countryCode                         string
	characterSet                        string //[]string
	principalLanguageOfMessage          string
	alternateCharacterSetHandlingScheme string
	messageProfileIdentifier            string //[]string
}

func ParseMSH(line string, encodingChars *utils.EncodingChars) *MSH {
	// r, _ := regexp.Compile(`([^\|])`)
	// tokens := r.FindAllString(line, -1)
	msh := MSH{}
	tokens := strings.Split(line, "|")

	for i := range tokens {
		tokens[i] = strings.TrimSuffix(tokens[i], "|")
	}

	s, err := strconv.ParseInt(tokens[12], 10, 64)
	if err != nil {
		s = 0
	}

	o := reflect.ValueOf(&msh).Elem()
	for i := 0; i < len(tokens); i++ {
		f := o.Field(i)

		switch f.Type().Kind() {
		case reflect.String:
			f.SetString(tokens[i])
		case reflect.TypeOf(new(*time.Time)).Kind():
			t, _ := time.Parse("20060102150405", tokens[i])
			f.Set(reflect.ValueOf(t))
		}

	}

	return &msh
}

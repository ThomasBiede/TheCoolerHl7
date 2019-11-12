package parser

import (
	"hl7/utils"
	"reflect"
	"strings"
	"time"
)

type MSH struct {
	FieldSeparator                      string
	EncodingChars                       string
	SendingApplication                  string
	SendingFacility                     string
	ReceivingApplication                string
	ReceivingFacility                   string
	DateTimeOfMessage                   *time.Time
	Security                            string
	MessageType                         string
	MessageControlID                    string
	ProcessingID                        string
	VersionID                           string
	SequenceNumber                      string
	ContinuationPointer                 string
	AcceptAcknowledgmentType            string
	ApplicationAcknowledgmentType       string
	CountryCode                         string
	CharacterSet                        []string
	PrincipalLanguageOfMessage          string
	AlternateCharacterSetHandlingScheme string
	MessageProfileIdentifier            []string
}

func ParseMSH(line string, encodingChars *utils.EncodingChars) *MSH {
	// r, _ := regexp.Compile(`([^\|])`)
	// tokens := r.FindAllString(line, -1)
	// s, err := strconv.ParseInt(tokens[12], 10, 64)
	// if err != nil {
	// 	s = 0
	// }
	msh := MSH{}
	tokens := strings.Split(line, "|")

	for i := range tokens {
		tokens[i] = strings.TrimSuffix(tokens[i], "|")
	}

	o := reflect.ValueOf(&msh).Elem()
	for i := 0; i < len(tokens); i++ {
		f := o.Field(i)

		switch f.Type().Kind() {

		case reflect.String:
			f.SetString(tokens[i])

		case reflect.TypeOf(new(time.Time)).Kind():
			formatStr := "20060102150405"
			t, _ := time.Parse(formatStr, tokens[i])
			field := reflect.New(reflect.TypeOf(t))
			field.Elem().Set(reflect.ValueOf(t))
			reflect.ValueOf(&msh).Elem().Field(i).Set(field)

		case reflect.TypeOf(new([]string)).Kind():
			d := encodingChars.GetDelimiters()[1]
			subTokens := strings.Split(tokens[i], d)
			for i := range subTokens {
				subTokens[i] = strings.TrimSuffix(subTokens[i], d)
			}
			f.Set(reflect.ValueOf(subTokens))
		}

	}

	return &msh
}

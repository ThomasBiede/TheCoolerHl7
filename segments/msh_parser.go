package segments

import (
	"hl7/utils"
	"reflect"
	"strconv"
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
	SequenceNumber                      int64
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
	msh := MSH{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&msh).Elem()
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) > 0 {
			f := o.Field(i)

			switch f.Type().Kind() {

			case reflect.String:
				f.SetString(tokens[i])

			case reflect.Int64:
				v, _ := strconv.ParseInt(tokens[i], 10, 64)
				f.SetInt(v)

			case reflect.TypeOf(new(time.Time)).Kind():
				formatStr := "20060102150405"
				t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
				field := reflect.New(reflect.TypeOf(t))
				field.Elem().Set(reflect.ValueOf(t))
				reflect.ValueOf(&msh).Elem().Field(i).Set(field)

			case reflect.Slice:
				d := encodingChars.GetDelimiters()[1]
				if strings.Contains(tokens[i], d) {
					subTokens := strings.Split(tokens[i], d)
					for i := range subTokens {
						subTokens[i] = strings.TrimSuffix(subTokens[i], d)
					}
					f.Set(reflect.ValueOf(subTokens))
				} else {
					f.Set(reflect.ValueOf([]string{tokens[i]}))
				}
			}
		}
	}

	return &msh
}

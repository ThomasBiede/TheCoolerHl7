package datatypes

import (
	"hl7/utils"
	"reflect"
	"time"
)

type XPN struct {
	FamilyName                                  string
	GivenName                                   string
	SecondAndFurtherGivenNamesOrInitialsThereof string
	SuffixegJrOrIii                             string
	PrefixegDr                                  string
	DegreeegMd                                  string
	NameTypeCode                                string
	NameRepresentationCode                      string
	NameContext                                 string
	NameValidityRange                           string
	NameAssemblyOrder                           string
	EffectiveDate                               time.Time
	ExpirationDate                              time.Time
	ProfessionalSuffix                          string
}

func ParseXPN(line string, encodingChars *utils.EncodingChars) *XPN {
	xpn := XPN{}

	tokens := utils.SplitAndTrim(line, encodingChars.GetDelimiters()[0])

	o := reflect.ValueOf(&xpn).Elem()
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) > 0 {
			f := o.Field(i)

			switch f.Type() {

			case reflect.TypeOf(""):
				f.SetString(tokens[i])

			case reflect.TypeOf(time.Time{}):
				formatStr := "20060102150405"
				t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
				f.Set(reflect.ValueOf(t))
			}
		}
	}

	return &xpn
}

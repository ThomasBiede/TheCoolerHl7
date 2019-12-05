package datatypes

import (
	"hl7/utils"
	"reflect"
	"time"
)

type XCN struct {
	IdNumber                                    string
	FamilyName                                  string
	GivenName                                   string
	SecondAndFurtherGivenNamesOrInitialsThereof string
	SuffixegJrOrIii                             string
	PrefixegDr                                  string
	DegreeegMd                                  string
	SourceTable                                 string
	AssigningAuthority                          string
	NameTypeCode                                string
	IdentifierCheckDigit                        string
	CheckDigitScheme                            string
	IdentifierTypeCode                          string
	AssigningFacility                           string
	NameRepresentationCode                      string
	NameContext                                 string
	NameValidityRange                           string
	NameAssemblyOrder                           string
	EffectiveDate                               time.Time
	ExpirationDate                              time.Time
	ProfessionalSuffix                          string
	AssigningJurisdiction                       string
	AssigningAgencyOrDepartment                 string
}

func ParseXCN(line string, encodingChars *utils.EncodingChars) *XCN {
	xcn := XCN{}

	tokens := utils.SplitAndTrim(line, encodingChars.GetDelimiters()[0])

	o := reflect.ValueOf(&xcn).Elem()
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

	return &xcn
}

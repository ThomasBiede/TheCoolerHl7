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
	EffectiveDate                               *time.Time
	ExpirationDate                              *time.Time
	ProfessionalSuffix                          string
	AssigningJurisdiction                       string
	AssigningAgencyOrDepartment                 string
}

func ParseXCN(line string, encodingChars *utils.EncodingChars) *XCN {
	xcn := XCN{}

	tokens := utils.SplitAndTrim(line, encodingChars.GetDelimiters()[0])

	o := reflect.ValueOf(&xcn).Elem()
	for i := 0; i < len(tokens); i++ {
		f := o.Field(i)

		switch f.Type().Kind() {

		case reflect.String:
			f.SetString(tokens[i])

		case reflect.TypeOf(new(time.Time)).Kind():
			formatStr := "20060102150405"
			t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
			field := reflect.New(reflect.TypeOf(t))
			field.Elem().Set(reflect.ValueOf(t))
			reflect.ValueOf(&xcn).Elem().Field(i).Set(field)

		}
	}

	return &xcn
}

package datatypes

import (
	"hl7/utils"
	"reflect"
)

type XTN struct {
	TelephoneNumber                string
	TelecommunicationUseCode       string
	TelecommunicationEquipmentType string
	EmailAddress                   string
	CountryCode                    string
	AreaCityCode                   string
	LocalNumber                    string
	Extension                      string
	AnyText                        string
	ExtensionPrefix                string
	SpeedDialCode                  string
	UnformattedTelephoneNumber     string
}

func ParseXTN(line string, encodingChars *utils.EncodingChars) *XTN {
	xtn := XTN{}

	tokens := utils.SplitAndTrim(line, encodingChars.GetDelimiters()[0])

	o := reflect.ValueOf(&xtn).Elem()
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) > 0 {
			f := o.Field(i)

			f.SetString(tokens[i])
		}
	}

	return &xtn
}

package datatypes

import (
	"hl7/utils"
	"reflect"
	"time"
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

			switch f.Type().Kind() {

			case reflect.String:
				f.SetString(tokens[i])

			case reflect.TypeOf(new(time.Time)).Kind():
				formatStr := "20060102150405"
				t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
				field := reflect.New(reflect.TypeOf(t))
				field.Elem().Set(reflect.ValueOf(t))
				reflect.ValueOf(&xtn).Elem().Field(i).Set(field)
			}
		}
	}

	return &xtn
}

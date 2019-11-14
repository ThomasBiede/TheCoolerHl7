package datatypes

import (
	"hl7/utils"
	"reflect"
	"time"
)

type XAD struct {
	StreetAddress              string
	OtherDesignation           string
	City                       string
	StateOrProvince            string
	ZipOrPostalCode            string
	Country                    string
	AddressType                string
	OtherGeographicDesignation string
	CountyParishCode           string
	CensusTract                string
	AddressRepresentationCode  string
	AddressValidityRange       string
	EffectiveDate              *time.Time
	ExpirationDate             *time.Time
}

func ParseXAD(line string, encodingChars *utils.EncodingChars) *XAD {
	xad := XAD{}

	tokens := utils.SplitAndTrim(line, encodingChars.GetDelimiters()[0])

	o := reflect.ValueOf(&xad).Elem()
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
				reflect.ValueOf(&xad).Elem().Field(i).Set(field)
			}
		}
	}

	return &xad
}

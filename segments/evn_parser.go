package segments

import (
	"hl7/datatypes"
	"hl7/utils"
	"reflect"
	"time"
)

type EVN struct {
	EventTypeCode        string
	RecordedDateTime     time.Time
	DateTimePlannedEvent time.Time
	EventReasonCode      string
	OperatorID           []*datatypes.XCN //XCN
	EventOccurred        time.Time
	EventFacility        string
}

func ParseEVN(line string, encodingChars *utils.EncodingChars) *EVN {
	evn := EVN{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&evn).Elem()
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

			case reflect.TypeOf([]*datatypes.XCN{}):
				var xcns []*datatypes.XCN
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[1])
				for j := 0; j < len(stokens); j++ {
					xcns = append(xcns, datatypes.ParseXCN(stokens[j], encodingChars))
				}

				f.Set(reflect.ValueOf(xcns))
			}
		}
	}

	return &evn
}

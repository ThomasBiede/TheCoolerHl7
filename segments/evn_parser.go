package segments

import (
	"hl7/utils"
	"reflect"
	"strings"
	"time"
)

type EVN struct {
	EventTypeCode        string
	RecordedDateTime     *time.Time
	DateTimePlannedEvent *time.Time
	EventReasonCode      string
	OperatorID           []string //XCN
	EventOccurred        *time.Time
	EventFacility        string
}

func ParseEVN(line string, encodingChars *utils.EncodingChars) *EVN {
	evn := EVN{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&evn).Elem()
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
			reflect.ValueOf(&evn).Elem().Field(i).Set(field)

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

	return &evn
}

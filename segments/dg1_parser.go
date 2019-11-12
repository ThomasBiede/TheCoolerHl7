package segments

import (
	"hl7/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type DG1 struct {
	SetIDDG1                string
	DiagnosisCodingMethod   string
	DiagnosisCodeDG1        string
	DiagnosisDescription    string
	DiagnosisDateTime       *time.Time
	DiagnosisType           string
	MajorDiagnosticCategory string
	DiagnosticRelatedGroup  string
	DRGApprovalIndicator    string
	DRGGrouperReviewCode    string
	OutlierType             string
	OutlierDays             float32
	OutlierCost             float32
	GrouperVersionAndType   string
	DiagnosisPriority       string
	DiagnosingClinician     []string //XCN
	DiagnosisClassification string
	ConfidentialIndicator   string
	AttestationDateTime     *time.Time
	DiagnosisIdentifier     string
	DiagnosisActionCode     string
}

func ParseDG1(line string, encodingChars *utils.EncodingChars) *DG1 {
	dg1 := DG1{}

	tokens := strings.Split(line, "|")
	for i := range tokens {
		tokens[i] = strings.TrimSuffix(tokens[i], "|")
	}

	o := reflect.ValueOf(&dg1).Elem()
	for i := 0; i < len(tokens); i++ {
		f := o.Field(i)

		switch f.Type().Kind() {

		case reflect.String:
			f.SetString(tokens[i])

		case reflect.Float32:
			v, _ := strconv.ParseFloat(tokens[12], 32)
			f.SetFloat(v)

		case reflect.TypeOf(new(time.Time)).Kind():
			formatStr := "20060102150405"
			t, _ := time.Parse(formatStr, tokens[i])
			field := reflect.New(reflect.TypeOf(t))
			field.Elem().Set(reflect.ValueOf(t))
			reflect.ValueOf(&dg1).Elem().Field(i).Set(field)

		case reflect.TypeOf(new([]string)).Kind():
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

	return &dg1
}

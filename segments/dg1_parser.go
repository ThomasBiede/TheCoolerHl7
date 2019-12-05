package segments

import (
	"hl7/datatypes"
	"hl7/utils"
	"reflect"
	"strconv"
	"time"
)

type DG1 struct {
	SetIDDG1                string
	DiagnosisCodingMethod   string
	DiagnosisCodeDG1        string
	DiagnosisDescription    string
	DiagnosisDateTime       time.Time
	DiagnosisType           string
	MajorDiagnosticCategory string
	DiagnosticRelatedGroup  string
	DRGApprovalIndicator    string
	DRGGrouperReviewCode    string
	OutlierType             string
	OutlierDays             float64
	OutlierCost             float64
	GrouperVersionAndType   string
	DiagnosisPriority       string
	DiagnosingClinician     []datatypes.XCN //XCN
	DiagnosisClassification string
	ConfidentialIndicator   string
	AttestationDateTime     time.Time
	DiagnosisIdentifier     string
	DiagnosisActionCode     string
}

func ParseDG1(line string, encodingChars *utils.EncodingChars) *DG1 {
	dg1 := DG1{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&dg1).Elem()
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) > 0 {
			f := o.Field(i)

			switch f.Type() {

			case reflect.TypeOf(""):
				f.SetString(tokens[i])

			case reflect.TypeOf(1.0):
				v, _ := strconv.ParseFloat(tokens[i], 64)
				f.SetFloat(v)

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

	return &dg1
}

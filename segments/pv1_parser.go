package segments

import (
	"hl7/datatypes"
	"hl7/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type PV1 struct {
	SetIDPV1                string
	PatientClass            string
	AssignedPatientLocation string
	AdmissionType           string
	PreadmitNumber          string
	PriorPatientLocation    string
	AttendingDoctor         []*datatypes.XCN // XCN
	ReferringDoctor         []*datatypes.XCN // XCN
	ConsultingDoctor        []*datatypes.XCN // XCN
	HospitalService         string
	TemporaryLocation       string
	PreadmitTestIndicator   string
	ReadmissionIndicator    string
	AdmitSource             string
	AmbulatoryStatus        []string
	VIPIndicator            string
	AdmittingDoctor         []*datatypes.XCN // XCN
	PatientType             string
	VisitNumber             string
	FinancialClass          []string
	ChargePriceIndicator    string
	CourtesyCode            string
	CreditRating            string
	ContractCode            []string
	ContractEffectiveDate   []time.Time
	ContractAmount          []float64
	ContractPeriod          []float64
	InterestCode            string
	TransfertoBadDebtCode   string
	TransfertoBadDebtDate   time.Time
	BadDebtAgencyCode       string
	BadDebtTransferAmount   float64
	BadDebtRecoveryAmount   float64
	DeleteAccountIndicator  string
	DeleteAccountDate       time.Time
	DischargeDisposition    string
	DischargedtoLocation    string
	DietType                string
	ServicingFacility       string
	BedStatus               string
	AccountStatus           string
	PendingLocation         string
	PriorTemporaryLocation  string
	AdmitDateTime           time.Time
	DischargeDateTime       []time.Time
	CurrentPatientBalance   float64
	TotalCharges            float64
	TotalAdjustments        float64
	TotalPayments           float64
	AlternateVisitID        string
	VisitIndicator          string
	OtherHealthcareProvider []datatypes.XCN // XCN
}

func ParsePV1(line string, encodingChars *utils.EncodingChars) *PV1 {
	pv1 := PV1{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&pv1).Elem()

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

			case reflect.TypeOf([]string{}):
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

			case reflect.TypeOf([]*datatypes.XCN{}):
				var xcns []*datatypes.XCN
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[1])
				for j := 0; j < len(stokens); j++ {
					xcns = append(xcns, datatypes.ParseXCN(stokens[j], encodingChars))
				}

				f.Set(reflect.ValueOf(xcns))

			case reflect.TypeOf([]time.Time{}):
				var times []time.Time
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[2])
				for j := 0; j < len(stokens); j++ {
					formatStr := "20060102150405"
					t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
					times = append(times, t)
				}
				f.Set(reflect.ValueOf(times))

			case reflect.TypeOf([]float64{}):
				var floats []float64
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[1])
				for j := 0; j < len(stokens); j++ {
					v, _ := strconv.ParseFloat(tokens[i], 64)
					floats = append(floats, v)
				}
			}
		}
	}

	return &pv1
}

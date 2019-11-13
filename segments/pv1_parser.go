package segments

import (
	"fmt"
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
	AttendingDoctor         []string // XCN
	ReferringDoctor         []string // XCN
	ConsultingDoctor        []string // XCN
	HospitalService         string
	TemporaryLocation       string
	PreadmitTestIndicator   string
	ReadmissionIndicator    string
	AdmitSource             string
	AmbulatoryStatus        []string
	VIPIndicator            string
	AdmittingDoctor         []string // XCN
	PatientType             string
	VisitNumber             string
	FinancialClass          []string
	ChargePriceIndicator    string
	CourtesyCode            string
	CreditRating            string
	ContractCode            []string
	ContractEffectiveDate   []*time.Time
	ContractAmount          []float32
	ContractPeriod          []float32
	InterestCode            string
	TransfertoBadDebtCode   string
	TransfertoBadDebtDate   *time.Time
	BadDebtAgencyCode       string
	BadDebtTransferAmount   float32
	BadDebtRecoveryAmount   float32
	DeleteAccountIndicator  string
	DeleteAccountDate       *time.Time
	DischargeDisposition    string
	DischargedtoLocation    string
	DietType                string
	ServicingFacility       string
	BedStatus               string
	AccountStatus           string
	PendingLocation         string
	PriorTemporaryLocation  string
	AdmitDateTime           *time.Time
	DischargeDateTime       []*time.Time
	CurrentPatientBalance   float32
	TotalCharges            float32
	TotalAdjustments        float32
	TotalPayments           float32
	AlternateVisitID        string
	VisitIndicator          string
	OtherHealthcareProvider []string // XCN
}

func ParsePV1(line string, encodingChars *utils.EncodingChars) *PV1 {
	pv1 := PV1{}

	tokens := strings.Split(line, "|")
	for i := range tokens {
		tokens[i] = strings.TrimSuffix(tokens[i], "|")
	}
	o := reflect.ValueOf(&pv1).Elem()

	for i := 0; i < o.NumField(); i++ {
		fmt.Println(o.Field(i).Type())
	}

	for i := 0; i < len(tokens); i++ {
		f := o.Field(i)

		switch f.Type().Kind() {

		case reflect.String:
			f.SetString(tokens[i])

		case reflect.Float32:
			v, _ := strconv.ParseFloat(tokens[i], 32)
			f.SetFloat(v)

		case reflect.TypeOf(new(time.Time)).Kind():
			formatStr := "20060102150405"
			t, _ := time.Parse(formatStr, tokens[i])
			field := reflect.New(reflect.TypeOf(t))
			field.Elem().Set(reflect.ValueOf(t))
			reflect.ValueOf(&pv1).Elem().Field(i).Set(field)

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

	return &pv1
}

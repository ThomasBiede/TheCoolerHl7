package segments

import (
	"hl7/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type PID struct {
	SetIDPID                    string
	PatientID                   string
	PatientIdentifierList       []string
	AlternatePatientIDPID       []string
	PatientName                 []string // XPN
	MothersMaidenName           []string // XPN
	DateTimeofBirth             *time.Time
	AdministrativeSex           string
	PatientAlias                []string //XPN
	Race                        []string
	PatientAddress              []string // XAD
	CountyCode                  string
	PhoneNumberHome             []string //XTN
	PhoneNumberBusiness         []string //XTN
	PrimaryLanguage             string
	MaritalStatus               string
	Religion                    string
	PatientAccountNumber        string
	SSNNumberPatient            string
	DriversLicenseNumberPatient int64
	MothersIdentifier           []string
	EthnicGroup                 []string
	BirthPlace                  string
	MultipleBirthIndicator      string
	BirthOrder                  float32
	Citizenship                 []string
	VeteransMilitaryStatus      string
	Nationality                 string
	PatientDeathDateandTime     *time.Time
	PatientDeathIndicator       string
	IdentityUnknownIndicator    string
	IdentityReliabilityCode     []string
	LastUpdateDateTime          *time.Time
	LastUpdateFacility          string
	SpeciesCode                 string
	BreedCode                   string
	Strain                      string
	ProductionClassCode         string
	TribalCitizenship           []string
}

func ParsePID(line string, encodingChars *utils.EncodingChars) *PID {
	pid := PID{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&pid).Elem()
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) > 0 {
			f := o.Field(i)

			switch f.Type().Kind() {

			case reflect.String:
				f.SetString(tokens[i])

			case reflect.Int64:
				v, _ := strconv.ParseInt(tokens[i], 10, 64)
				f.SetInt(v)

			case reflect.Float32:
				v, _ := strconv.ParseFloat(tokens[i], 32)
				f.SetFloat(v)

			case reflect.TypeOf(new(time.Time)).Kind():
				formatStr := "20060102150405"
				t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
				field := reflect.New(reflect.TypeOf(t))
				field.Elem().Set(reflect.ValueOf(t))
				reflect.ValueOf(&pid).Elem().Field(i).Set(field)

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
	}

	return &pid
}

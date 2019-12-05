package segments

import (
	"hl7/datatypes"
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
	PatientName                 []*datatypes.XPN // XPN
	MothersMaidenName           []*datatypes.XPN // XPN
	DateTimeofBirth             time.Time
	AdministrativeSex           string
	PatientAlias                []*datatypes.XPN //XPN
	Race                        []string
	PatientAddress              []*datatypes.XAD // XAD
	CountyCode                  string
	PhoneNumberHome             []*datatypes.XTN //XTN
	PhoneNumberBusiness         []*datatypes.XTN //XTN
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
	BirthOrder                  float64
	Citizenship                 []string
	VeteransMilitaryStatus      string
	Nationality                 string
	PatientDeathDateandTime     time.Time
	PatientDeathIndicator       string
	IdentityUnknownIndicator    string
	IdentityReliabilityCode     []string
	LastUpdateDateTime          time.Time
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

			switch f.Type() {

			case reflect.TypeOf(""):
				f.SetString(tokens[i])

			case reflect.TypeOf(1):
				v, _ := strconv.ParseInt(tokens[i], 10, 64)
				f.SetInt(v)

			case reflect.TypeOf(1.0):
				v, _ := strconv.ParseFloat(tokens[i], 32)
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

			case reflect.TypeOf([]*datatypes.XPN{}):
				var xpns []*datatypes.XPN
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[1])
				for j := 0; j < len(stokens); j++ {
					xpns = append(xpns, datatypes.ParseXPN(stokens[j], encodingChars))
				}

				f.Set(reflect.ValueOf(xpns))

			case reflect.TypeOf([]*datatypes.XAD{}):
				var xads []*datatypes.XAD
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[1])
				for j := 0; j < len(stokens); j++ {
					xads = append(xads, datatypes.ParseXAD(stokens[j], encodingChars))
				}

				f.Set(reflect.ValueOf(xads))

			case reflect.TypeOf([]*datatypes.XTN{}):
				var xtns []*datatypes.XTN
				stokens := utils.SplitAndTrim(tokens[i], encodingChars.GetDelimiters()[1])
				for j := 0; j < len(stokens); j++ {
					xtns = append(xtns, datatypes.ParseXTN(stokens[j], encodingChars))
				}

				f.Set(reflect.ValueOf(xtns))
			}
		}
	}

	return &pid
}

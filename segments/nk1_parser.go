package segments

import (
	"hl7/datatypes"
	"hl7/utils"
	"reflect"
	"strings"
	"time"
)

type NK1 struct {
	SetIDNK1                                 string
	NkName                                   []*datatypes.XPN //XPN
	Relationship                             string
	Address                                  []*datatypes.XAD //XAD
	PhoneNumber                              []*datatypes.XTN //XTN
	BusinessPhoneNumber                      []*datatypes.XTN //XTN
	ContactRole                              string
	StartDate                                time.Time
	EndDate                                  time.Time
	NextofKinAssociatedPartiesJobTitle       string
	NextofKinAssociatedPartiesJobCodeClass   string
	NextofKinAssociatedPartiesEmployeeNumber string
	OrganizationNameNK1                      []string
	MaritalStatus                            string
	AdministrativeSex                        string
	DateTimeofBirth                          time.Time
	LivingDependency                         []string
	AmbulatoryStatus                         []string
	Citizenship                              []string
	PrimaryLanguage                          string
	LivingArrangement                        string
	PublicityCode                            string
	ProtectionIndicator                      string
	StudentIndicator                         string
	Religion                                 string
	MothersMaidenName                        []*datatypes.XPN //XPN
	Nationality                              string
	EthnicGroup                              []string
	ContactReason                            []string
	ContactPersonsName                       []*datatypes.XPN //XPN
	ContactPersonsTelephoneNumber            []*datatypes.XTN //XTN
	ContactPersonsAddress                    []*datatypes.XAD //XAD
	NextofKinAssociatedPartysIdentifiers     []string
	JobStatus                                string
	Race                                     []string
	Handicap                                 string
	ContactPersonSocialSecurityNumber        string
	NextofKinBirthPlace                      string
	VIPIndicator                             string
}

func ParseNK1(line string, encodingChars *utils.EncodingChars) *NK1 {
	nk1 := NK1{}

	tokens := utils.SplitAndTrim(line, "|")

	o := reflect.ValueOf(&nk1).Elem()
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

			case reflect.TypeOf([]string{}):
				d := encodingChars.GetDelimiters()[1]
				if strings.Contains(tokens[i], d) {
					subTokens := strings.Split(tokens[i], d)
					for i := range subTokens {
						subTokens[i] = strings.TrimSuffix(subTokens[i], d)
					}
					f.Set(reflect.ValueOf(subTokens))
				} else {
					s := []string{tokens[i]}
					f.Set(reflect.ValueOf(s))
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

	return &nk1
}

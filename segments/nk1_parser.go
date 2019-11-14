package segments

import (
	"hl7/utils"
	"reflect"
	"strings"
	"time"
)

type NK1 struct {
	SetIDNK1                                 string
	NkName                                   []string //XPN
	Relationship                             string
	Address                                  []string //XAD
	PhoneNumber                              []string //XTN
	BusinessPhoneNumber                      []string //XTN
	ContactRole                              string
	StartDate                                *time.Time
	EndDate                                  *time.Time
	NextofKinAssociatedPartiesJobTitle       string
	NextofKinAssociatedPartiesJobCodeClass   string
	NextofKinAssociatedPartiesEmployeeNumber string
	OrganizationNameNK1                      []string
	MaritalStatus                            string
	AdministrativeSex                        string
	DateTimeofBirth                          *time.Time
	LivingDependency                         []string
	AmbulatoryStatus                         []string
	Citizenship                              []string
	PrimaryLanguage                          string
	LivingArrangement                        string
	PublicityCode                            string
	ProtectionIndicator                      string
	StudentIndicator                         string
	Religion                                 string
	MothersMaidenName                        []string //XPN
	Nationality                              string
	EthnicGroup                              []string
	ContactReason                            []string
	ContactPersonsName                       []string //XPN
	ContactPersonsTelephoneNumber            []string //XTN
	ContactPersonsAddress                    []string //XAD
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

			switch f.Type().Kind() {

			case reflect.String:
				f.SetString(tokens[i])

			case reflect.TypeOf(new(time.Time)).Kind():
				formatStr := "20060102150405"
				t, _ := time.Parse(formatStr[0:len(tokens[i])], tokens[i])
				field := reflect.New(reflect.TypeOf(t))
				field.Elem().Set(reflect.ValueOf(t))
				reflect.ValueOf(&nk1).Elem().Field(i).Set(field)

			case reflect.Slice:
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
			}
		}
	}

	return &nk1
}

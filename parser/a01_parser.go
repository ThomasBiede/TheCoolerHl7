package parser

import (
	"fmt"
	"hl7/segments"
	"hl7/utils"
	"io/ioutil"
	"strings"
)

type ADT_A01 struct {
	MessageHeader                               *segments.MSH
	SoftwareSegment                             []string
	EventType                                   *segments.EVN
	PatientIdentification                       *segments.PID
	PatientAdditionalDemographic                *segments.PD1
	Role                                        []string
	NextofKinAssociatedParties                  []*segments.NK1
	PatientVisit                                *segments.PV1
	PatientVisitAdditionalInformation           string
	Role1                                       []string
	Disability                                  []string
	ObservationResult                           []string
	PatientAllergyInformation                   []string
	Diagnosis                                   []*segments.DG1
	DiagnosisRelatedGroup                       string
	Procedures                                  string
	Role2                                       []string
	Guarantor                                   []string
	Insurance                                   string
	InsuranceAdditionalInformation              string
	InsuranceAdditionalInformationCertification []string
	Role3                                       []string
	Accident                                    string
	UB82                                        string
	UB92Data                                    string
	PatientDeathandAutopsy                      string
}

func (a *ADT_A01) ParseFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.SplitAfter(string(data), "\n")
	var delimiter *utils.EncodingChars

	for _, w := range lines {
		subStr := w[:3]
		v := w[4:]

		switch subStr {
		case "MSH":
			delimiter = utils.NewEncodingChars(w[4:8])
			msh := segments.ParseMSH(v, delimiter)
			a.MessageHeader = msh

		case "SFT":
			a.SoftwareSegment = append(a.SoftwareSegment, v)

		case "EVN":
			evn := segments.ParseEVN(v, delimiter)
			a.EventType = evn

		case "DG1":
			dg1 := segments.ParseDG1(v, delimiter)
			a.Diagnosis = append(a.Diagnosis, dg1)

		case "NK1":
			nk1 := segments.ParseNK1(v, delimiter)
			a.NextofKinAssociatedParties = append(a.NextofKinAssociatedParties, nk1)

		case "PD1":
			pd1 := segments.ParsePD1(v, delimiter)
			a.PatientAdditionalDemographic = pd1

		case "PID":
			pid := segments.ParsePID(v, delimiter)
			a.PatientIdentification = pid

		case "PV1":
			pv1 := segments.ParsePV1(v, delimiter)
			a.PatientVisit = pv1

		case "PV2":
			a.PatientVisitAdditionalInformation = v

		case "ROL":
			a.Role = append(a.Role, v)

		case "DB1":
			a.Disability = append(a.Disability, v)

		case "OBX":
			a.ObservationResult = append(a.ObservationResult, v)

		case "AL1":
			a.PatientAllergyInformation = append(a.PatientAllergyInformation, v)

		case "DRG":
			a.DiagnosisRelatedGroup = v

		case "PR1":
			a.Procedures = v

		case "GT1":
			a.Guarantor = append(a.Guarantor, v)

		case "IN1":
			a.Insurance = v

		case "IN2":
			a.InsuranceAdditionalInformation = v

		case "IN3":
			a.InsuranceAdditionalInformationCertification = append(a.InsuranceAdditionalInformationCertification, v)

		case "ACC":
			a.Accident = v

		case "UB1":
			a.UB92Data = v

		case "PDA":
			a.PatientDeathandAutopsy = v
		}
	}
}

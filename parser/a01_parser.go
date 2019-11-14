package parser

import (
	"fmt"
	"hl7/segments"
	"hl7/utils"
	"io/ioutil"
	"strings"
)

type ADT_A01 struct {
	messageHeader                               *segments.MSH
	softwareSegment                             []string
	eventType                                   *segments.EVN
	patientIdentification                       *segments.PID
	patientAdditionalDemographic                *segments.PD1
	role                                        []string
	nextofKinAssociatedParties                  []*segments.NK1
	patientVisit                                *segments.PV1
	patientVisitAdditionalInformation           string
	role1                                       []string
	disability                                  []string
	observationResult                           []string
	patientAllergyInformation                   []string
	diagnosis                                   []*segments.DG1
	diagnosisRelatedGroup                       string
	procedures                                  string
	role2                                       []string
	guarantor                                   []string
	insurance                                   string
	insuranceAdditionalInformation              string
	insuranceAdditionalInformationCertification []string
	role3                                       []string
	accident                                    string
	uB82                                        string
	uB92Data                                    string
	patientDeathandAutopsy                      string
}

func (a *ADT_A01) ParseFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.SplitAfter(string(data), "\n")

	var delimiter *utils.EncodingChars

	for _, v := range lines {
		subStr := v[:3]

		switch subStr {
		case "MSH":
			delimiter = utils.NewEncodingChars(v[4:8])
			msh := segments.ParseMSH(v, delimiter)
			a.messageHeader = msh

		case "SFT":
			a.softwareSegment = append(a.softwareSegment, v)

		case "EVN":
			evn := segments.ParseEVN(v, delimiter)
			a.eventType = evn

		case "DG1":
			dg1 := segments.ParseDG1(v, delimiter)
			fmt.Println(dg1)

		case "NK1":
			nk1 := segments.ParseNK1(v, delimiter)
			a.nextofKinAssociatedParties = append(a.nextofKinAssociatedParties, nk1)

		case "PD1":
			pd1 := segments.ParsePD1(v, delimiter)
			a.patientAdditionalDemographic = pd1

		case "PID":
			pid := segments.ParsePID(v, delimiter)
			a.patientIdentification = pid

		case "PV1":
			pv1 := segments.ParsePV1(v, delimiter)
			a.patientVisit = pv1

		case "PV2":
			a.patientVisitAdditionalInformation = v

		case "ROL":
			a.role = append(a.role, v)
		}
	}
}

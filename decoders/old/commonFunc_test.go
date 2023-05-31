package decoders_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ISEMS-MRSICT/datamodels"

	mstixo "github.com/av-belyakov/methodstixobjects"
)

var _ = Describe("CommonFunc", func() {
	var (
		reportDomainObjectsSTIX = datamodels.ReportDomainObjectsSTIX{
			ReportDomainObjectsSTIX: mstixo.ReportDomainObjectsSTIX{
				CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
					Type: "report",
					ID:   "report--t66e6-2r33h8-fefer442-3r3",
				},
				Name:        "name_logic",
				Description: "vv ieii i viivr ibir",
				ReportTypes: []mstixo.OpenVocabTypeSTIX{
					"attack-pattern",
					"intrusion-set",
					"observed-data",
				},
				ObjectRefs: []mstixo.IdentifierTypeSTIX{
					"attack-pattern--t7277733773r77-f3f37377e377",
					"intrusion-set--wdtw52525e5-f2662fe2e",
					"observed-data--fdeyde2721-113334e3",
				},
			},
		}

		reportTypeOV = []string{
			"attack-pattern",
			"campaign",
			"identity",
			"indicator",
			"intrusion-set",
			"malware",
			"observed-data",
			"threat-actor",
			"threat-report",
			"tool",
			"vulnerability",
		}
	)

	changeReportTypes := func(rdostix datamodels.ReportDomainObjectsSTIX) datamodels.ReportDomainObjectsSTIX {
		newrdostix := datamodels.ReportDomainObjectsSTIX{
			ReportDomainObjectsSTIX: mstixo.ReportDomainObjectsSTIX{
				CommonPropertiesObjectSTIX:       rdostix.CommonPropertiesObjectSTIX,
				CommonPropertiesDomainObjectSTIX: rdostix.CommonPropertiesDomainObjectSTIX,
				Name:                             rdostix.Name,
				Description:                      rdostix.Description,
				Published:                        rdostix.Published,
			},
		}
		newrdostix.ObjectRefs = append(newrdostix.ObjectRefs, rdostix.ObjectRefs...)

		for _, v := range newrdostix.ObjectRefs {
			typeName := strings.Split(string(v), "--")[0]

			for _, tn := range reportTypeOV {
				if typeName == tn {
					var isExist bool
					for _, nrt := range newrdostix.ReportTypes {
						if string(nrt) == typeName {
							isExist = true
						}
					}

					if !isExist {
						newrdostix.ReportTypes = append(newrdostix.ReportTypes, mstixo.OpenVocabTypeSTIX(typeName))
					}
				}
			}
		}

		return newrdostix
	}

	Context("Тест 1. Тестируем изменение значений поля 'report_types' STIX объекта 'Report' при изменении содержимого поля 'object_ref' того же объекта", func() {
		It("При добавлении id SDO объектов, типов которых еще нет в поле 'report_types', в поле 'object_ref' их тип должен добавится в поле 'report_types'", func() {
			reportDomainObjectsSTIX.ObjectRefs = append(reportDomainObjectsSTIX.ObjectRefs, mstixo.IdentifierTypeSTIX("indicator--bywgd737f-g38g83-38f3"))
			rstix := changeReportTypes(reportDomainObjectsSTIX)

			//fmt.Printf("Report SDO, field report_types: '%v\n", rstix.ReportTypes)

			Expect(len(rstix.ReportTypes)).Should(Equal(4))
		})

		It("При добавлении id SDO объекта, типа, который уже есть в поле 'report_types' никаких изменений бть не должно", func() {
			reportDomainObjectsSTIX.ObjectRefs = append(
				reportDomainObjectsSTIX.ObjectRefs,
				[]mstixo.IdentifierTypeSTIX{
					mstixo.IdentifierTypeSTIX("indicator--bywgd737f-g38g83-38f3"),
					mstixo.IdentifierTypeSTIX("observed-data--yfye73-3378883-37f7rgf"),
				}...)
			rstix := changeReportTypes(reportDomainObjectsSTIX)

			//fmt.Printf("Report SDO, field report_types: '%v\n", rstix.ReportTypes)

			Expect(len(rstix.ReportTypes)).Should(Equal(4))
		})

		It("При удалении некоторых id SDO объектов из поля 'object_ref' их тип должен удалится из поля 'report_types'", func() {
			list := []mstixo.IdentifierTypeSTIX{}

			list = append(list, reportDomainObjectsSTIX.ObjectRefs[0])
			list = append(list, reportDomainObjectsSTIX.ObjectRefs[2])

			fmt.Printf("List: '%v'\n", list)

			reportDomainObjectsSTIX.ObjectRefs = list
			rstix := changeReportTypes(reportDomainObjectsSTIX)

			fmt.Printf("Report SDO, field report_types: '%v\n", rstix.ReportTypes)

			Expect(len(rstix.ReportTypes)).Should(Equal(2))
		})

		It("При добавлении определенных id SDO или CDO объектов в поле 'object_ref' их тип НЕ должен добавится в поле 'report_types'", func() {
			reportDomainObjectsSTIX.ObjectRefs = append(
				reportDomainObjectsSTIX.ObjectRefs,
				[]mstixo.IdentifierTypeSTIX{
					mstixo.IdentifierTypeSTIX("note--52624t3-rewewrg82-235544"),
					mstixo.IdentifierTypeSTIX("file--e3f4t4t-44r3r3r3r-444t4t"),
				}...)
			rstix := changeReportTypes(reportDomainObjectsSTIX)

			fmt.Printf("Report SDO, field object_refs: '%v\n", rstix.ObjectRefs)

			Expect(len(rstix.ReportTypes)).Should(Equal(2))
		})

		It("При удалении определенных id SDO или CDO объектов из поля 'object_ref' их тип НЕ должен удалится из поля 'report_types'", func() {

		})
	})
})

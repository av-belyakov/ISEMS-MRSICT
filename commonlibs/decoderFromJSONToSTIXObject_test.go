package commonlibs_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type resultProcessingListSTIXObject struct {
	DataType string
	Data     interface{}
}

type fieldTypeSTIXObject struct {
	Type string `json:"type"`
}

var _ = Describe("DecoderFromJSONToSTIXObject", func() {
	var (
		docJSON                        []byte
		errReadFile                    error
		errUnmarchalReq                error
		errUnmarchalList               error
		modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
		commonPropertiesObjectSTIX     datamodels.CommonPropertiesObjectSTIX
		listSTIXObjectJSON             datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
		listResult                     []*resultProcessingListSTIXObject
		fieldTypeSTIXObject            fieldTypeSTIXObject
	)

	numSTIXObj := map[string]int{}
	numSTIXType := map[string]int{}

	var _ = BeforeSuite(func() {
		docJSON, errReadFile = ioutil.ReadFile("../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		errUnmarchalList = json.Unmarshal(*modAPIRequestProcessingReqJSON.RequestDetails, &listSTIXObjectJSON)

		for _, item := range listSTIXObjectJSON {
			numCurrent := 1

			err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
			if err != nil {
				continue
			}

			resultDecodingSTIXObject, typeSTIXObject, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
			if err != nil {
				fmt.Println(err)

				continue
			}

			listResult = append(listResult, &resultProcessingListSTIXObject{
				DataType: typeSTIXObject,
				Data:     resultDecodingSTIXObject,
			})

			if num, ok := numSTIXObj[typeSTIXObject]; ok {
				numCurrent = numCurrent + num
			}

			numSTIXObj[typeSTIXObject] = numCurrent

			numType := 1
			err = json.Unmarshal(*item, &fieldTypeSTIXObject)
			if err != nil {
				fmt.Println(err)

				continue
			}

			if n, ok := numSTIXType[fieldTypeSTIXObject.Type]; ok {
				numType = numType + n
			}

			numSTIXType[fieldTypeSTIXObject.Type] = numType
		}

		fmt.Println("----------- STIX objects -----------")
		for k, v := range numSTIXObj {
			fmt.Printf("Key: '%s', Value: '%d'\n", k, v)
		}

		listTypeKeys := make([]string, 0, len(numSTIXType))

		fmt.Println("----------- STIX types -----------")
		for k := range numSTIXType {
			listTypeKeys = append(listTypeKeys, k)
		}

		sort.Strings(listTypeKeys)

		for _, k := range listTypeKeys {
			fmt.Printf("Key: '%s', Value: '%d'\n", k, numSTIXType[k])
		}
	})

	Context("Тест 1. Чтение тестового файла", func() {
		It("При чтении файла не должно быть ошибок", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Декодируем все STIX объекты содержащиеся в JSON файле", func() {
		It("При декодировании объекта запроса не долно быть ошибок", func() {
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())
		})

		It("При декодировании списка STIX объектов в тип interface{} не должно быть ошибок", func() {
			Expect(errUnmarchalList).ShouldNot(HaveOccurred())
		})

		It("Должно быть получено определенное количество STIX объектов", func() {
			Expect(len(listResult)).Should(Equal(64))
		})

		It("Должен быть найден 1 объект типа 'relationship'", func() {
			Expect(numSTIXType["relationship"]).Should(Equal(1))
		})

		It("Должен быть найден 3 объекта типа 'location'", func() {
			Expect(numSTIXType["location"]).Should(Equal(3))
		})

		It("Должен быть найден 2 объекта типа 'malware'", func() {
			Expect(numSTIXType["malware"]).Should(Equal(2))
		})

		It("Должен быть найден 3 объекта типа 'email-addr'", func() {
			Expect(numSTIXType["email-addr"]).Should(Equal(3))
		})

		It("Должен быть найден 8 объектов типа 'file'", func() {
			Expect(numSTIXType["file"]).Should(Equal(8))
		})
	})

	/*
			if fieldTypeSTIXObject.Type == "email-message" {
			fmt.Println(resultDecodingSTIXObject)
		}
	*/

	Context("Тест 3. Выполняем приведение типов для объектов STIX", func() {
		It("Должны быть успешно выполненны все приведения типов", func() {
			var (
				err error
				num int
			)

			for k := range listResult {
				if listResult[k].DataType == "domain object stix" {
					switch (listResult[k].Data).(type) {
					case datamodels.AttackPatternDomainObjectsSTIX:
						num++
					case datamodels.CampaignDomainObjectsSTIX:
						num++
					case datamodels.CourseOfActionDomainObjectsSTIX:
						num++
					case datamodels.GroupingDomainObjectsSTIX:
						num++
					case datamodels.IdentityDomainObjectsSTIX:
						num++
					case datamodels.IndicatorDomainObjectsSTIX:
						num++
					case datamodels.InfrastructureDomainObjectsSTIX:
						num++
					case datamodels.IntrusionSetDomainObjectsSTIX:
						num++
					case datamodels.LocationDomainObjectsSTIX:
						num++
					case datamodels.MalwareDomainObjectsSTIX:
						num++
					case datamodels.MalwareAnalysisDomainObjectsSTIX:
						num++
					case datamodels.NoteDomainObjectsSTIX:
						num++
					case datamodels.ObservedDataDomainObjectsSTIX:
						num++
					case datamodels.OpinionDomainObjectsSTIX:
						num++
					case datamodels.ReportDomainObjectsSTIX:
						num++
					case datamodels.ThreatActorDomainObjectsSTIX:
						num++
					case datamodels.ToolDomainObjectsSTIX:
						num++
					case datamodels.VulnerabilityDomainObjectsSTIX:
						num++
					}
				} else if listResult[k].DataType == "cyber observable object stix" {
					switch data := (listResult[k].Data).(type) {
					case datamodels.ArtifactCyberObservableObjectSTIX:
						num++
					case datamodels.AutonomousSystemCyberObservableObjectSTIX:
						num++
					case datamodels.DirectoryCyberObservableObjectSTIX:
						num++
					case datamodels.DomainNameCyberObservableObjectSTIX:
						num++
					case datamodels.EmailAddressCyberObservableObjectSTIX:
						num++
					case datamodels.EmailMessageCyberObservableObjectSTIX:
						fmt.Printf("STIX type object:'%s'\n", data.Type)
						fmt.Println("AdditionalHeaderFields:")

						for k, v := range data.AdditionalHeaderFields {
							fmt.Printf("Key: '%s', Value: '%v'\n", k, v)
						}

						num++
					case datamodels.FileCyberObservableObjectSTIX:
						for k, v := range data.Extensions {
							if k == "pdf-ext" {
								fmt.Printf("STIX type object:'%s'\n", data.Type)
								fmt.Printf("Extensions: '%s'\n", k)

								list, ok := (*v).(map[string]string)
								if !ok {
									continue
								}

								for a, b := range list {
									fmt.Printf("Name: '%s', Value: '%s'\n", a, b)
								}
							}
						}

						num++
					case datamodels.IPv4AddressCyberObservableObjectSTIX:
						num++
					case datamodels.IPv6AddressCyberObservableObjectSTIX:
						num++
					case datamodels.MACAddressCyberObservableObjectSTIX:
						num++
					case datamodels.MutexCyberObservableObjectSTIX:
						num++
					case datamodels.NetworkTrafficCyberObservableObjectSTIX:
						fmt.Printf("STIX type object:'%s'\n", data.Type)
						fmt.Println("Extensions:")

						for k, v := range data.Extensions {
							fmt.Printf("Key: '%s', Value: '%v'\n", k, *v)
						}

						num++
					case datamodels.ProcessCyberObservableObjectSTIX:
						num++
					case datamodels.SoftwareCyberObservableObjectSTIX:
						num++
					case datamodels.URLCyberObservableObjectSTIX:
						num++
					case datamodels.UserAccountCyberObservableObjectSTIX:
						num++
					case datamodels.WindowsRegistryKeyCyberObservableObjectSTIX:
						num++
					case datamodels.X509CertificateCyberObservableObjectSTIX:
						num++
					}
				} else if listResult[k].DataType == "relationship object stix" {
					switch (listResult[k].Data).(type) {
					case datamodels.RelationshipObjectSTIX:
						num++
					case datamodels.SightingObjectSTIX:
						num++
					}
				} else {
					err = fmt.Errorf("Error, type object STIX not found")

					break
				}
			}

			Expect(err).ShouldNot(HaveOccurred())
			Expect(num).Should(Equal(64))
		})
	})
})

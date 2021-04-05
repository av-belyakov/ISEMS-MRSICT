package decoders_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/decoders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecoderFromJSONToSTIXObject", func() {
	var (
		docJSON                        []byte
		errso                          error
		errReadFile                    error
		errUnmarchalReq                error
		errUnmarchalList               error
		modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
		listSTIXObjectJSON             datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
		listSTIXObj                    []*datamodels.ElementSTIXObject
	)

	countSTIXObj := map[string]int{}

	var _ = BeforeSuite(func() {
		docJSON, errReadFile = ioutil.ReadFile("../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		errUnmarchalList = json.Unmarshal(*modAPIRequestProcessingReqJSON.RequestDetails, &listSTIXObjectJSON)

		listSTIXObj, errso = decoders.GetListSTIXObjectFromJSON(listSTIXObjectJSON)

		for _, i := range listSTIXObj {
			n := 1

			if num, ok := countSTIXObj[i.DataType]; ok {
				countSTIXObj[i.DataType] = num + n

				continue
			}

			countSTIXObj[i.DataType] = n
		}
	})

	/*
				if fieldTypeSTIXObject.Type == "email-message" {
				fmt.Println(resultDecodingSTIXObject)
			}

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
		})*/

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
	})

	Context("Тест 3. Декодируем все STIX объекты содержащиеся в JSON файле (выполняем декодирование через методы типов объектов STIX)", func() {
		It("При декодировании объекта запроса не долно быть ошибок", func() {
			Expect(errso).ShouldNot(HaveOccurred())
		})

		It("Должно быть получено определенное количество STIX объектов (64)", func() {
			Expect(len(listSTIXObj)).Should(Equal(64))
		})

		It("Должен быть найден 1 объект типа 'relationship'", func() {
			Expect(countSTIXObj["relationship"]).Should(Equal(1))
		})

		It("Должен быть найден 3 объекта типа 'location'", func() {
			Expect(countSTIXObj["location"]).Should(Equal(3))
		})

		It("Должен быть найден 2 объекта типа 'malware'", func() {
			Expect(countSTIXObj["malware"]).Should(Equal(2))
		})

		It("Должен быть найден 3 объекта типа 'email-addr'", func() {
			Expect(countSTIXObj["email-addr"]).Should(Equal(3))
		})

		It("Должен быть найден 8 объектов типа 'file'", func() {
			Expect(countSTIXObj["file"]).Should(Equal(8))
		})

		Context("Тест 4. Проверяем полученный список объектов STIX на корректно отработанное поле Extensions", func() {
			It("Должен быть найдено 8 STIX объектов типа 'file', 6 из которых содержат заолненное поле Extensions", func() {
				var numFieldExtensions, numObjFile int

				for _, i := range listSTIXObj {
					if i.DataType != "file" {
						continue
					}

					fmt.Printf("==== Type STIX object: '%s' ====\nObject STIX:'%v'\n", i.DataType, i.Data)

					if obj, ok := i.Data.(datamodels.FileCyberObservableObjectSTIX); ok {
						for k, v := range obj.Extensions {
							fmt.Printf("	Extensions name:'%s'\n	Extensions value:'%v'\n", k, v)
						}

						if len(obj.Extensions) > 0 {
							numFieldExtensions++
						}
					}

					numObjFile++
				}

				Expect(numFieldExtensions).Should(Equal(6))
				Expect(numObjFile).Should(Equal(8))
			})
		})

		Context("Тест 5. Проверяем полученный список объектов STIX на наличие спецефичных полей для некоторых типов STIX объектов", func() {
			It("Список объектов STIX не должен быть пустым", func() {

				Expect(len(listSTIXObj)).ShouldNot(Equal(0))

				for _, i := range listSTIXObj {

					fmt.Printf("Test 5. STIX object type: '%s'\n", i.DataType)

					switch i.DataType {
					//  1. Для Domain Objects STIX
					case "attack-pattern":
						obj := i.GetAttackPatternDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "attack-pattern")).Should(BeTrue())
					case "campaign":
						obj := i.GetCampaignDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "campaign")).Should(BeTrue())
					case "course-of-action":
						obj := i.GetCourseOfActionDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "course-of-action")).Should(BeTrue())
					case "grouping":
						obj := i.GetGroupingDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "grouping")).Should(BeTrue())
					case "identity":
						obj := i.GetIdentityDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "identity")).Should(BeTrue())
					case "indicator":
						obj := i.GetIndicatorDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "indicator")).Should(BeTrue())
					case "infrastructure":
						obj := i.GetInfrastructureDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "infrastructure")).Should(BeTrue())
					case "intrusion-set":
						obj := i.GetIntrusionSetDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "intrusion-set")).Should(BeTrue())
					case "location":
						obj := i.GetLocationDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "location")).Should(BeTrue())
					case "malware":
						obj := i.GetMalwareDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "malware")).Should(BeTrue())
					case "malware-analysis":
						obj := i.GetMalwareAnalysisDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "malware-analysis")).Should(BeTrue())
					case "note":
						obj := i.GetNoteDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "note")).Should(BeTrue())
					case "observed-data":
						obj := i.GetObservedDataDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "observed-data")).Should(BeTrue())
					case "opinion":
						obj := i.GetOpinionDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "opinion")).Should(BeTrue())
					case "report":
						obj := i.GetReportDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "report")).Should(BeTrue())
					case "threat-actor":
						obj := i.GetThreatActorDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "threat-actor")).Should(BeTrue())
					case "tool":
						obj := i.GetToolDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "tool")).Should(BeTrue())
					case "vulnerability":
						obj := i.GetVulnerabilityDomainObjectsSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "vulnerability")).Should(BeTrue())

						//  2. Для Relationship Objects STIX
					case "relationship":
						obj := i.GetRelationshipObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "relationship")).Should(BeTrue())
					case "sighting":
						obj := i.GetSightingObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "sighting")).Should(BeTrue())

						//  3. Для Cyber-observable Objects STIX
					case "artifact":
						obj := i.GetArtifactCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "artifact")).Should(BeTrue())
					case "autonomous-system":
						obj := i.GetAutonomousSystemCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "autonomous-system")).Should(BeTrue())
					case "directory":
						obj := i.GetDirectoryCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "directory")).Should(BeTrue())
					case "domain-name":
						obj := i.GetDomainNameCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "domain-name")).Should(BeTrue())
					case "email-addr":
						obj := i.GetEmailAddressCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "email-addr")).Should(BeTrue())
					case "email-message":
						obj := i.GetEmailMessageCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "email-message")).Should(BeTrue())
					case "file":
						obj := i.GetFileCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "file")).Should(BeTrue())
					case "ipv4-addr":
						obj := i.GetIPv4AddressCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "ipv4-addr")).Should(BeTrue())
					case "ipv6-addr":
						obj := i.GetIPv6AddressCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "ipv6-addr")).Should(BeTrue())
					case "mac-addr":
						obj := i.GetMACAddressCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "mac-addr")).Should(BeTrue())
					case "mutex":
						obj := i.GetMutexCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "mutex")).Should(BeTrue())
					case "network-traffic":
						obj := i.GetNetworkTrafficCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "network-traffic")).Should(BeTrue())
					case "process":
						obj := i.GetProcessCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "process")).Should(BeTrue())
					case "software":
						obj := i.GetSoftwareCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "software")).Should(BeTrue())
					case "url":
						obj := i.GetURLCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "url")).Should(BeTrue())
					case "user-account":
						obj := i.GetUserAccountCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "user-account")).Should(BeTrue())
					case "windows-registry-key":
						obj := i.GetWindowsRegistryKeyCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "windows-registry-key")).Should(BeTrue())
					case "x509-certificate":
						obj := i.GetX509CertificateCyberObservableObjectSTIX()

						Expect(obj).ShouldNot(BeNil())
						Expect(strings.Contains(obj.ID, "x509-certificate")).Should(BeTrue())
					}
				}
			})
		})
	})
})

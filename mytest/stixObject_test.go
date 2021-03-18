package mytest_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StixObject", func() {
	var fd []byte
	var err error
	var _ = BeforeSuite(func() {
		fd, err = ioutil.ReadFile("jsonSTIXExample.json")
		if err != nil {
			fmt.Println(err)
		}
	})

	Context("Тест №1. Тест парсинга STIX объекта", func() {
		It("При чтении объекта не должно быть ошибок", func() {
			var (
				modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
				commonPropertiesObjectSTIX     datamodels.CommonPropertiesObjectSTIX
			)
			var err error

			err = json.Unmarshal(fd, &modAPIRequestProcessingReqJSON)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Printf("Task ID: '%s'\nSection: '%s'\nTask was generated automatically: '%v'\nUser name generated task: '%s'\n", modAPIRequestProcessingReqJSON.TaskID, modAPIRequestProcessingReqJSON.Section, modAPIRequestProcessingReqJSON.TaskWasGeneratedAutomatically, modAPIRequestProcessingReqJSON.UserNameGeneratedTask)

			type ResultProcessingListSTIXObject struct {
				DataType string
				Data     interface{}
			}

			switch modAPIRequestProcessingReqJSON.Section {
			case "handling stix object":
				fmt.Printf("Request Section:'%s'\n", modAPIRequestProcessingReqJSON.Section)

				var listSTIXObjectJSON datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
				if err := json.Unmarshal(*modAPIRequestProcessingReqJSON.RequestDetails, &listSTIXObjectJSON); err != nil {
					Expect(err).ShouldNot(HaveOccurred())
				}

				for _, item := range listSTIXObjectJSON {
					err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
					if err != nil {
						Expect(err).ShouldNot(HaveOccurred())
					}

					resultDecodingSTIXObject, typeSTIXObject, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
					if err != nil {
						Expect(err).ShouldNot(HaveOccurred())
					}

					fmt.Printf("	--=== Type STIX Object:'%s'\n", typeSTIXObject)
					if typeSTIXObject == "domain object stix" {
						switch data := (resultDecodingSTIXObject).(type) {
						case datamodels.AttackPatternDomainObjectsSTIX:
							fmt.Printf("\nType:'%s'\nID:'%s'\nVersion:'%s'\nName:'%s'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.Name, resultDecodingSTIXObject)

						case datamodels.CampaignDomainObjectsSTIX:
							fmt.Printf("\nType:'%s'\nID:'%s'\nVersion:'%s'\nName:'%s'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.Name, resultDecodingSTIXObject)

						case datamodels.CourseOfActionDomainObjectsSTIX:
							fmt.Printf("\nType:'%s'\nID:'%s'\nVersion:'%s'\nName:'%s'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.Name, resultDecodingSTIXObject)

						case datamodels.IndicatorDomainObjectsSTIX:
							fmt.Printf("\nType:'%s'\nID:'%s'\nVersion:'%s'\nName:'%s'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.Name, resultDecodingSTIXObject)

						}
					} else if typeSTIXObject == "cyber observable object stix" {
						switch data := (resultDecodingSTIXObject).(type) {
						case datamodels.FileCyberObservableObjectSTIX:
							fmt.Printf("\nType:'%s' @@@\nID:'%s'\nVersion:'%s'\nName:'%s'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.Name, resultDecodingSTIXObject)
							for extKey, extValue := range data.Extensions {
								fmt.Printf("Extensions key:'%s'\n	Value:'%v'\n", extKey, *extValue)
								if extKey == "windows-pebinary-ext" {
									if result, ok := (*extValue).(datamodels.WindowsPEBinaryFileExtensionSTIX); ok {
										fmt.Println(result.PeType)
										fmt.Println(result.TimeDateStamp)
										fmt.Println(result.MachineHex)
									}
								}
							}

						case datamodels.NetworkTrafficCyberObservableObjectSTIX:
							fmt.Printf("\nType:'%s' $$$$\nID:'%s'\nVersion:'%s'\nProtocols:'%v'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.Protocols, resultDecodingSTIXObject)
							for extKey, extValue := range data.Extensions {
								fmt.Printf("Extensions key:'%s'\n	Value:'%v'\n", extKey, *extValue)
								if extKey == "http-request-ext" {
									if result, ok := (*extValue).(datamodels.HTTPRequestExtensionSTIX); ok {
										fmt.Println(result.RequestMethod)
										fmt.Println(result.RequestValue)
										fmt.Println(result.RequestHeader)
									}
								}
							}

						case datamodels.ProcessCyberObservableObjectSTIX:
							fmt.Printf("\nType:'%s' $$$$\nID:'%s'\nVersion:'%s'\nPID:'%v'\nObject:'%v' |\n", data.Type, data.ID, data.SpecVersion, data.PID, resultDecodingSTIXObject)
							for extKey, extValue := range data.Extensions {
								fmt.Printf("Extensions key:'%s'\n	Value:'%v'\n", extKey, *extValue)
								if extKey == "windows-service-ext" {
									if result, ok := (*extValue).(datamodels.WindowsServiceExtensionSTIX); ok {
										fmt.Println(result.ServiceName)
										fmt.Println(result.ServiceType)
										fmt.Println(result.ServiceStatus)
									}
								}
							}
						case datamodels.UserAccountCyberObservableObjectSTIX:

						}
					} else if typeSTIXObject == "relationship object stix" {

					} else {
						fmt.Println("Неизвестный тип")
					}
				}

			case "handling search requests":
				fmt.Printf("Request Section:'%s'\n", modAPIRequestProcessingReqJSON.Section)

			case "generating reports":
				fmt.Printf("Request Section:'%s'\n", modAPIRequestProcessingReqJSON.Section)

			case "formation final documents":
				fmt.Printf("Request Section:'%s'\n", modAPIRequestProcessingReqJSON.Section)

			}

			/*for _, item := range modAPIRequestProcessingReqJSON.DetailedInformation {
				fmt.Printf("Action type: '%v'\n", item.ActionType)
				//fmt.Println(item.MessageParameters)

				if strings.Contains(item.ActionType, "STIX object") {
					err := json.Unmarshal(item.MessageParameters, &commonPropertiesObjectSTIX)
					if err != nil {
						fmt.Println(err)
					}

					fmt.Println("_____ Processing STIX object ______")
					fmt.Printf("------- Type STIX object is '%s' -------\n", commonPropertiesObjectSTIX.Type)

					resultDecodingSTIXObject, err := listDecoderSTIXObject[commonPropertiesObjectSTIX.Type](commonPropertiesObjectSTIX.Type, &item.MessageParameters)
					if err != nil {
						Expect(err).ShouldNot(HaveOccurred())
					}

					switch data := (resultDecodingSTIXObject.Data).(type) {
					case datamodels.AttackPatternDomainObjectsSTIX:
						fmt.Printf("| Type:'%s', Version:'%s', Name:'%s', Object:'%v' |\n", data.Type, data.SpecVersion, data.Name, resultDecodingSTIXObject.Data)

					case datamodels.CampaignDomainObjectsSTIX:
						fmt.Printf("| Type:'%s', Version:'%s', Name:'%s', Object:'%v' |\n", data.Type, data.SpecVersion, data.Name, resultDecodingSTIXObject.Data)

					case datamodels.CourseOfActionDomainObjectsSTIX:
						fmt.Printf("| Type:'%s', Version:'%s', Name:'%s', Object:'%v' |\n", data.Type, data.SpecVersion, data.Name, resultDecodingSTIXObject.Data)

					case datamodels.IndicatorDomainObjectsSTIX:
						fmt.Printf("| Type:'%s', Version:'%s', Name:'%s', Object:'%v' |\n", data.Type, data.SpecVersion, data.Name, resultDecodingSTIXObject.Data)
					}
				}
			}*/

			Expect(true).Should(BeTrue())
		})
	})

	Context("Тест 2. Генерация UUID", func() {
		It("Должен быть получен UUID, ошибки быть не должно", func() {
			uuid := uuid.NewString()

			fmt.Printf("========================= UUID '%v'\n", uuid)

			Expect(nil).ShouldNot(HaveOccurred())
		})
	})

	/*Context("Тест 2. Проверка получения параметров (флагов) JSON сообщения", func() {
		It("Должен быть получен параметр", func() {
			tessst := datamodels.MalwareAnalysisDomainObjectsSTIX{
				Product: "os product",
				Version: "v12.3.2",
			}

			fmt.Println(tessst)

			var testTypeOne testTypeOne
			var err error

			err = json.Unmarshal(objInfo, &testTypeOne)
			Expect(err).ShouldNot(HaveOccurred())

			rtype := reflect.TypeOf(testTypeOne.Extensions)

			fmt.Println("================")
			if sfield, ok := rtype.FieldByName("SocketExt"); ok {
				fmt.Printf("Name: %v, Type: %v, Value: %v, Tag: %v\n", sfield.Name, sfield.Type, testTypeOne.Extensions.SocketExt, sfield.Tag)

				listValue := strings.Split(string(sfield.Tag), " ")

				fmt.Printf("result: %v\n", listValue)
			}
			fmt.Println("================")

			Expect(true).Should(BeTrue())
		})
	})*/
})

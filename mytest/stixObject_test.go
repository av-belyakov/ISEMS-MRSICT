package mytest_test

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StixObject", func() {
	reqJSON := []byte(`{
		"task_id": "ni883f838ggr8h9ehr3rgr38rg8f8g8wgd8g38",
		"section": "handling stix object",
		"task_was_generated_automatically": true,
		"user_name_generated_task": "Иванов Петр Васильевич",
		"request_details":	[
			{
				"type": "attack-pattern",
				"spec_version": "2.1",
				"id": "attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"name": "Spear Phishing as Practiced by Adversary X",
				"description": "A particular form of spear phishing where the attacker claims that the target had won a contest, including personal details, to get them to click on a link.",
				"external_references": [{ "source_name": "capec", "external_id": "CAPEC-163" }]
			},
			{
				"type": "course-of-action",
				"spec_version": "2.1",
				"id": "course-of-action--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"name": "Add TCP port 80 Filter Rule to the existing Block UDP 1434 Filter",
				"description": "This is how to add a filter rule to block inbound access to TCP port 80 to the existing UDP 1434 filter ..."
			},
			{
				"type": "indicator",
				"spec_version": "2.1",
				"id": "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"indicator_types": ["malicious-activity"],
				"name": "Poison Ivy Malware",
				"description": "This file is part of Poison Ivy",
				"pattern": "[ file:hashes.'SHA-256' = '4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877' ]",
				"pattern_type": "stix",
				"valid_from": "2016-01-01T00:00:00Z"
			}
		]}`)

	Context("Тест №1. Тест парсинга STIX объекта", func() {
		It("При чтении объекта не должно быть ошибок", func() {
			var (
				modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
				commonPropertiesObjectSTIX     datamodels.CommonPropertiesObjectSTIX
			)
			var err error

			err = json.Unmarshal(reqJSON, &modAPIRequestProcessingReqJSON)
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
				if err := json.Unmarshal(modAPIRequestProcessingReqJSON.RequestDetails, &listSTIXObjectJSON); err != nil {
					Expect(err).ShouldNot(HaveOccurred())
				}

				for _, item := range listSTIXObjectJSON {
					err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
					if err != nil {
						Expect(err).ShouldNot(HaveOccurred())
					}

					resultDecodingSTIXObject, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
					if err != nil {
						Expect(err).ShouldNot(HaveOccurred())
					}

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

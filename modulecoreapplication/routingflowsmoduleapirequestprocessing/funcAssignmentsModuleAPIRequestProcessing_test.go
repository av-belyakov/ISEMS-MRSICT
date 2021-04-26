package routingflowsmoduleapirequestprocessing_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FuncAssignmentsModuleAPIRequestProcessing", func() {
	var (
		errReadFile, errUnmarchalReq, errUnmarchalToSTIX error
		errCheckSTIXObjects, errDecSearch, errChecker    error
		docJSON                                          []byte
		l, sanitizeListElement                           []*datamodels.ElementSTIXObject
		modAPIRequestProcessingReqJSON                   datamodels.ModAPIRequestProcessingReqJSON
		searchReq, newSearchReq                          datamodels.ModAPIRequestProcessingResJSONSearchReqType
	)

	testSearchReq := json.RawMessage([]byte(`{
		"collection_name": "stix object",
		"search_parameters": {
			"documents_id": ["attack-pattern--1853f6a4-458f-5b4e-9b0f-ded361ae1002", "tool--0853f6a4-638f-5b4e-9b0f-ded001ae3822"],
			"documents_type": ["attack-pattern"],
			"created": {
				"start": "2015-12-21T19:59:11.000Z",
				"end": "2015-12-21T21:59:45.000Z"
			},
			"specific_search_fields": [
				{
					"object_name": "attack-pattern",
					"search_fields": {
						"name": "attack pattern$$ example to yahoo.com",
						"aliases": ["ap aliase 1", "ap aliase 2"]
					}
				},	
				{
					"object_name": "campaign",
					"search_fields": {
						"name": "comp name example\n",
						"first_seen": {
							"start": "2016-05-12T08:17:27.000Z",
							"end": "2016-05-12T12:31:17.000Z"
						},
						"last_seen": {
							"start": "2016-10-12T10:17:47.000Z",
							"end": "2016-05-12T10:29:02.000Z"
						}
					}
				},
				{
					"object_name": "ipv4-addr",
					"search_fields": {
						"value": ["124.12.5.33/31", "67.45.2.1/32", "89.0.213.4"]
					}
				},
				{
					"object_name": "ipv6-addr",
					"search_fields": {
						"value": ["2001:0db8::/96", "2001:0ab8:85a3:0000:0000:9a1e:0370:7334"]
					}
				},
				{
					"object_name": "report",
					"search_fields": {
						"name": "example report name"
					}
				}
			]
		}
	}`))

	var _ = BeforeSuite(func() {
		docJSON, errReadFile = ioutil.ReadFile("../../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)
		errCheckSTIXObjects = routingflowsmoduleapirequestprocessing.CheckSTIXObjects(l)
		sanitizeListElement = routingflowsmoduleapirequestprocessing.SanitizeSTIXObject(l)
		searchReq, errDecSearch = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectReqSearchParameters(&testSearchReq)
		newSearchReq, errChecker = routingflowsmoduleapirequestprocessing.CheckSearchSTIXObject(&searchReq)
	})

	Context("Тест 1. Проверка на наличие ошибок при предварительном преобразовании из JSON", func() {
		It("Ошибок при предварительном преобразовании из JSON быть не должно", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверяем функцию 'UnmarshalJSONObjectSTIXReq'", func() {
		It("должен быть получен список из 65 STIX объектов, ошибок быть не должно", func() {
			Expect(errUnmarchalToSTIX).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(65))
		})
	})

	Context("Тест 3. Выполнение валидации STIX объектов", func() {
		It("При выполнении валидации не должно быть ошибок", func() {
			Expect(errCheckSTIXObjects).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Выполнение санитаризации STIX объектов", func() {
		It("Количество STIX объектов после выполнения санитаризации должно соответствовать количеству объектов исходного среза", func() {
			Expect(len(sanitizeListElement)).To(Equal(len(l)))
		})
	})

	Context("Тест 5. Декодирование JSON документа, содержащего запросы к поисковой машине", func() {
		It("При декодирования запроса ошибок быть не должно", func() {

			fmt.Printf("Search Result:'%v'\n", searchReq)

			Expect(errDecSearch).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 6. Выполнение валидации и саниторизации запросов к поисковой машине", func() {
		It("Должна быть успешно выполненна валидация и саниторизация запросов", func() {

			if sp, ok := newSearchReq.SearchParameters.(datamodels.SearchThroughCollectionSTIXObjectsType); ok {
				fmt.Printf("DocumentsID: '%s'\n", sp.DocumentsID)
				fmt.Printf("DocumentsType: '%s'\n", sp.DocumentsType)
				fmt.Printf("Created: '%v'\n", sp.Created)
				for _, v := range sp.SpecificSearchFields {
					fmt.Printf("NEW search request, SpecificSearchFields.ObjectName:'%s', Name: '%s'\n", v.ObjectName, v.SearchFields.Name)
				}
			}

			Expect(errChecker).ShouldNot(HaveOccurred())
		})
	})
})

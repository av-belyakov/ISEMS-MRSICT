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
		errReadFile, errUnmarchalReq, errUnmarchalToSTIX, errUnmarchalSearchReq         error
		errCheckSTIXObjects, errDecSearch, errChecker, errReadFileJSONSearchSTIXExample error
		docJSON, docJSONSearchSTIX                                                      []byte
		l, sanitizeListElement                                                          []*datamodels.ElementSTIXObject
		modAPIRequestProcessingReqJSON                                                  datamodels.ModAPIRequestProcessingReqJSON
		modAPIRequestProcessingSearchReqJSON                                            datamodels.ModAPIRequestProcessingReqJSON
		searchReq, newSearchReq                                                         datamodels.ModAPIRequestProcessingResJSONSearchReqType
	)

	var _ = BeforeSuite(func() {
		docJSON, errReadFile = ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)
		errCheckSTIXObjects = routingflowsmoduleapirequestprocessing.CheckSTIXObjects(l)
		sanitizeListElement = routingflowsmoduleapirequestprocessing.SanitizeSTIXObject(l)

		docJSONSearchSTIX, errReadFileJSONSearchSTIXExample = ioutil.ReadFile("../../mytest/test_resources/jsonSearchSTIXExample.json")
		errUnmarchalSearchReq = json.Unmarshal(docJSONSearchSTIX, &modAPIRequestProcessingSearchReqJSON)
		searchReq, errDecSearch = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectReqSearchParameters(modAPIRequestProcessingSearchReqJSON.RequestDetails)
		newSearchReq, errChecker = routingflowsmoduleapirequestprocessing.CheckSearchSTIXObject(&searchReq)
	})

	Context("Тест 1. Проверка на наличие ошибок при предварительном преобразовании из JSON", func() {
		It("Ошибок при предварительном преобразовании из JSON быть не должно", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())

			Expect(errReadFileJSONSearchSTIXExample).ShouldNot(HaveOccurred())
			Expect(errUnmarchalSearchReq).ShouldNot(HaveOccurred())
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
					fmt.Printf("NEW search request, Name: '%s'\n", v.Name)
				}
			}

			Expect(errChecker).ShouldNot(HaveOccurred())
		})
	})
})

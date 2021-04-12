package routingflowsmoduleapirequestprocessing_test

import (
	"encoding/json"
	"io/ioutil"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FuncAssignmentsModuleAPIRequestProcessing", func() {
	var (
		errReadFile, errUnmarchalReq, errUnmarchalToSTIX, errCheckSTIXObjects error
		docJSON                                                               []byte
		l, sanitizeListElement                                                []*datamodels.ElementSTIXObject
		modAPIRequestProcessingReqJSON                                        datamodels.ModAPIRequestProcessingReqJSON
	)

	var _ = BeforeSuite(func() {
		docJSON, errReadFile = ioutil.ReadFile("../../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)
		errCheckSTIXObjects = routingflowsmoduleapirequestprocessing.CheckSTIXObjects(l)
		sanitizeListElement = routingflowsmoduleapirequestprocessing.SanitizeSTIXObject(l)
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
})

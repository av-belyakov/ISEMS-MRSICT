package commonlibs_test

import (
	"encoding/json"
	"io/ioutil"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type resultProcessingListSTIXObject struct {
	DataType string
	Data     interface{}
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
	)

	var _ = BeforeSuite(func() {
		docJSON, errReadFile = ioutil.ReadFile("../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		errUnmarchalList = json.Unmarshal(*modAPIRequestProcessingReqJSON.RequestDetails, &listSTIXObjectJSON)

		for _, item := range listSTIXObjectJSON {
			err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
			if err != nil {
				continue
			}

			resultDecodingSTIXObject, typeSTIXObject, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
			if err != nil {
				continue
			}

			listResult = append(listResult, &resultProcessingListSTIXObject{
				DataType: typeSTIXObject,
				Data:     resultDecodingSTIXObject,
			})
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

		It("", func() {
		})
	})
})

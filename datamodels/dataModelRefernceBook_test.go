package datamodels_test

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"os"
	//"path/filepath"

	"ISEMS-MRSICT/datamodels"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type JsonRawTestData []json.RawMessage

//HiLevelNilAndErrorParametersAssertion -
// Утверждения проверки nil-значений и ошибок в json.Unmarshal
func HiLevelNilAndErrorParametersAssertion(vv json.RawMessage) {

	//	for _, v := range vv {
	var trbr datamodels.ReferencesBookReq
	err := json.Unmarshal(vv, &trbr)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(trbr.RequestDetails).ShouldNot(BeNil())
	Expect(len(trbr.RequestDetails)).ShouldNot(Equal(0))
	for _, v := range trbr.RequestDetails {
		Expect(v.OP).ShouldNot(BeNil())
		Expect(v.Name).ShouldNot(BeNil())
	}
	//	}

}

var _ = Describe("RefernceBook", func() {

	var (
		fd       []byte
		testData json.RawMessage
		//testData  []interface{}
		err, err1 error
	)

	var _ = BeforeSuite(func() {
		dir, _ := os.Getwd()
		filePath := filepath.Join(dir, "..", "mytest/test_resources/ReferersBookAPIHierarchicalNotationResponseExample.json")
		fd, err = ioutil.ReadFile(filePath)
		err1 = json.Unmarshal(fd, &testData)
	})

	Context("Подготовка тестовых данных", func() {
		It("Открытие файл с тестовыми данными", func() {
			Expect(err).NotTo(HaveOccurred())
			l := len(fd)
			Expect(fd).NotTo(HaveLen(0), fmt.Sprintf("Файл с тестовыми данными содержитне %b байт", l))
		})
		It("Первичне преобразование тестового набора", func() {
			Expect(err1).NotTo(HaveOccurred())
		})
		It("Не пустой набор тестовых данных", func() {
			l := len(testData)
			Expect(l).NotTo(Equal(0), fmt.Sprintf("Число тестовых объектов = %b", l))
			//for i, v := range testData {
			//	l := len(v)
			Expect(l).NotTo(Equal(0), fmt.Sprintf("Тестовый объект имеет размер %b байт", l))
			//}
		})
	})

	Describe("Тестирование Custom Unmarshal-инга запросов JSON API для объектов ReferenceBook", func() {

		Context("Тест№1 Корректность заполнения поля Parameters при Unmarshal-инге тестового набора", func() {

			It("Проверка отсутствия nil значений в поле Parameters и возниконовения ошибок при работе json.Unmarshal", func() {
				HiLevelNilAndErrorParametersAssertion(testData)
			})

		})

	})

	/*
		Context("Tect №1 Unmarshal ReferenceBookAPIReqJSON в интерфейсный тип RefernceBooker", func() {
			var rBooker datamodels.RefernceBooker
			err := json.Unmarshal(v, &rBooker)
			It("При Unmarshal не должно быть ошибок", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(rBooker.GetCommand()).Should(Equal("get"))
			})

		})
	}*/
})

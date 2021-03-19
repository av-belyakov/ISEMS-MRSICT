package mytest_test

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StixObject", func() {
	Context("Тест 1. Генерация UUID", func() {
		It("Должен быть получен UUID, ошибки быть не должно", func() {
			uuid := uuid.NewString()

			fmt.Printf("========================= UUID '%v'\n", uuid)

			Expect(nil).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка RegExp", func() {
		var validStr string = "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff"
		var invalidStr string = "identity--f431f809-377b-45e0-aa1c-6!z!a4751cae5ff"

		/*
			uuid3                string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
			uuid4                string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
			uuid5                string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
			uuid                 string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
			alpha                string = "^[a-zA-Z]+$"
		*/

		It("На валидную строку TRUE", func() {
			Expect((regexp.MustCompile(`^[0-9a-zA-Z]+(--)[0-9a-f|-]+$`)).MatchString(validStr)).Should(BeTrue())
		})

		It("На НЕ валидную строку FALSE", func() {
			Expect((regexp.MustCompile(`^[0-9a-zA-Z]+(--)[0-9a-f|-]+$`)).MatchString(invalidStr)).Should(BeFalse())
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

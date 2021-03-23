package mytest_test

import (
	"fmt"
	"regexp"
	"strings"

	govalidator "github.com/asaskevich/govalidator"
	"github.com/kennygrant/sanitize"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gosanitize "github.com/whosonfirst/go-sanitize"
)

var _ = Describe("RegexpMatchstring", func() {
	Context("Тест 1. Тестируем регулярные выражения для валидации поля 'SpecVersion' STIX объектов", func() {
		var pattern string = `^[0-9a-z.]+$`
		var validStr string = "2.1"
		var invalidStr string = "2,1"

		It("Должно быть TRUE на валидное содержимое поля", func() {
			Expect((regexp.MustCompile(pattern)).MatchString(validStr)).Should(BeTrue())
		})
		It("Должно быть TRUE на валидное содержимое поля", func() {
			Expect((regexp.MustCompile(pattern)).MatchString(invalidStr)).Should(BeFalse())
		})
	})

	Context("Тест 2. Тестируем регулярные выражения для валидации поля 'CreatedByRef' STIX объектов", func() {
		var pattern string = `^[0-9a-zA-Z]+(--)[0-9a-f|-]+$`
		var validStr string = "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff"
		var invalidStr string = "identity--f431f809-377b-45e0-aa1c-6!z!a4751cae5ff"

		It("Должно быть TRUE на валидное содержимое поля", func() {
			Expect((regexp.MustCompile(pattern)).MatchString(validStr)).Should(BeTrue())
		})

		It("Должно быть TRUE на валидное содержимое поля", func() {
			Expect((regexp.MustCompile(pattern)).MatchString(invalidStr)).Should(BeFalse())
		})
	})

	Context("Тест 3. Тестируем регулярные выражения для валидации поля 'Lang' STIX объектов", func() {
		var pattern string = `^[a-zA-Z-]+$`
		var validStr string = "ru"
		var invalidStr string = "3 ру"

		It("Должно быть TRUE на валидное содержимое поля", func() {
			Expect((regexp.MustCompile(pattern)).MatchString(validStr)).Should(BeTrue())
		})

		It("Должно быть TRUE на валидное содержимое поля", func() {
			Expect((regexp.MustCompile(pattern)).MatchString(invalidStr)).Should(BeFalse())
		})
	})

	Context("Тест 4. С помощбю sanitize Проверяем функцию выполняющую 'очистку' строк от нежелательных символов или вырожений", func() {
		str := "Mozilla/5.0 (Windows; U; Windows NT 5.1; <' \nen-US; rv:1.6)Gecko/20040113"
		resultStr := sanitize.Accents(str)

		fmt.Printf("String sanitize result: '%v'\n", resultStr)

		It("Исходная строка должна содержать указанное невалидное значение", func() {
			Expect(strings.Contains(str, "<'")).Should(BeTrue())
		})

		It("Результирующая строка не должна содержать невалидное значение", func() {
			Expect(strings.Contains(resultStr, "<'")).Should(BeFalse())
		})
	})

	Context("Тест 5. С помощью go-sanitize Проверяем функцию выполняющую 'очистку' строк от нежелательных символов или вырожений", func() {
		str := "Mozilla/5.0 (Windows; U; Windows NT 5.1; \n en-US; rv:1.6)Gecko/20040113"
		opts := gosanitize.DefaultOptions()
		resultStr, _ := gosanitize.SanitizeString(str, opts)

		fmt.Printf("String sanitize result: '%v'\n", resultStr)

		It("Исходная строка должна содержать указанное невалидное значение", func() {
			Expect(strings.Contains(str, "\n")).Should(BeTrue())
		})

		It("Результирующая строка не должна содержать невалидное значение", func() {
			Expect(strings.Contains(resultStr, "\n")).Should(BeFalse())
		})
	})

	Context("Тест 6. Тестируем пакет 'govalidator'", func() {
		str := "Mozilla/5.0 (Windows; U; Windows NT 5.1; \n en-US; &where where rv:1.6)Gecko/20040113"
		resultStr := govalidator.BlackList(str, "\n")

		fmt.Printf("++++++ govalidator: '%s'\n", resultStr)

		It("Должен быть удален невалидный символ '\n'", func() {
			Expect(strings.Contains(resultStr, "\n")).Should(BeFalse())
		})
	})

	/*Context("Тест .", func(){
		It("", func(){

		})
	})*/
})

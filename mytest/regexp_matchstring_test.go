package mytest_test

import (
	"ISEMS-MRSICT/commonlibs"
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

	Context("Тест 4. С помощью sanitize Проверяем функцию выполняющую 'очистку' строк от нежелательных символов или вырожений", func() {
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
		str := "Mozilla/5.0 (Windows; U; Windows NT 5.1; \n \ten-'US; rv:1.6)Gecko/20040113"
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

	Context("Тест 6. Тестируем функцию которая преобразовывает некоторые символы в их HTML код", func() {
		It("В строке должны быть заменены все, некоторые специальные символы в их HTML код", func() {
			str := `Mozilla/5.0 (Windows; U; Windows NT 5.1; \n en-US<; $where " \twhere' rv>:1.6)Gecko/20040113`
			strRes := commonlibs.StringSanitize(str)

			fmt.Printf("+ stringSanitize: '%s'\n", str)
			fmt.Printf("+ stringSanitize: '%s'\n", strRes)

			charOne := strings.Contains(strRes, "\n")
			charTwo := strings.Contains(strRes, "<")
			charThree := strings.Contains(strRes, "\t")
			charFour := strings.Contains(strRes, ">")
			charFive := strings.Contains(strRes, "$")

			fmt.Printf("charOne: '%v', charTwo: '%v', charThree: '%v', charFour: '%v', charFive: '%v'\n", charOne, charTwo, charThree, charFour, charFive)

			charIsExist := charOne || charTwo || charThree || charFour || charFive

			Expect(charIsExist).Should(BeFalse())
		})
	})

	Context("Тест 7. Тестируем пакет 'govalidator'", func() {
		It("Должен быть удален невалидный символ '\n'", func() {

			specialCharacters := [][2]string{
				{"$", " &#36; "},
				{"\"", " &quot; "},
				{"'", " &apos; "},
				{"<", " &lt; "},
				{">", " &gt; "},
				{"\\n", " &#010; "},
				{"\\t", " &#009; "},
				{"\\r", " &#013; "},
			}
			var resultStrTwo string

			str := `Mozilla/5.0 (Windows; U; Windows NT 5.1; \n en-US<; $where " \twhere' rv>:1.6)Gecko/20040113`
			resultStr := govalidator.ReplacePattern(str, "(\\|\"|\\')", "\\/")

			fmt.Printf("++++++ govalidator: '%s'\n", str)
			fmt.Printf("++++++ govalidator: '%s'\n", resultStr)

			resultStrTwo = str
			for ch := range specialCharacters {

				//fmt.Printf("--- 1. '%s', 2. '%s' ---\n", specialCharacters[ch][0], specialCharacters[ch][1])

				resultStrTwo = govalidator.ReplacePattern(resultStrTwo, specialCharacters[ch][0], specialCharacters[ch][1])
			}

			fmt.Printf("++++++ MyFunc: '%s'\n", resultStrTwo)

			Expect(strings.Contains(resultStr, "\n")).Should(BeFalse())
		})
	})

	/*Context("Тест .", func(){
		It("", func(){

		})
	})*/
})

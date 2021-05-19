package mytest_test

import (
	"fmt"
	"regexp"
	"strings"

	"ISEMS-MRSICT/commonlibs"

	govalidator "github.com/asaskevich/govalidator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

	Context("Тест 4. Тестируем регулярне вырожения для проверки поля ID", func() {
		It("Должно быть TRUE на валидное содержимое поля ID", func() {
			id := "attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5"
			Expect(regexp.MustCompile(`^(attack-pattern--)[0-9a-f|-]+$`).MatchString(id)).Should(BeTrue())
		})
		It("Должно быть FALSE на валидное содержимое поля ID", func() {
			id := "attack-pattxern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5"
			Expect(regexp.MustCompile(`^(attack-pattern--)[0-9a-f|-]+$`).MatchString(id)).Should(BeFalse())
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

	Context("Тест 8. Делим список елементов на части, где размер части не превышает заданное значение", func() {
		const chunkSize int = 10

		It("Список состоящий из 73 элементов должен быть разделен на 8 частей (по 10 элементов в части)", func() {
			leo := make([]int, 0, 73)
			for i := 1; i <= 73; i++ {
				leo = append(leo, i)
			}

			num := commonlibs.GetCountChunk(int64(len(leo)), chunkSize)

			nleo := make([][]int, 0, num)
			min := 0
			max := chunkSize
			for i := 0; i < num; i++ {
				if i == 0 {
					nleo = append(nleo, leo[:max])
				} else if i == num-1 {
					nleo = append(nleo, leo[min:])
				} else {
					nleo = append(nleo, leo[min:max])
				}

				min = min + chunkSize
				max = max + chunkSize
			}

			fmt.Printf("==== Before list element: '%v'\n, ==== In total chunks: '%d'\n, ==== list elemet: '%v'\n", leo, num, nleo)

			Expect(len(nleo)).Should(Equal(8))
		})

		It("Список состоящий из 17 элементов должен быть разделен на 2 частей (по 10 элементов в части)", func() {
			leo := make([]int, 0, 17)
			for i := 1; i <= 17; i++ {
				leo = append(leo, i)
			}

			num := commonlibs.GetCountChunk(int64(len(leo)), chunkSize)

			nleo := make([][]int, 0, num)
			min := 0
			max := chunkSize
			for i := 0; i < num; i++ {
				if i == 0 {
					nleo = append(nleo, leo[:max])
				} else if i == num-1 {
					nleo = append(nleo, leo[min:])
				} else {
					nleo = append(nleo, leo[min:max])
				}

				min = min + chunkSize
				max = max + chunkSize
			}

			fmt.Printf("==== 1 Before list element: '%v'\n, ==== 1 In total chunks: '%d'\n, ==== 1 list elemet: '%v'\n", leo, num, nleo)

			Expect(len(nleo)).Should(Equal(2))
		})
	})

	/*Context("Тест .", func(){
		It("", func(){

		})
	})*/
})

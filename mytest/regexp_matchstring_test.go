package mytest_test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	govalidator "github.com/asaskevich/govalidator"
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

	Context("Тест 8. Выполняем проверку типа 'attack-pattern' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа AttackPatternDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			apbyte := json.RawMessage([]byte(`{
	"type": "attack-pattern",
	"spec_version": "2.1",
	"id": "attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5",
	"created": "2016-05-12T08:17:27.000Z",
	"modified": "2016-05-12T08:17:27.000Z",
	"name": "Spear Phishing as Practiced $< by Adversary\n' X",
	"description": "A particular form of spear phishing where the attacker claims that $ the target <> had won a contest, including personal details, to get them to click on a link.",
	"external_references": [{ "source_name": "capec", "external_id": "CAPEC-163" }],
	"labels": ["__lable 1", "__lable 2", "__lable$ 3 <.'"],
	"aliases": ["new ali&ase", "aliase <>", "Ali12$$"]
}`))
			var apo datamodels.AttackPatternDomainObjectsSTIX
			apoTmp, err := apo.DecoderJSON(&apbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newapo, ok := apoTmp.(datamodels.AttackPatternDomainObjectsSTIX)
			apoIsTrue := newapo.CheckingTypeFields()
			newapo = newapo.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newapo.ToStringBeautiful())

			Expect(apoIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 9. Выполняем проверку типа 'campaign' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа CampaignDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			cobyte := json.RawMessage([]byte(`{
				"type": "campaign",
				"spec_version": "2.1",
				"id": "campaign--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"modified": "2016-04-06T20:03:00.000Z",
				"name": "Green Group Attacks 'Against' Finance",
				"description": "Campaign by Green Group against a series of targets in the financialservices $sector.",
				"aliases": ["aa12$", "asd_3'", "\nerrx"],
				"first_seen": "2016-04-12T20:12:05.000Z",
				"last_seen": "2016-04-06T20:22:10.000Z",
				"objective": "Example objectiv<e!"
}`))
			var co datamodels.CampaignDomainObjectsSTIX
			coTmp, err := co.DecoderJSON(&cobyte)

			Expect(err).ShouldNot(HaveOccurred())

			newco, ok := coTmp.(datamodels.CampaignDomainObjectsSTIX)
			coIsTrue := newco.CheckingTypeFields()
			newco = newco.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newco.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 10. Выполняем проверку типа 'course-of-action' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа CourseOfActionDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			cabyte := json.RawMessage([]byte(`{
				"type": "course-of-action",
				"spec_version": "2.1",
				"id": "course-of-action--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"name": "Add TCP port 80 Filter Rule to the existing Block UDP 1434 Filter",
				"description": "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ..."
}`))
			var ca datamodels.CourseOfActionDomainObjectsSTIX
			caTmp, err := ca.DecoderJSON(&cabyte)

			Expect(err).ShouldNot(HaveOccurred())

			newca, ok := caTmp.(datamodels.CourseOfActionDomainObjectsSTIX)
			coIsTrue := newca.CheckingTypeFields()
			newca = newca.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newca.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 11. Выполняем проверку типа 'grouping' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа GroupingDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			gdbyte := json.RawMessage([]byte(`{
				"type": "grouping",
				"spec_version": "2.1",
				"id": "grouping--84e4d88f-44ea-4bcd-bbf3-b2c1c320bcb3",
				"created_by_ref": "identity--a463ffb3-1bd9-4d94-b02d-74e4f1658283",
				"created": "2015-12-21T19:59:11.000Z",
				"modified": "2015-12-21T19:59:11.000Z",
				"name": "The Black Vine Cyberespionage Group",
				"description": "A simple collection of Black Vine Cyberespionage Group attributed intel",
				"context": "suspicious-activity",
				"object_refs": [
					"indicator--26ffb872-1dd9-446e-b6f5-d58527e5b5d2",
					"campaign--83422c77-904c-4dc1-aff5-5c38f3a2c55c",
					"relationship--f82356ae-fe6c-437c-9c24-6b64314ae68a",
					"file--0203b5c8-f8b6-4ddb-9ad0-527d727f968b"
					]
}`))
			var gd datamodels.GroupingDomainObjectsSTIX
			gdTmp, err := gd.DecoderJSON(&gdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newgd, ok := gdTmp.(datamodels.GroupingDomainObjectsSTIX)
			coIsTrue := newgd.CheckingTypeFields()
			newgd = newgd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newgd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 12. Выполняем проверку типа 'identity' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа IdentityDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			idbyte := json.RawMessage([]byte(`{
				"type": "identity",
				"spec_version": "2.1",
				"id": "identity--023d105b-752e-4e3c-941c-7d3f3cb15e9e",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:00.000Z",
				"modified": "2016-04-06T20:03:00.000Z",
				"name": "John Smith",
				"identity_class": "individual"
}`))
			var id datamodels.IdentityDomainObjectsSTIX
			idTmp, err := id.DecoderJSON(&idbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newid, ok := idTmp.(datamodels.IdentityDomainObjectsSTIX)
			coIsTrue := newid.CheckingTypeFields()
			newid = newid.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newid.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 13. Выполняем проверку типа 'indicator' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа IndicatorDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			idbyte := json.RawMessage([]byte(`{
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
}`))
			var id datamodels.IndicatorDomainObjectsSTIX
			idTmp, err := id.DecoderJSON(&idbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newid, ok := idTmp.(datamodels.IndicatorDomainObjectsSTIX)
			coIsTrue := newid.CheckingTypeFields()
			newid = newid.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newid.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 14. Выполняем проверку типа 'infrastructure' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа InfrastructureDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			idbyte := json.RawMessage([]byte(`{
				"type":"infrastructure",
				"spec_version": "2.1",
				"id":"infrastructure--38c47d93-d984-4fd9-b87b-d69d0841628d",
				"created":"2016-05-07T11:22:30.000Z",
				"modified":"2016-05-07T11:22:30.000Z",
				"name":"Poison Ivy C2",
				"infrastructure_types": ["command-and-control"]
}`))
			var id datamodels.InfrastructureDomainObjectsSTIX
			idTmp, err := id.DecoderJSON(&idbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newid, ok := idTmp.(datamodels.InfrastructureDomainObjectsSTIX)
			coIsTrue := newid.CheckingTypeFields()
			newid = newid.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newid.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 15. Выполняем проверку типа 'infrastructure' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа IntrusionSetDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			isdbyte := json.RawMessage([]byte(`{
				"type": "intrusion-set",
				"spec_version": "2.1",
				"id": "intrusion-set--4e78f46f-a023-4e5f-bc24-71b3ca22ec29",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"name": "Bobcat Breakin",
				"description": "Incidents usually feature a shared TTP of a bobcat being released within the building containing network access, scaring users to leave their computers without locking them first. Still determining where the threat actors are getting the bobcats.",
				"aliases": ["Zookeeper"],
				"goals": ["acquisition-theft", "harassment", "damage"]
}`))
			var isd datamodels.IntrusionSetDomainObjectsSTIX
			isdTmp, err := isd.DecoderJSON(&isdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newisd, ok := isdTmp.(datamodels.IntrusionSetDomainObjectsSTIX)
			coIsTrue := newisd.CheckingTypeFields()
			newisd = newisd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newisd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 16. Выполняем проверку типа 'location' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа LocationDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			ldbyte := json.RawMessage([]byte(`{
				"type": "location",
				"spec_version": "2.1",
				"id": "location--a6e9345f-5a15-4c29-8bb3-7dcc5d168d64",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:00.000Z",
				"modified": "2016-04-06T20:03:00.000Z",
				"latitude": 48.8566,
				"longitude": 2.3522,
				"region": "south-eastern-asia",
				"country": "th",
				"administrative_area": "Tak",
				"postal_code": "63170"
}`))
			var ld datamodels.LocationDomainObjectsSTIX
			ldTmp, err := ld.DecoderJSON(&ldbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newld, ok := ldTmp.(datamodels.LocationDomainObjectsSTIX)
			coIsTrue := newld.CheckingTypeFields()
			newld = newld.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newld.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 16. Выполняем проверку типа 'malware' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа MalwareDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "malware",
				"spec_version": "2.1",
				"id": "malware--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"name": "Cryptolocker",
				"description": "A variant of the cryptolocker family",
				"malware_types": ["ransomware"],
				"is_family": false    
}`))
			var md datamodels.MalwareDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.MalwareDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			fmt.Println(newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	/*Context("Тест .", func(){
		It("", func(){

		})
	})*/
})

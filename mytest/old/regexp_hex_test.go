package mytest_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	//"ISEMS-MRSICT/mytest"

	"regexp"
)

// HashesTypeSTIX тип "hashes", по терминологии STIX, содержащий хеш значения, где <тип_хеша>:<хеш>
type HashesTypeSTIX map[string]string

// CheckHashesTypeSTIX выполняет проверку значений типа HashesTypeSTIX
func (htstix *HashesTypeSTIX) CheckHashesTypeSTIX() bool {
	if len(*htstix) == 0 {
		return true
	}

	pattern := regexp.MustCompile(`^[0-9a-zA-Z-_=]+$`)
	for k, v := range *htstix {
		if !pattern.MatchString(k) {
			return false
		}

		if !pattern.MatchString(v) {
			return false
		}
	}

	return true
}

var _ = Describe("RegexpHex", func() {
	Context("Тест 1. Выполняем проверку регулярного выражения", func() {
		It("При валидном содержимом должно быть TRUE", func() {
			htstix := HashesTypeSTIX{
				"MD5":      "dwffw33gg3",
				"SHA-1":    "gg4g4g4tg4hg4gh4",
				"SHA3-256": "g4344g4g3efg3f3f3",
				//"SHA3-256": "jtjtj6",
				//"SHA-1":    "ngnmty66j6j",
				//"MD5":      "6h66hrghtjhymymthnngn",
			}

			Expect(htstix.CheckHashesTypeSTIX()).Should(BeTrue())
		})
	})
})

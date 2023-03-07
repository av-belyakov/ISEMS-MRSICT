package mytest_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ISEMS-MRSICT/datamodels"
)

var _ = Describe("MetodsStixObject", func() {
	Context("Тест 30. Выполняем проверку типа 'network-traffic' методом ValidateStruct", func() {
		It("На валидное содержимое типа NetworkTrafficCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
	   				"type": "network-traffic",
	   				"spec_version": "2.1",
	   				"id": "network-traffic--630d7bb1-0bbc-53a6-a6d4-f3c2d35c2734",
	   				"src_ref": "ipv4-addr--e42c19c8-f9fe-5ae9-9fc8-22c398f78fb",
	   				"dst_ref": "ipv4-addr--03b708d9-7761-5523-ab75-5ea096294a68",
	   				"src_port": 2487,
	   				"dst_port": 1723,
	   				"protocols": ["ipv4", "tcp"],
	   				"src_byte_count": 147600,
	   				"src_packets": 100,
	   				"dst_byte_count": 935750,
	   				"encapsulates_refs": ["network-traffic--53e0bf48-2eee-5c03-8bde-ed7049d2c0a3"],
	   				"ipfix": {
	   					"minimumIpTotalLength": 32,
	   					"maximumIpTotalLength": 2556
	   				},
	   				"extensions": {
	   					"http-request-ext": {
	   						"request_method": "get",
	   						"request_value": "/download.html",
	   						"request_version": "http/1.1",
	   						"request_header": {
	   							"Accept-Encoding": "gzip,deflate",
	   							"User-Agent": "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.6)Gecko/20040113",
	   							"Host": "www.example.com"
	   						}
	   					},
	   					"socket-ext": {
	   						"is_listening": true,
	   						"address_family": "AF_INET",
	   						"socket_type": "SOCK_STREAM"
	   					}
	   				}
	   }`))
			var md datamodels.NetworkTrafficCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.NetworkTrafficCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())
		})
	})
})

package mytest_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"ISEMS-MRSICT/datamodels"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testTypeOne struct {
	Type        string         `json:"type"`
	SpecVersion string         `json:"spec_version"`
	ID          string         `json:"id"`
	DstRef      string         `json:"dst_ref"`
	Protocols   []string       `json:"protocols"`
	Extensions  testTypeCommon `json:"extensions"`
}

type testTypeTwo struct {
	Type        string                 `json:"type"`
	SpecVersion string                 `json:"spec_version"`
	ID          string                 `json:"id"`
	DstRef      string                 `json:"dst_ref"`
	Protocols   []string               `json:"protocols"`
	Extensions  map[string]interface{} `json:"extensions"`
}

type testTypeCommon struct {
	HttpRequestExt testTypeCommanHTTP `json:"http-request-ext"`
	SocketExt      testTypeSocketExt  `json:"socket-ext" bson:"socket-ext" required:"true" minValue:"3" maxValue:"10"`
	TcpExt         testTypeTCPExt     `json:"tcp-ext"`
	Rdsv           testRdsv           `json:"rdsv"`
}

type testRdsv struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

type testTypeCommanHTTP struct {
	RequestMethod  string                `json:"request_method"`
	RequestValue   string                `json:"request_value"`
	RequestVersion string                `json:"request_version"`
	RequestHeader  testTypeHTTPReqHeader `json:"request_header"`
}

type testTypeHTTPReqHeader struct {
	AcceptEncoding string `json:"Accept-Encoding"`
	UserAgent      string `json:"User-Agent"`
	Host           string `json:"Host"`
}

type testTypeSocketExt struct {
	AddressFamily string `json:"address_family" bson:"address_family"`
	IsBlocking    bool   `json:"is_blocking" bson:"is_blocking"`
	IsListening   bool   `json:"is_listening" bson:"is_listening"`
	SocketType    string `json:"socket_type" bson:"socket_type"`
}

type testTypeTCPExt struct {
	SrcFlagsHex string `json:"src_flags_hex"`
	DstFlagsHex string `json:"dst_flags_hex"`
}

//`json:""`

var _ = Describe("StixObject", func() {
	objInfo := []byte(`{
		"type": "network-traffic",
		"spec_version": "2.1",
		"id": "network-traffic--f8ae967a-3dc3-5cdf-8f94-8505abff00c2",
		"dst_ref": "ipv4-addr--6da8dad3-4de3-5f8e-ab23-45d0b8f12f16",
		"protocols": [
			"tcp", "http"
		],
		"extensions": {
			"http-request-ext": {
				"request_method": "get",
				"request_value": "/download.html",
				"request_version": "http/1.1",
				"request_header": {
					"Accept-Encoding": "gzip,deflate",
					"User-Agent": "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.6) Gecko/20040113",
					"Host": "www.example.com"
				}
			},
			"socket-ext": {
				"is_listening": true,
				"address_family": "AF_INET",
				"socket_type": "SOCK_STREAM"
			},
			"tcp-ext": {
				"src_flags_hex": "00000002"
			},
			"undef_type_test": "test string"
		}
	}`)

	Context("Тест №1. Тест парсинга STIX объекта", func() {
		/*
			"tcp-ext": {
				"src_flags_hex": "00000002"
			},
		*/

		It("При чтении объекта не должно быть ошибок", func() {
			var testTypeOne testTypeOne
			//var testTypeTwo testTypeTwo
			var err error

			err = json.Unmarshal(objInfo, &testTypeOne)
			Expect(err).ShouldNot(HaveOccurred())

			//err = json.Unmarshal(objInfo, &testTypeTwo)
			//Expect(err).ShouldNot(HaveOccurred())

			/*			fmt.Println(testTypeOne)
						fmt.Println("--------------- Extensions Before Processing ----------------")

						for n, v := range testTypeOne.Extensions {
							fmt.Printf("'%v'\n    '%v'\n", n, v)
						}*/

			fmt.Printf("+++ '%v' +++\n", testTypeOne.Extensions)

			fmt.Println("--------------- Extensions After Processing ----------------")

			/*			for n, v := range testTypeOne.Extensions {

						switch n {
						case "http-request-ext":
							fmt.Println("****** TypeName:http-request-ext ******")
							fmt.Println(v)

							newMsg := testTypeCommanHTTP{}
							err = json.Unmarshal(*v, &newMsg)
							if err != nil {
								fmt.Println(err)
							}

							fmt.Println(newMsg)

							//if result, ok := v.(testTypeCommanHTTP); ok {
							//	fmt.Printf("'%v'\n    '%v'\n", n, result)
							//	fmt.Printf("        Method:%v\n", result.RequestMethod)
							//	fmt.Printf("        Value:%v\n", result.RequestValue)
							//	fmt.Printf("        Version:%v\n", result.RequestVersion)
							//	fmt.Printf("        Host:%v\n", result.RequestHeader.Host)
							//	fmt.Printf("        AcceptEncoding:%v\n", result.RequestHeader.AcceptEncoding)
							//	fmt.Printf("        UserAgent:%v\n", result.RequestHeader.UserAgent)
							//}
						case "socket-ext":
							fmt.Println("****** TypeName:socket-ext ******")
							fmt.Println(v)

							//if result, ok := v.(testTypeSocketExt); ok {
							//	fmt.Printf("'%v'\n    '%v'\n", n, result)
							//}
						case "tcp-ext":
							fmt.Println("****** TypeName:tcp-ext ******")
							fmt.Println(v)

							//if result, ok := v.(testTypeTCPExt); ok {
							//	fmt.Printf("'%v'\n    '%v'\n", n, result)
							//}

						default:
							fmt.Printf("Undefined type: '%v'!!!\n", n)
						}

						//fmt.Printf("'%v'\n    '%v'\n", n, v)
					}*/

			Expect(true).Should(BeTrue())
		})
	})

	Context("Тест 2. Проверка получения параметров (флагов) JSON сообщения", func() {
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
	})
})

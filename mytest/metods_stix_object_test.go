package mytest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

var (
	connectError error
	qp           interactionmongodb.QueryParameters
	dm           datamodels.MongoDBSettings
	cdmd         interactionmongodb.ConnectionDescriptorMongoDB
)

var _ = BeforeSuite(func() {
	dm = datamodels.MongoDBSettings{
		Host:     "192.168.9.208",
		Port:     37017,
		User:     "module-isems-mrsict",
		Password: "vkL6Znj$Pmt1e1",
		NameDB:   "isems-mrsict",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	cdmd = interactionmongodb.ConnectionDescriptorMongoDB{
		Ctx:       ctx,
		CtxCancel: cancel,
	}

	connectError = cdmd.CreateConnection(&dm)

	qp = interactionmongodb.QueryParameters{
		NameDB:         "isems-mrsict",
		CollectionName: "stix_object_collection",
		ConnectDB:      cdmd.Connection,
	}
})

var _ = AfterSuite(func() {
	//ctxCancel()
	cdmd.CtxCancel()
})

var _ = Describe("MetodsStixObject", func() {
	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Получаем информацию о STIX объект с ID 'network-traffic--68ead79b-28db-4811-85f6-971f46548343'", func() {
		It("должен быть получен список из 1 STIX объекта, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.id", Value: "network-traffic--68ead79b-28db-4811-85f6-971f46548343"}})

			//fmt.Printf("___ ERROR find ID:'network-traffic--68ead79b-28db-4811-85f6-971f46548343' - '%v'\n", err)
			/*l := []*datamodels.NetworkTrafficCyberObservableObjectSTIX{}
			for cur.Next(context.Background()) {
				var model datamodels.NetworkTrafficCyberObservableObjectSTIX
				_ = cur.Decode(&model)

				l = append(l, &model)
			}*/

			listelm := interactionmongodb.GetListElementSTIXObject(cur)

			for _, v := range listelm {
				//fmt.Printf("Found STIX object with ID:'network-traffic--68ead79b-28db-4811-85f6-971f46548343' - '%v'\n\tIPFix = %v\n", *v, &v.IPFix)
				fmt.Printf("Found STIX object with ID:'network-traffic--68ead79b-28db-4811-85f6-971f46548343' - '%v'\n", *v)
			}

			Expect(err).ShouldNot(HaveOccurred())
			// Expect(len(l)).Should(Equal(1))
			Expect(len(listelm)).Should(Equal(1))
		})
	})

	Context("Тест 3. Выполняем проверку типа 'network-traffic' методом ValidateStruct", func() {
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

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.ValidateStruct()).Should(BeTrue())
		})
	})
})

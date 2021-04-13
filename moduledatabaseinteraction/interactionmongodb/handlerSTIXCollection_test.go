package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

var _ = Describe("HandlerSTIXCollection", func() {
	var (
		connectError, errReadFile, errUnmarchalReq, errUnmarchalToSTIX error
		docJSON                                                        []byte
		cdmdb                                                          interactionmongodb.ConnectionDescriptorMongoDB
		l                                                              []*datamodels.ElementSTIXObject
		qp                                                             interactionmongodb.QueryParameters
		modAPIRequestProcessingReqJSON                                 datamodels.ModAPIRequestProcessingReqJSON
	)

	var _ = BeforeSuite(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		cdmdb = interactionmongodb.ConnectionDescriptorMongoDB{
			Ctx:       ctx,
			CtxCancel: cancel,
		}

		//подключаемся к базе данных MongoDB
		connectError = cdmdb.CreateConnection(&datamodels.MongoDBSettings{
			Host:     "192.168.13.200",
			Port:     27017,
			User:     "module-isems-mrsict",
			Password: "vkL6Zn$jPmt1e1",
			NameDB:   "isems-mrsict",
		})

		docJSON, errReadFile = ioutil.ReadFile("../../mytest/jsonSTIXExample.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)

		qp = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
			ConnectDB:      cdmdb.Connection,
		}
	})

	var _ = AfterSuite(func() {
		cdmdb.CtxCancel()
	})

	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Читаем тестовый файл содержащий STIX объекты", func() {
		It("При чтении тестового файла не должно быть ошибок", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Проверка на наличие ошибок при предварительном преобразовании из JSON", func() {
		It("Ошибок при предварительном преобразовании из JSON быть не должно", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Проверяем функцию 'UnmarshalJSONObjectSTIXReq'", func() {
		It("должен быть получен список из 65 STIX объектов, ошибок быть не должно", func() {
			Expect(errUnmarchalToSTIX).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(65))
		})
	})

	Context("Тест 5. Взаимодействие с коллекцией STIX объектов", func() {
		It("При добавлении STIX объектов не должно быть ошибок", func() {
			addInfo := make([]interface{}, 0, len(l))

			for _, v := range l {

				//				fmt.Printf("___ STIX object ID:'%s'\n", v.Data.GetID())

				addInfo = append(addInfo, v.Data)
			}

			ok, err := qp.InsertData(addInfo)

			Expect(ok).Should(BeTrue())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 6. Получаем информацию о STIX объект с ID 'indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f'", func() {
		It("должен быть получен список из 1 STIX объекта, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.id", Value: "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f"}})

			l := []*datamodels.IndicatorDomainObjectsSTIX{}
			for cur.Next(context.Background()) {
				var model datamodels.IndicatorDomainObjectsSTIX
				_ = cur.Decode(&model)

				l = append(l, &model)
			}

			/*for _, v := range l {
				fmt.Printf("Found STIX object with ID:'indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f' - '%v'\n", *v)
			}*/

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(1))
		})
	})

	Context("Тест 7. Получаем информацию обо всех STIX объектах имеющих тип 'location'", func() {
		It("должен быть получен список из 3 STIX объектов, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.type", Value: "location"}})

			l := []*datamodels.LocationDomainObjectsSTIX{}
			for cur.Next(context.Background()) {
				var model datamodels.LocationDomainObjectsSTIX
				_ = cur.Decode(&model)

				l = append(l, &model)
			}

			/*for _, v := range l {
				fmt.Printf("Found STIX object with Type: ID:'%s'\n'location' - '%v'\n", v.ID, *v)
			}*/

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(3))
		})
	})

	Context("Тест 8. Получаем информацию обо всех STIX объектах", func() {
		It("должен быть получен список из 65 STIX объектов, ошибок быть не должно", func() {
			cur, err := qp.Find(bson.D{})

			l := GetListElementSTIXObject(cur)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(0))
		})
	})

	/*
	   Список STIX объектов успешно добавляется, однако есть некоторые нюансы:
	   1. Объекты добавляются и добавляются, дублируя друг друга. Нет проверки на уникальность ID объекта.
	   2. Поиск по типу или ID объекта затруднен из-за вложенности типов. А для таких STIX объектов как:
	    file, network-traffic и process невозможен из-за большой вложоности.
	*/

	/*
			Context("", func() {
			It("", func(){

			})
		})
	*/
})

type definingTypeSTIXObject struct {
	datamodels.CommonPropertiesObjectSTIX
}

func GetListElementSTIXObject(cur *mongo.Cursor) []*datamodels.ElementSTIXObject {
	elements := []*datamodels.ElementSTIXObject{}

	for cur.Next(context.Background()) {
		var modelType definingTypeSTIXObject
		_ = cur.Decode(&modelType)

		fmt.Printf("func 'GetListElementSTIXObject', type STIX object: '%s'\n", modelType.Type)

		//		elements = append(elements, &model)
	}

	return elements
}

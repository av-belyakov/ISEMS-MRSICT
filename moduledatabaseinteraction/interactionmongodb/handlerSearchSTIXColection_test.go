package interactionmongodb_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

var _ = Describe("HandlerSearchSTIXColection", func() {
	var (
		connectError error
		cdmdb        interactionmongodb.ConnectionDescriptorMongoDB
		qp           interactionmongodb.QueryParameters = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
		}
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

		qp.ConnectDB = cdmdb.Connection

	})

	var _ = AfterSuite(func() {
		cdmdb.CtxCancel()
	})

	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Выполнение поисковых запросов по коллекции STIX объектов", func() {
		It("Поиск только по списку Document ID, должно быть найдено определенное количество объектов", func() {
			ldid := []string{
				"intrusion-set--0c7e22ad-b099-4dc3-b0df-2ea3f49ae2e6",
				"indicator--d38a99ae-c5ee-4542-bc12-dfe68b48cc08",
				"location--a6e9345f-5a15-4c29-8bb3-8dac5d168d64",
			}

			sizeElem, err := qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
				DocumentsID: ldid,
			}))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(len(ldid))))
		})

		It("Поиск ТОЛЬКО по списку Document Type, должно быть найдено определенное количество объектов", func() {
			sizeElem, err := qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
				DocumentsType: []string{"grouping", "location", "report"},
			}))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(5)))
		})

		It("Поиск ТОЛЬКО по времени создания STIX объекта, должно быть найдено определенное количество объектов", func() {
			tcs, errsp := time.Parse(time.RFC3339, "2015-12-21T19:59:11.000Z")
			Expect(errsp).ShouldNot(HaveOccurred())

			tce, errep := time.Parse(time.RFC3339, "2016-08-21T21:34:10.000Z")
			Expect(errep).ShouldNot(HaveOccurred())

			qrotc := datamodels.SearchThroughCollectionSTIXObjectsType{
				Created: struct {
					Start time.Time "json:\"start\""
					End   time.Time "json:\"end\""
				}{
					Start: tcs,
					End:   tce,
				},
			}

			sizeElem, err := qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(19)))

			cur, errfl := qp.FindAllWithLimit(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc), &interactionmongodb.FindAllWithLimitOptions{
				Offset:        1,
				LimitMaxSize:  100,
				SortAscending: false,
			})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			/*fmt.Println("____ last 19 elements ___")
			for _, v := range elemSTIXObj {
				fmt.Printf("\tid: '%s'\n", v.Data.GetID())
			}*/

			Expect(errfl).ShouldNot(HaveOccurred())
			Expect(int64(len(elemSTIXObj))).Should(Equal(int64(19)))
		})

		It("Поиск ТОЛЬКО по времени модификации STIX объекта, должно быть найдено определенное количество объектов", func() {
			tms, errsp := time.Parse(time.RFC3339, "2016-05-01T19:59:11.000Z")
			Expect(errsp).ShouldNot(HaveOccurred())

			tme, errep := time.Parse(time.RFC3339, "2016-05-30T21:34:10.000Z")
			Expect(errep).ShouldNot(HaveOccurred())

			qrotc := datamodels.SearchThroughCollectionSTIXObjectsType{
				Modified: struct {
					Start time.Time "json:\"start\""
					End   time.Time "json:\"end\""
				}{
					Start: tms,
					End:   tme,
				},
			}
			sizeElem, err := qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(7)))
		})

		It("Поиск ТОЛЬКО по времени создания или модификации STIX объекта, должно быть найдено определенное количество объектов", func() {
			tcs, errsp := time.Parse(time.RFC3339, "2015-12-21T19:59:11.000Z")
			Expect(errsp).ShouldNot(HaveOccurred())

			tce, errep := time.Parse(time.RFC3339, "2016-08-21T21:34:10.000Z")
			Expect(errep).ShouldNot(HaveOccurred())

			tms, errsp := time.Parse(time.RFC3339, "2015-12-21T19:59:11.000Z")
			Expect(errsp).ShouldNot(HaveOccurred())

			tme, errep := time.Parse(time.RFC3339, "2016-08-21T21:34:10.000Z")
			Expect(errep).ShouldNot(HaveOccurred())

			qrotc := datamodels.SearchThroughCollectionSTIXObjectsType{
				Created: struct {
					Start time.Time "json:\"start\""
					End   time.Time "json:\"end\""
				}{
					Start: tcs,
					End:   tce,
				},
				Modified: struct {
					Start time.Time "json:\"start\""
					End   time.Time "json:\"end\""
				}{
					Start: tms,
					End:   tme,
				},
			}

			sizeElem, err := qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(20)))
		})

		It("Поиск по времени модификации STIX объекта и его типу, должно быть найдено определенное количество объектов", func() {
			tms, errsp := time.Parse(time.RFC3339, "2016-05-01T19:59:11.000Z")
			Expect(errsp).ShouldNot(HaveOccurred())

			tme, errep := time.Parse(time.RFC3339, "2016-05-30T21:34:10.000Z")
			Expect(errep).ShouldNot(HaveOccurred())

			qrotc := datamodels.SearchThroughCollectionSTIXObjectsType{
				DocumentsType: []string{"vulnerability"},
				Modified: struct {
					Start time.Time "json:\"start\""
					End   time.Time "json:\"end\""
				}{
					Start: tms,
					End:   tme,
				},
			}
			sizeElem, err := qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(1)))
		})
	})
})

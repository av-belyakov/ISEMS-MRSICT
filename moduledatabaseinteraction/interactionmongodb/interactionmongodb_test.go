package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

var _ = Describe("Interactionmongodb", Ordered, func() {
	var (
		connectError error
		cdmdb        interactionmongodb.ConnectionDescriptorMongoDB
		qp           interactionmongodb.QueryParameters = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
		}
	)

	BeforeAll(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		cdmdb = interactionmongodb.ConnectionDescriptorMongoDB{
			Ctx:       ctx,
			CtxCancel: cancel,
		}

		//подключаемся к базе данных MongoDB
		connectError = cdmdb.CreateConnection(&datamodels.MongoDBSettings{
			//Host:     "test-uchet-db.cloud.gcm",
			Host:     "127.0.0.1",
			Port:     27017,
			User:     "module-isems-mrsict",
			Password: "vkL6Znj$Pmt1e1",
			NameDB:   "isems-mrsict",
		})

		qp.ConnectDB = cdmdb.Connection
	})

	tempStorage := memorytemporarystoragecommoninformation.NewTemporaryStorage()
	//fmt.Println(tempStorage)

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
			Expect(sizeElem).Should(Equal(int64(90)))

			tcs, errep := time.Parse(time.RFC3339, "2021-09-28T09:00:00.000Z")
			tce, errep := time.Parse(time.RFC3339, "2021-09-28T21:00:00.000Z")
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

			sizeElem, err = qp.CountDocuments(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc))

			Expect(err).ShouldNot(HaveOccurred())
			Expect(sizeElem).Should(Equal(int64(38)))

			cur, errfl := qp.FindAllWithLimit(interactionmongodb.CreateSearchQueriesSTIXObject(&qrotc), &interactionmongodb.FindAllWithLimitOptions{
				Offset:        1,
				LimitMaxSize:  100,
				SortAscending: false,
			})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			//fmt.Println("SEARCH ELEMENTS = ", len(elemSTIXObj))
			//fmt.Println("____ should be 38 elements ___")
			//for k, v := range elemSTIXObj {
			//	fmt.Printf("\t%d. id: '%s'\n", k+1, v.Data.GetID())
			//}

			Expect(errfl).ShouldNot(HaveOccurred())
			Expect(int64(len(elemSTIXObj))).Should(Equal(int64(38)))
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
			Expect(sizeElem).Should(Equal(int64(52)))
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

	Context("Тест 3. Тестируем обработчики формирующие поисковые запросы, в формате BSON, направленные к БД MongoDB", func() {
		It("Должен быть успешно сформирован BSON запрос для обработки поля 'value'. При передачи запроса на обработку в БД ошибок быть не должно", func() {
			vr := interactionmongodb.HandlerValueField([]string{
				"https://example.com/research/index.html", // есть в коллекции БД	url--c1477287-23ac-5971-a010-5c287877fa60
				//"vdb76@yandex.ru",
				"lana_manata_mail@yahoo.com", // есть в коллекции БД	 email-addr--89f52ea8-d6ef-51e9-8fce-6a29236436ed
				//"john@example.com",
				"2001:db8:85a3::8a2e:370:7334", // есть в коллекции БД	ipv6-addr--1e61d36c-a16c-53b7-a80f-2a00161c96b1
				"91.123.6.66",                  // есть в коллекции БД	 ipv4-addr--2f372154-779c-43e0-9828-ecf8ed907d70
				//"basd-89@bk.com",
			})

			fmt.Printf("\tВ ЭТОМ ТЕСТЕ неверно отрабатывает поиск !!!! \nBSON document: '%v'\n", vr)

			cur, err := qp.FindAllWithLimit(bson.D{vr}, &interactionmongodb.FindAllWithLimitOptions{
				Offset:        1,
				LimitMaxSize:  100,
				SortField:     "",
				SortAscending: false,
			})

			fmt.Printf("\t------------------------ CUR: '%v'\n", cur)
			fmt.Println("Error: ", err)

			//
			//cur, err := qp.FindAllWithLimit(interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
			//	DocumentsType: []string{"grouping", "location", "report", "malware", "attack-pattern"},
			//	SpecificSearchFields: []datamodels.SpecificSearchFieldsSTIXObjectType{{
			//		Name: "sss",
			//		Value: []string{
			//			"https://talks.golang.org/2012/10things.slide#2",
			//			"vdb76@yandex.ru",
			//	Offset:        1,
			//	LimitMaxSize:  100,
			//	SortField:     "",
			//	SortAscending: false,
			//})
			//

			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			fmt.Printf("Found '%d' elements\n", len(elemSTIXObj))

			for _, v := range elemSTIXObj {

				fmt.Printf("Data type STIX: '%s', value: '%v'\n", v.DataType, v)

			}

			//
			// Вот тут поиск не отрабатывает как нужно, не должно быть 0 найденных
			//

			Expect(len(elemSTIXObj)).Should(Equal(4))
		})
	})

	/*Context("Тест 3.1. Тестируем обработчики формирующие поисковые запросы, в формате BSON, направленные к БД MongoDB", func() {
		It("Должен быть успешно сформирован BSON запрос для обработки поля 'value'. При передачи запроса на обработку в БД ошибок быть не должно", func() {
			sr := interactionmongodb.HandlerSpecificSearchFields([]string{""}, &datamodels.SpecificSearchFieldsSTIXObjectType{
				//Name:    "sss",
				//Aliases: []string{"df", "fe"},
				//FirstSeen: struct {
				//	Start time.Time "json:\"start\""
				//	End   time.Time "json:\"end\""
				//}{},
				//LastSeen: struct {
				//	Start time.Time "json:\"start\""
				//	End   time.Time "json:\"end\""
				//}{},
				Value: []string{
					"https://talks.golang.org/2012/10things.slide#2",
					"vdb76@yandex.ru",
					"john@example.com",                        // есть в коллекции БД
					"2001:0db8:85a3:0000:0000:8a2e:0370:7334", // есть в коллекции БД
					"198.51.100.3",                            // есть в коллекции БД
					"basd-89@bk.com",
				},
			})

			//sr := interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
			//	DocumentsType: []string{"report"},
			//	SpecificSearchFields: []datamodels.SpecificSearchFieldsSTIXObjectType{{
			//		Value: []string{
			//			"https://talks.golang.org/2012/10things.slide#2",
			//			"vdb76@yandex.ru",
			//			"john@example.com",                        // есть в коллекции БД
			//			"2001:0db8:85a3:0000:0000:8a2e:0370:7334", // есть в коллекции БД
			//			"198.51.100.3",                            // есть в коллекции БД
			//			"basd-89@bk.com",
			//		},
			//	}},
			//})

			//fmt.Printf("_________||||| %v |||||________\n", sr)

			_, err := qp.FindAllWithLimit(
				sr,
				&interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  100,
					SortField:     "",
					SortAscending: false,
				})

			Expect(err).ShouldNot(HaveOccurred())

			//elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			//			fmt.Printf("Found '%d' elements\n", len(elemSTIXObj))

			//			}
		})
	})*/

	Context("Тест 4. Формируем поисковые запросы с сортировкой и с выборкой по определенным полям", func() {
		It("Должен быть успешно выполнен запрос информации из БД и отсортирован по полю время создания (commonpropertiesdomainobjectstix.created)", func() {
			cur, err := qp.FindAllWithLimit(
				bson.D{}, &interactionmongodb.FindAllWithLimitOptions{
					SortField:     "commonpropertiesdomainobjectstix.created",
					SortAscending: false,
				})

			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			//fmt.Printf("Found '%d' elements and sorted\n", len(elemSTIXObj))
			//for _, v := range elemSTIXObj {
			//	fmt.Printf("Data type STIX: '%s', time created: '%v'\n", v.DataType, v.Data)
			//}

			//
			//   Сортировка работает, однако следует помнить что, сортировка может выполнятся только по одной из групп STIX DO,
			//   STIX CO или STIX RO или по всем сразу. Данный факт зафисит от типа сортируемого поля, для разных STIX объектов
			//   может быть использованно разное поле. Например, для STIX DO 'malware' поле 'commonpropertiesdomainobjectstix',
			//   а для STIX CO 'file' поле 'optionalcommonpropertiescyberobservableobjectstix'
			//

			Expect(len(elemSTIXObj)).Should(Equal(396))
		})
	})

	Context("Тест 5. Тестируем поиск STIX DO 'grouping' с определенными именами, перечень в списке", func() {
		It("Должен быть получен список из трех ID", func() {
			done := make(chan interface{})

			go func() {
				tmp := map[string]string{}
				configFileSettings := map[string]datamodels.StorageApplicationCommonListType{}

				//проверяем наличие файлов с дефолтными настройками приложения
				row, err := os.ReadFile("../../defaultsettingsfiles/settingsStatusesDecisionsMadeComputerThreats.json")
				Expect(err).ShouldNot(HaveOccurred())

				err = json.Unmarshal(row, &tmp)
				Expect(err).ShouldNot(HaveOccurred())

				for k, v := range tmp {
					configFileSettings[k] = datamodels.StorageApplicationCommonListType{Description: v}
				}

				listID, err := interactionmongodb.GetIDGroupingObjectSTIX(qp, configFileSettings)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(listID)).Should(Equal(3))

				close(done)
			}()

			Eventually(done, 0.5).Should(BeClosed())
		})
	})

	Context("Тест 6. Тестируем получение ID решений к компьютерным угрозам", func() {
		It("Должен быть получен ID типа 'successfully implemented computer threat'", func() {
			_, err := tempStorage.GetIDDecisionsMadeSuccessfully()

			//fmt.Printf("\tTYPE: 'successfully implemented computer threat' = %s\n", id)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Должен быть получен ID типа 'unsuccessfully computer threat'", func() {
			_, err := tempStorage.GetIDDecisionsMadeUnsuccessfully()

			//fmt.Printf("\tTYPE: 'unsuccessfully computer threat' = %s\n", id)

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Должен быть получен ID типа 'false positive'", func() {
			_, err := tempStorage.GetIDDecisionsMadeFalsePositive()

			//fmt.Printf("\tTYPE: 'false positive' = %s\n", id)

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 7. Тестируем поиск STIX DO 'grouping' с определенными именами, перечень в списке", func() {
		It("Должен быть получен список из 35 ID", func() {
			tmp := map[string]string{}
			configFileSettings := map[string]datamodels.StorageApplicationCommonListType{}

			//проверяем наличие файлов с дефолтными настройками приложения
			row, err := os.ReadFile("../../defaultsettingsfiles/settingsComputerThreatTypes.json")
			Expect(err).ShouldNot(HaveOccurred())

			err = json.Unmarshal(row, &tmp)
			Expect(err).ShouldNot(HaveOccurred())

			for k, v := range tmp {
				configFileSettings[k] = datamodels.StorageApplicationCommonListType{Description: v}
			}

			listID, err := interactionmongodb.GetIDGroupingObjectSTIX(qp, configFileSettings)
			Expect(err).ShouldNot(HaveOccurred())

			//fmt.Printf("Test 5. List ID grouping: '%v'\n", listID)

			Expect(tempStorage.SetListComputerThreat(listID)).ShouldNot(HaveOccurred())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(listID)).Should(Equal(35))
		})
	})

	Context("Тест 8. Тестируем поиск STIX DO типа 'report' как опубликованные так и нет", func() {
		It("Должен быть получен список из определенного количества элементов типа 'report', поле 'published' которых содержит пустое значение", func() {
			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType:        []string{"report"},
					SpecificSearchFields: []datamodels.SpecificSearchFieldsSTIXObjectType{{}},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  100,
					SortField:     "commonpropertiesdomainobjectstix.created",
					SortAscending: false,
				})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(elemSTIXObj)).Should(Equal(32))
		})

		It("Должен быть получен список из определенного количества элементов типа 'report', с заполненным полем 'published'", func() {
			t, errtp := time.Parse(time.RFC3339, "2016-01-01T00:00:01.000Z")
			Expect(errtp).ShouldNot(HaveOccurred())

			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType:        []string{"report"},
					SpecificSearchFields: []datamodels.SpecificSearchFieldsSTIXObjectType{{Published: t}},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  100,
					SortField:     "commonpropertiesdomainobjectstix.created",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			Expect(len(elemSTIXObj)).Should(Equal(9))
		})

		It("Должен быть получен список из определенного количества элементов типа 'report', 'grouping', 'location', с заполненным полем 'published'", func() {
			t, errtp := time.Parse(time.RFC3339, "2016-01-01T00:00:01.000Z")
			Expect(errtp).ShouldNot(HaveOccurred())

			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType:        []string{"grouping", "location", "report"},
					SpecificSearchFields: []datamodels.SpecificSearchFieldsSTIXObjectType{{Published: t}},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  100,
					SortField:     "commonpropertiesdomainobjectstix.created",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)
			listType := map[string]int{}
			for _, v := range elemSTIXObj {
				objType := v.Data.GetType()

				if num, ok := listType[objType]; ok {
					listType[objType] = num + 1
				} else {
					listType[objType] = 1
				}
			}

			//fmt.Println("____ COUNT ELEMENT TYPES____")
			//for k, v := range listType {
			//	fmt.Printf("Element type: '%s' = %d\n", k, v)
			//}
			//fmt.Println("____________________________")

			Expect(len(elemSTIXObj)).Should(Equal(90))
			Expect(listType["grouping"]).Should(Equal(43))
			Expect(listType["location"]).Should(Equal(6))
			Expect(listType["report"]).Should(Equal(41))
		})

		It("Должен быть получен STIX объект типа 'grouping' с заданным именем в поле Name", func() {
			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType:        []string{"grouping"},
					SpecificSearchFields: []datamodels.SpecificSearchFieldsSTIXObjectType{{Name: "unsuccessfully computer threat"}},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  100,
					SortField:     "",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			Expect(len(elemSTIXObj)).Should(Equal(1))
		})
	})

	Context("Тест 9. Поиск по типу STIX объектов и нестандартным полям", func() {
		It("Должен быть найден STIX объект с заданным содержимым нестандартного поля outside_specification.computer_threat_type", func() {
			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType: []string{"report"},
					OutsideSpecificationSearchFields: datamodels.OutsideSpecificationSearchFieldsType{
						//DecisionsMadeComputerThreat: "",
						//DecisionsMadeComputerThreat: "successfully implemented computer threat",
						ComputerThreatType: "SQL-Injection",
					},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  200,
					SortField:     "",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			//
			//fmt.Println("Test 9.1:")
			//for k, v := range elemSTIXObj {
			//	fmt.Printf("%d. STIX object type: '%s', value: '%v'\n%s\n", k, v.DataType, v.Data.GetID(), v.Data.ToStringBeautiful())
			//}
			//

			Expect(len(elemSTIXObj)).Should(Equal(1))
		})

		It("Должен быть найден STIX объект с заданным содержимым нестандартного поля outside_specification.decisions_made_computer_threat", func() {
			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType: []string{"report"},
					OutsideSpecificationSearchFields: datamodels.OutsideSpecificationSearchFieldsType{
						DecisionsMadeComputerThreat: "successfully implemented computer threat",
						ComputerThreatType:          "",
					},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  200,
					SortField:     "",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			//
			//fmt.Println("Test 9.2:")
			//for k, v := range elemSTIXObj {
			//	fmt.Printf("%d. STIX object type: '%s', value: '%v'\n%s\n", k, v.DataType, v.Data.GetID(), v.Data.ToStringBeautiful())
			//}
			//

			Expect(len(elemSTIXObj)).Should(Equal(10))
		})

		It("Должен быть найден STIX объект с заданным содержимым нестандартного поля outside_specification.decisions_made_computer_threat и outside_specification.decisions_made_computer_threat", func() {
			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(&datamodels.SearchThroughCollectionSTIXObjectsType{
					DocumentsType: []string{"report"},
					OutsideSpecificationSearchFields: datamodels.OutsideSpecificationSearchFieldsType{
						DecisionsMadeComputerThreat: "successfully implemented computer threat",
						ComputerThreatType:          "malware",
					},
				}), &interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  200,
					SortField:     "",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)

			//
			//	fmt.Println("Test 9.3:")
			//	for k, v := range elemSTIXObj {
			//		fmt.Printf("%d. STIX object type: '%s', value: '%v'\n", k, v.DataType, v.Data.GetID())
			//	}
			//

			Expect(len(elemSTIXObj)).Should(Equal(10))
		})
	})

	Context("Тест 10. Поиск по типу 'report' STIX объектов и нестандартным полям типов и статусов компьютерных угроз", func() {
		It("Должно быть получена статистическая информация о статусах компьютерных атак", func() {
			collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
			opts := options.Aggregate().SetAllowDiskUse(true).SetMaxTime(2 * time.Second)

			cur, err := collection.Aggregate(
				context.TODO(),
				mongo.Pipeline{
					bson.D{bson.E{Key: "$match", Value: bson.D{
						bson.E{Key: "commonpropertiesobjectstix.type", Value: "report"},
					}}},
					bson.D{
						bson.E{Key: "$group", Value: bson.D{
							bson.E{Key: "_id", Value: "$outside_specification.decisions_made_computer_threat"},
							bson.E{Key: "count", Value: bson.D{
								bson.E{Key: "$sum", Value: 1},
							}},
						}}}},
				opts)

			Expect(err).ShouldNot(HaveOccurred())

			var results []bson.M
			err = cur.All(context.TODO(), &results)

			//fmt.Printf("||| RESULT (decisions_made_computer_threat) |||\n'%v'\n", results)
			//for k, v := range results {
			//	name, ok := v["_id"].(string)
			//	if !ok {
			//		fmt.Println("Convert 'Name' ERROR")
			//	}

			//	count, ok := v["count"].(int32)
			//	if !ok {
			//		fmt.Println("Convert 'Count' ERROR")
			//	}

			//	fmt.Printf("%d. Name:%s, Count:%s\n", k+1, name, fmt.Sprintln(count))
			//}

			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Должно быть получена информация о типах компьютерных атак", func() {
			collection := qp.ConnectDB.Database(qp.NameDB).Collection(qp.CollectionName)
			opts := options.Aggregate().SetAllowDiskUse(true) //.SetMaxTime(2 * time.Second)

			cur, err := collection.Aggregate(
				context.TODO(),
				mongo.Pipeline{
					bson.D{bson.E{Key: "$match", Value: bson.D{
						bson.E{Key: "commonpropertiesobjectstix.type", Value: "report"},
					}}},
					bson.D{
						bson.E{Key: "$group", Value: bson.D{
							bson.E{Key: "_id", Value: "$outside_specification.computer_threat_type"},
							bson.E{Key: "count", Value: bson.D{
								bson.E{Key: "$sum", Value: 1},
							}},
						}}}},
				opts)

			Expect(err).ShouldNot(HaveOccurred())

			var results []bson.M
			err = cur.All(context.TODO(), &results)

			//fmt.Printf("||| RESULT (computer_threat_type) |||\n'%v'\n", results)
			//for k, v := range results {
			//	fmt.Printf("%d. ID:%s, Count:%v\n", k+1, v["_id"], v["count"])
			//}

			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	AfterAll(func() {
		cdmdb.CtxCancel()
	})
})

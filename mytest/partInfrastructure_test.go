package mytest_test

import (
	"context"
	"fmt"
	"time"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PartInfrastructure", func() {
	var (
		connectError error
		cdmdb        interactionmongodb.ConnectionDescriptorMongoDB
		qp           interactionmongodb.QueryParameters = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
		}
		tst *memorytemporarystoragecommoninformation.TemporaryStorageType
	)

	var _ = BeforeSuite(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		cdmdb = interactionmongodb.ConnectionDescriptorMongoDB{
			Ctx:       ctx,
			CtxCancel: cancel,
		}

		//подключаемся к базе данных MongoDB
		connectError = cdmdb.CreateConnection(&datamodels.MongoDBSettings{
			Host:     "127.0.0.1",
			Port:     27017,
			User:     "module-isems-mrsict",
			Password: "vkL6Zn$jPmt1e1",
			NameDB:   "isems-mrsict",
		})

		qp.ConnectDB = cdmdb.Connection
		//инициализируем временное хранилище
		tst = memorytemporarystoragecommoninformation.NewTemporaryStorage()
	})

	var _ = AfterSuite(func() {
		cdmdb.CtxCancel()
	})

	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Тестирование части инфраструктуры, начиная с обработки поискового запроса и до формирования JSON ответа пользователю", func() {
		It("В результате выполнения 'пустого' запроса к поисковой машине должно быть полученно 7 частей JSON сообщений, отправляемых пользователю", func() {
			const appTaskID string = "38g8g48t8y59g5h9g7g75g4884"

			cur, errfl := qp.FindAllWithLimit(interactionmongodb.CreateSearchQueriesSTIXObject(
				&datamodels.SearchThroughCollectionSTIXObjectsType{}),
				&interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  100,
					SortAscending: false,
				})

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)
			Expect(errfl).ShouldNot(HaveOccurred())
			Expect(len(elemSTIXObj)).Should(Equal(65))

			//сохраняем найденные значения во временном хранилище
			erranfi := tst.AddNewFoundInformation(
				appTaskID,
				&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
					Collection:  "stix_object_collection",
					ResultType:  "full_found_info",
					Information: elemSTIXObj,
				})
			//сохраняем найденные значения во временном хранилище
			/*erranfi := tst.AddNewFoundInformation(
			appTaskID,
			&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
				Collection:  "stix_object_collection",
				ResultType:  "only_count",
				Information: int64(len(elemSTIXObj)),
			})*/
			Expect(erranfi).ShouldNot(HaveOccurred())

			errhsr := handlingSearchRequestsSTIXObject(
				10,
				&datamodels.ModuleDataBaseInteractionChannel{
					CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{},
					Section:                             "stix object test",
					Command:                             "",
					AppTaskID:                           appTaskID,
				},
				tst,
				&memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType{
					TemporaryStorageTaskType: memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
						TaskStatus: "completed",
						TaskParameters: datamodels.ModAPIRequestProcessingResJSONSearchReqType{
							CollectionName: "stix object",
						},
					},
				})
			Expect(errhsr).ShouldNot(HaveOccurred())

			_, errgfibid := tst.GetFoundInformationByID(appTaskID)
			Expect(errgfibid).Should(HaveOccurred())
		})
	})
})

func handlingSearchRequestsSTIXObject(
	maxChunkSize int,
	data *datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {
	if ti.TaskStatus != "completed" {
		return nil
	}

	tp, ok := ti.TaskParameters.(datamodels.ModAPIRequestProcessingResJSONSearchReqType)
	if !ok {
		return fmt.Errorf("type conversion error, line 166")
	}

	//обрабатываем результаты опираясь на типы коллекций
	if tp.CollectionName == "stix object" {
		//делаем запрос к временному хранилищу информации
		result, err := tst.GetFoundInformationByID(data.AppTaskID)
		if err != nil {
			return err
		}

		taskID, di, err := tst.GetTaskByID(data.AppTaskID)
		if err != nil {
			return err
		}

		fmt.Printf("func 'handlerDataBaseResponse', Found Information: '%v'\n", result)

		msgRes := datamodels.ModAPIRequestProcessingResJSON{
			ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
				TaskID:  di.ClientTaskID,
				Section: data.Section,
			},
			IsSuccessful: true,
		}

		//для КРАТКОЙ информации, только колличество, по найденным STIX объектам
		if result.Collection == "stix_object_collection" && result.ResultType == "only_count" {
			numFound, ok := result.Information.(int64)
			if !ok {
				return fmt.Errorf("type conversion error, line 191")
			}

			msgRes.AdditionalParameters = struct {
				NumberDocumentsFound int64 `json:"number_documents_found"`
			}{
				NumberDocumentsFound: numFound,
			}

			fmt.Printf("func 'handlingSearchRequestsSTIXObject', Collection: 'stix_object_collection', ResultType: 'only_count', result: '%v'\n", msgRes)
		}

		//для ПОЛНОЙ информации по найденным STIX объектам
		if result.Collection == "stix_object_collection" && result.ResultType == "full_found_info" {
			listElemSTIXObj, ok := result.Information.([]*datamodels.ElementSTIXObject)
			if !ok {
				return fmt.Errorf("type conversion error, line 220")
			}

			sestixo := len(listElemSTIXObj)
			listMsgRes := make([]interface{}, 0, sestixo)
			for _, v := range listElemSTIXObj {
				listMsgRes = append(listMsgRes, v.Data)
			}

			//обрабатываем полученный список STIX объектов, в том числе если он превышает размер в 100 объектов
			if sestixo < maxChunkSize {
				msgRes.AdditionalParameters = datamodels.ResJSONParts{
					TotalNumberParts:      1,
					GivenSizePart:         maxChunkSize,
					NumberTransmittedPart: 1,
					TransmittedData:       listMsgRes,
				}

				fmt.Printf("func 'handlingSearchRequestsSTIXObject', Collection: 'stix_object_collection', ResultType: 'full_found_info', ALL result: '%v'\n", msgRes)
			} else {
				num := commonlibs.GetCountChunk(int64(sestixo), maxChunkSize)

				min := 0
				max := maxChunkSize
				for i := 0; i < num; i++ {
					data := datamodels.ResJSONParts{
						TotalNumberParts:      num,
						GivenSizePart:         maxChunkSize,
						NumberTransmittedPart: i + 1,
					}

					if i == 0 {
						data.TransmittedData = listMsgRes[:max]
					} else if i == num-1 {
						data.TransmittedData = listMsgRes[min:]
					} else {
						data.TransmittedData = listMsgRes[min:max]
					}

					min = min + maxChunkSize
					max = max + maxChunkSize

					msgRes.AdditionalParameters = data

					fmt.Printf("func 'handlingSearchRequestsSTIXObject', Collection: 'stix_object_collection', ResultType: 'full_found_info', result NUMDER %d: '%v'\n", i+1, msgRes)
				}
			}
		}

		//удаляем задачу и результаты поиска информации, если они есть
		tst.DeletingTaskByID(data.AppTaskID)
		tst.DeletingFoundInformationByID(data.AppTaskID)
	}

	return nil
}

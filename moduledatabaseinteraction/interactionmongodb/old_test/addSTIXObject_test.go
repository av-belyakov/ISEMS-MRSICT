package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ISEMS-MRSICT/commonhandlers"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

func getListSettings(f string, appConfig *datamodels.AppConfig) (map[string]datamodels.StorageApplicationCommonListType, error) {
	tmp := map[string]string{}
	configFileSettings := map[string]datamodels.StorageApplicationCommonListType{}

	//проверяем наличие файлов с дефолтными настройками приложения
	row, err := os.ReadFile(path.Join(appConfig.RootDir, f))
	if err != nil {
		return configFileSettings, fmt.Errorf("Error! The file '%s' with default settings not found.", f)
	}

	err = json.Unmarshal(row, &tmp)
	if err != nil {
		return configFileSettings, err
	}

	for k, v := range tmp {
		configFileSettings[k] = datamodels.StorageApplicationCommonListType{Description: v}
	}

	return configFileSettings, err
}

var _ = Describe("AddSTIXObject", Ordered, func() {
	var (
		docJSON []byte
		l       []*datamodels.ElementSTIXObject
		qp      interactionmongodb.QueryParameters = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
		}
		qpdo                                                           interactionmongodb.QueryParameters
		connectError, errReadFile, errUnmarchalReq, errUnmarchalToSTIX error
		cdmdb                                                          interactionmongodb.ConnectionDescriptorMongoDB
		modAPIRequestProcessingReqJSON                                 datamodels.ModAPIRequestProcessingReqJSON
		tst                                                            *memorytemporarystoragecommoninformation.TemporaryStorageType
		appConfig                                                      datamodels.AppConfig
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

		docJSON, errReadFile = os.ReadFile("../../mytest/test_resources/jsonSTIXExampleObjectRef.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)

		qpdo = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "accounting_differences_objects_collection",
			ConnectDB:      cdmdb.Connection,
		}

		appConfig.RootDir = "/Users/user/go/src/ISEMS-MRSICT"
		tst = memorytemporarystoragecommoninformation.NewTemporaryStorage()

		ssdmct, _ := getListSettings("defaultsettingsfiles/settingsStatusesDecisionsMadeComputerThreats.json", &appConfig)
		tst.SetListDecisionsMade(ssdmct)

		//fmt.Println(ssdmct)

		sctt, _ := getListSettings("defaultsettingsfiles/settingsComputerThreatTypes.json", &appConfig)
		tst.SetListComputerThreat(sctt)

		//fmt.Println(sctt)
	})

	AfterAll(func() {
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

	Context("Тест 3. Преобразуем срез байт в тип datamodels.ModAPIRequestProcessingReqJSON", func() {
		It("При преобразовании среза байт ошибок быть не должно", func() {
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Взаимодействие с коллекцией STIX объектов.", func() {
		It("При добавлении STIX объектов не должно быть ошибок. STIX объекты идентификаторы которых уже есть в БД добавлятся не должны.", func() {
			//получаем список ID STIX объектов предназначенных для добавление в БД
			listID := commonhandlers.GetListIDFromListSTIXObjects(l)

			fmt.Println("-----------=====================================================----------")

			countBefore := len(l)

			//выполняем запрос к БД, для получения полной информации об STIX объектах по их ID
			listElemetSTIXObject, err := interactionmongodb.FindSTIXObjectByID(qp, listID)
			Expect(err).ShouldNot(HaveOccurred())

			routingflowsmoduleapirequestprocessing.VerifyOutsideSpecificationFields(l, tst, "client-test")
			l := interactionmongodb.SavingAdditionalNameListSTIXObject(listElemetSTIXObject, l)

			//выполняем сравнение объекта из БД и полученного из файла
			//выполняем сравнение объектов и ищем внесенные изменения для каждого из STIX объектов
			listDifferentObject := interactionmongodb.ComparasionListSTIXObject(interactionmongodb.ComparasionListTypeSTIXObject{
				CollectionType: qp.CollectionName,
				OldList:        listElemetSTIXObject,
				NewList:        l,
			})

			//fmt.Println("only listDifferentObject")
			//fmt.Println(listDifferentObject)

			if len(listDifferentObject) > 0 {
				list := make([]interface{}, 0, len(listDifferentObject))

				for _, v := range listDifferentObject {
					list = append(list, v)
				}

				_, err = qpdo.InsertData(list, []mongo.IndexModel{
					{
						Keys: bson.D{
							{Key: "document_id", Value: 1},
						},
						Options: &options.IndexOptions{},
					},
				})
				Expect(err).ShouldNot(HaveOccurred())
			}

			countAfter := len(l)

			err = interactionmongodb.ReplacementElementsSTIXObject(qp, l)

			fmt.Println("-----------=====================================================----------")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(countBefore).Should(Equal(countAfter))
		})
	})

	Context("Тест 5. Проверяем функцию 'UnmarshalJSONObjectSTIXReq'", func() {
		It("Должен быть получен список из 66 STIX объектов, ошибок быть не должно", func() {
			Expect(errUnmarchalToSTIX).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(34))
		})
	})
})

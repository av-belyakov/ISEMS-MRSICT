package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
	row, err := ioutil.ReadFile(path.Join(appConfig.RootDir, f))
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

var _ = Describe("AddSTIXObjSetBackLink", func() {
	var (
		docJSON                                                        []byte
		l                                                              []*datamodels.ElementSTIXObject
		qp                                                             interactionmongodb.QueryParameters
		connectError, errReadFile, errUnmarchalReq, errUnmarchalToSTIX error
		cdmdb                                                          interactionmongodb.ConnectionDescriptorMongoDB
		modAPIRequestProcessingReqJSON                                 datamodels.ModAPIRequestProcessingReqJSON
		tst                                                            *memorytemporarystoragecommoninformation.TemporaryStorageType
		appConfig                                                      datamodels.AppConfig
	)

	var _ = BeforeSuite(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

		cdmdb = interactionmongodb.ConnectionDescriptorMongoDB{
			Ctx:       ctx,
			CtxCancel: cancel,
		}

		//подключаемся к базе данных MongoDB
		connectError = cdmdb.CreateConnection(&datamodels.MongoDBSettings{
			Host:     "test-uchet-db.cloud.gcm",
			Port:     27017,
			User:     "module-isems-mrsict",
			Password: "vkL6Zn$jPmt1e1",
			NameDB:   "isems-mrsict",
		})

		docJSON, errReadFile = ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExampleObjectRef.json")
		errUnmarchalReq = json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON)
		l, errUnmarchalToSTIX = routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)

		qp = interactionmongodb.QueryParameters{
			NameDB:         "isems-mrsict",
			CollectionName: "stix_object_collection",
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

	Context("Тест 3. Преобразуем срез байт в тип datamodels.ModAPIRequestProcessingReqJSON", func() {
		It("При преобразовании среза байт ошибок быть не должно", func() {
			Expect(errUnmarchalReq).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 4. Проверяем функцию 'UnmarshalJSONObjectSTIXReq'", func() {
		It("Должен быть получен список из 66 STIX объектов, ошибок быть не должно", func() {
			Expect(errUnmarchalToSTIX).ShouldNot(HaveOccurred())
			Expect(len(l)).Should(Equal(33))
		})
	})

	Context("Тест 5. Взаимодействие с коллекцией STIX объектов.", func() {
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

			countAfter := len(l)

			err = interactionmongodb.ReplacementElementsSTIXObject(qp, l)

			fmt.Println("-----------=====================================================----------")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(countBefore).Should(Equal(countAfter))
		})
	})
})

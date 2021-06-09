package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

func addListTestSTIXObject(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters) error {
	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON

	docJSON, err := ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExample_1.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON); err != nil {
		return err
	}

	l, err := routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)
	if err != nil {
		return err
	}

	if err := interactionmongodb.ReplacementElementsSTIXObject(qp, l); err != nil {
		return err
	}

	return nil
}

func delListTestSTIXObject(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters) error {
	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON

	docJSON, err := ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExample_1.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(docJSON, &modAPIRequestProcessingReqJSON); err != nil {
		return err
	}

	l, err := routingflowsmoduleapirequestprocessing.UnmarshalJSONObjectSTIXReq(modAPIRequestProcessingReqJSON)
	if err != nil {
		return err
	}

	listDelID := make([]string, 0, len(l))
	for _, v := range l {
		listDelID = append(listDelID, v.Data.GetID())
	}

	if _, err := qp.DeleteManyData(listDelID); err != nil {
		return err
	}

	return nil
}

func deleteObjTypeGrouping(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters, listID []string) error {

	return nil
}

var _ = Describe("HandlerDeleteSTIXObject", func() {
	var (
		connectError, errAddListObj error
		cdmdb                       interactionmongodb.ConnectionDescriptorMongoDB
		//tempStorage                 *memorytemporarystoragecommoninformation.TemporaryStorageType
		qp interactionmongodb.QueryParameters = interactionmongodb.QueryParameters{
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

		//tempStorage = memorytemporarystoragecommoninformation.NewTemporaryStorage()

		if connectError == nil {
			errAddListObj = addListTestSTIXObject(cdmdb, qp)
		}
	})

	var _ = AfterSuite(func() {
		cdmdb.CtxCancel()

		//_ = delListTestSTIXObject(cdmdb, qp)
	})

	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Добавление тестового набора STIX объектов (типы 'report', 'relationship' и 'grouping')", func() {
		It("При добавления тестового набора ошибок быть не должно", func() {
			Expect(errAddListObj).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Выполняем поиск и удаление объектов типа 'grouping' и связующих их объектов типа 'relationship'", func() {
		It("При удалении объектов 'grouping' ошибок быть недолжно", func() {

			/* НЕДОПИСАЛ ФУНКЦИЮ НАДО ДОДЕЛАТЬ */
			Expect(deleteObjTypeGrouping(cdmdb, qp, []string{
				"grouping--911f0f43-1e7c-4f97-7070-a787bdd3b100",
				"grouping--911f0f23-2e7c-4f07-9436-d6e7bdd3b236",
			})).ShouldNot(HaveOccurred())
		})

		It("Должны быть удалены 4 объекта", func() {

		})
	})
})

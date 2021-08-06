package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

func addListTestSTIXObject(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters) error {
	/*
						Описание тестового файла jsonSTIXExample_2.json
				Файл содержит 5 объекта: 1-'report', 2-'grouping', 1-'note', 1-'observed-data' при этом свойства ObjectRefs этих объектов содержит:
				1. type: "report" //"report--94e4d99f-67aa-4bcd-bbf3-b2c1c320aad7"
				"object_refs": [
		                "indicator--d38a99ae-c5ee-4542-bc12-dfe68b48cc08", //ссылается на объект в БД
		                "campaign--ce88a5a8-69ff-4349-86da-ac59b35c5672", //ссылается на объект в БД
		                "grouping--23af0f32-2e61-4f00-7020-b378bdd3b201", //ссылается на объект тестовом файле
		                "grouping--11ff0f44-1e8b-5607-ab36-ce89b1d11ce6" //ссылается на объект тестовом файле
		            ]
				2. type: "grouping" //"grouping--23af0f32-2e61-4f00-7020-b378bdd3b201"
				"object_refs": ["network-traffic--e7a939ca-78c6-5f27-8ae0-4ad112454626"] //ссылается на объект в БД
				3. type: "grouping" //"grouping--11ff0f44-1e8b-5607-ab36-ce89b1d11ce6"
				"object_refs": ["process--07bc30cad-ebc2-4579-881d-b9cdc7f2b33c"] //ссылается на объект в БД
				4. type: "note" //"note--addb1322-e64c-41cf-1031-acd6cbe3c12b"
				"object_refs": ["observed-data--c67d11ab-30ab-410a-abf9-10d888f412da"] //ссылается на объект тестовом файле
				5. type: "observed-data" //"observed-data--c67d11ab-30ab-410a-abf9-10d888f412da"
				"object_refs": [
		                "ipv4-addr--5853f6a4-638f-5b4e-9b0f-ded361ae3812", //ссылается на объект в БД
		                "domain-name--ecb120bf-2694-4902-a737-62b74539a41b" //ссылается на объект в БД
		            ]
				6. type: "relationship" //"relationship--57b56a43-b8b0-4cba-9deb-34e3e1faed9e"
				source_ref: "grouping--11ff0f44-1e8b-5607-ab36-ce89b1d11ce6",
				target_ref: "report--94e4d99f-67aa-4bcd-bbf3-b2c1c320aad7",

	*/

	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON

	docJSON, err := ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExample_2.json")
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

	//newList, err := creatingAdditionalRelationshipSTIXObject(l)
	newList, err := interactionmongodb.CreatingAdditionalRelationshipSTIXObject(qp, l)
	if err != nil {
		return err
	}

	fmt.Println("---+++=== ___newList___")
	for _, v := range newList {
		fmt.Printf("---+++=== ___newList___: Type: '%s', ID: '%s'\n", v.DataType, v.Data.GetID())
	}

	//запись в БД полученной от пользователя информации и дополнительно сформированных объектов типа 'relationship'
	if err := interactionmongodb.ReplacementElementsSTIXObject(qp, newList); err != nil {
		return err
	}

	/*
		создание объектов типа 'relationship', обеспечивающих обратные связи, выполняется успешно
		все дополнительно записывается в БД. Теперь необходимо заниматся удалением объектов типа 'relationship', обеспечивающих обратные связи
		которое должно выполнятся при исключении ID объекта из свойства 'object_ref'
	*/

	return nil
}

func delListTestSTIXObject(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters) error {
	/*
		Так же может возникнуть ситуация когда из свойства ObjectRefs объектов типа:
			- 'grouping'
			- 'report'
			- 'note'
			- 'observed-data'
			- 'opinion'
			были удалены ID некоторых объектов. Тогда нужно найти все объекты типа 'relationship' где в свойстве target_ref будет равно ID одного из
			объектов типа:
			- 'grouping'
			- 'report'
			- 'note'
			- 'observed-data'
			- 'opinion'
			то есть объекта из свойства ObjectRef которого были удалены ID какого то из объектов STIX. Далее надо сравнить список ID из свойства
			ObjectRef и список ID ссылок в полученный при чтении свойства source_ref всех найденных объектов типа 'relationship'. Если в свойстве
			source_ref ID объекта есть (для данного target_ref ID совпадающего с ID одного из пяти выше перечисленных объектов), а в объекте с полем
			ObjectRef данного ID нет, то это и будет лишним объектом типа 'relationship'. Добавляем его ID в список 'relationship' объектов
			подлежащих удалению

			1. Прочитать тестовый файл ../../mytest/test_resources/jsonSTIXExample_3.json (надо создать)

			2. проверить наличие списока объектов типа: 'grouping', 'report', 'note', 'observed-data', 'opinion' в БД по их ID
			полученных от пользователя, если объекты с такими ID уже есть в базе нужно сравнить содержимое поля 'object_ref'
			объекта полученного из БД и полученного от пользователя. Если в поле 'object_ref' объекта полученного из БД есть ссылка,
			а в том же объекте полученном от пользователя ее уже нет, то нужно выполнить поиск и удаление объекта типа 'relationship'
			в поле 'target_ref' которого ID одного из объектов типа 'grouping', 'report', 'note', 'observed-data', 'opinion',
			а в поле 'source_ref' которого ID объекта отсутствующего в объекте полученном от пользователя

			Описание тестового файла jsonSTIXExample_3.json
				Файл содержит 1 объект: 1-'report' при этом свойства ObjectRefs этих объектов содержит:
				1. type: "report" //"report--94e4d99f-67aa-4bcd-bbf3-b2c1c320aad7"
				"object_refs": [
		                "indicator--d38a99ae-c5ee-4542-bc12-dfe68b48cc08", //ссылается на объект в БД (БЫЛО УДАЛЕНО ИЗ СПИСКА!!!)
		                "campaign--ce88a5a8-69ff-4349-86da-ac59b35c5672", //ссылается на объект в БД
		                "grouping--23af0f32-2e61-4f00-7020-b378bdd3b201", //ссылается на объект тестовом файле (БЫЛО УДАЛЕНО ИЗ СПИСКА!!!)
		                "grouping--11ff0f44-1e8b-5607-ab36-ce89b1d11ce6" //ссылается на объект тестовом файле
		            ]
	*/

	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON

	docJSON, err := ioutil.ReadFile("../../mytest/test_resources/jsonSTIXExample_3.json")
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

	//return deleteOldRelationshipSTIXObject(l)
	return interactionmongodb.DeleteOldRelationshipSTIXObject(qp, l)
}

var _ = Describe("AddSTIXObjSetBackLink", func() {
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
	})

	Context("Тест 1. Проверка наличия установленного соединения с БД", func() {
		It("При установления соединения с БД ошибки быть не должно", func() {
			Expect(connectError).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Добавление тестового набора STIX объектов и установление дополнительных связей", func() {
		It("При добавления тестового набора ошибок быть не должно", func() {
			Expect(errAddListObj).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Удаление объектов типа 'relationship' обеспечивающих обратные связи, при удалении ID объекта из свойства 'object_ref'", func() {
		It("При удалении объектов типа 'relationship' ошибок быть не должно", func() {
			err := delListTestSTIXObject(cdmdb, qp)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

package interactionmongodb_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

func addListTestSTIXObject(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters) error {
	/*
		Нужно сделать автоматическое установление ОБРАТНЫХ связей между STIX объектами содержащими свойство ObjectRefs, такими как объекты типов:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		и с любыми другими объектами по средствам объектов типа 'relationship'.
		Нужно сделать автоматическое удаление объектов типа 'relationship' обеспечивающие обратную связь между объектами типов 'grouping'
		и 'report' и еще три выше перечисленных и другими объектами при удалении ссылок на объекты из поля ObjectRefs объектов типов
		'grouping' и 'report' и т.д.

		1. Создать новый тестовый набор STIX объектов, в нем должны быть объекты типа:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		только они имеют свойство ObjectRefs в котором есть ссылки на другие объекты

		2. Получить из этих объектов все ссылки содержащиеся в свойстве ObjectRefs в виде карты, где ключ будет являтся
		ID объекта типа:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'

		3. Выполнить поиск по принятым в запросе объектам, есть ли в них объекты, ID которых совпадают с ID в свойстве ObjectRefs
		какого либо из объектов типа:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		и создать список объектов типа 'relationship' обеспечивающих обратные связи с этими объектами

		4. Выполнить поиск объектов типа 'relationship' в БД где свойство target_ref будет равно ID одного из объектов типа:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		а свойство source_ref будет равно одному из ID в свойстве ObjectRefs данного объекта. Создать список объектов типа
		'relationship' обеспечивающих обратные связи с этими объектами. Это делается потому что принятый запрос может не содержать
		объекты с котороыми устанавливается связи при добавлении или изменении объектов вышеперечисленных объектов

		5. Так же может возникнуть ситуация когда из свойства ObjectRefs объектов типа:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		были удалены ID некоторых объектов. Тогда нужно найти все объекты типа 'relationship' где в свойстве target_ref будет равно ID одного из
		объектов типа:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		то есть объекта из свойства ObjectRef которого были удалены ID какого то из объектов STIX. Далее надо сравнить список ID из свойства
		ObjectRef и список ID ссылок в полученный при чтении свойства source_ref всех найденных объектов типа 'relationship'. Если в свойстве
		source_ref ID объекта есть (для данного target_ref ID совпадающего с ID одного из пяти выше перечисленных объектов), а в объекте с полем
		ObjectRef данного ID нет, то это и будет лишним объектом типа 'relationship'. Добавляем его ID в список 'relationship' объектов
		подлежащих удалению

	*/

	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON

	/*
						Описание тестового файла jsonSTIXExample_2.json
				Файл содержит 5 объекта: 1-'report', 2-'grouping', 1-'note', 1-'observed-data' при этом свойства ObjectRefs этих объектов содержит:
				1. type: "report"
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
				4. type: "note"
				"object_refs": ["observed-data--c67d11ab-30ab-410a-abf9-10d888f412da"] //ссылается на объект тестовом файле
				5. type: "observed-data"
				"object_refs": [
		                "ipv4-addr--5853f6a4-638f-5b4e-9b0f-ded361ae3812", //ссылается на объект в БД
		                "domain-name--ecb120bf-2694-4902-a737-62b74539a41b" //ссылается на объект в БД
		            ]

	*/
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

	if err := interactionmongodb.ReplacementElementsSTIXObject(qp, l); err != nil {
		return err
	}

	return nil
}

func delListTestSTIXObject(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters, listID []string) error {
	/*var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON

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
	}*/

	if _, err := qp.DeleteManyData(bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: listID}}}}); err != nil {
		return err
	}

	return nil
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
		//удаляем только добавленный объект типа 'report', остальные добавленные объекты должны быть удалены ранее функцией deleteObjTypeGrouping
		/*err := delListTestSTIXObject(cdmdb, qp, []string{"report--94e4d99f-67aa-4bcd-bbf3-b2c1c320aad7"})
		if err != nil {
			fmt.Println(err)
		}*/

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
})

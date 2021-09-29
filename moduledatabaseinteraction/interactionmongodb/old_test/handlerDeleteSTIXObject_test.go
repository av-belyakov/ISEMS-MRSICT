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

func deleteObjTypeGrouping(cdmdb interactionmongodb.ConnectionDescriptorMongoDB, qp interactionmongodb.QueryParameters, listID []string) error {
	var (
		listIDGroupingDel     []string // список объектов типа 'grouping' предназначеных для удаления
		listIDRelationshipDel []string // список объектов типа 'relationship' предназначеных для удаления
		listIDReporModify     []string // список объектов типа 'report' предназначенных для модификации
		listObjModiy          []*datamodels.ElementSTIXObject
	)
	sl := map[string]struct {
		targetRefsID   string
		relationshipID string
		listRefs       []datamodels.IdentifierTypeSTIX
	}{}

	//получаем все объекты предназначенные для удаления (проверка типа объекта, удаление возможно только объектов типа 'grouping' или
	//'relationship', осуществляется на этапе валидации входных параметров)
	listElementSTIXObject, err := interactionmongodb.FindSTIXObjectByID(qp, listID)
	if err != nil {
		return err
	}

	//обрабатываем все объекты типа 'grouping' и 'relationship' из полученной задачи и отмеченные для удаления
	for _, v := range listElementSTIXObject {
		if v.DataType == "grouping" {
			listIDGroupingDel = append(listIDGroupingDel, v.Data.GetID())

			element, ok := v.Data.(datamodels.GroupingDomainObjectsSTIX)
			if !ok {
				return fmt.Errorf("Error 1: %v", err)
			}

			if len(element.ObjectRefs) == 0 {
				continue
			}

			sl[v.Data.GetID()] = struct {
				targetRefsID   string
				relationshipID string
				listRefs       []datamodels.IdentifierTypeSTIX
			}{listRefs: element.ObjectRefs}
		}

		if v.DataType == "relationship" {
			listIDRelationshipDel = append(listIDRelationshipDel, v.Data.GetID())
		}
	}

	//ищем объекты типа 'relationship' являющиеся связующим звеном между объектами 'grouping' и другими объектами, чаще всего 'report'
	cur, err := qp.Find((bson.D{
		bson.E{Key: "commonpropertiesobjectstix.type", Value: "relationship"},
		bson.E{Key: "source_ref", Value: bson.D{{Key: "$in", Value: listIDGroupingDel}}},
	}))
	if err != nil {
		return fmt.Errorf("Error 2: %v", err)
	}

	//обрабатываем полученый список объектов типа 'relationship'
	for _, v := range interactionmongodb.GetListElementSTIXObject(cur) {
		if obj, ok := v.Data.(datamodels.RelationshipObjectSTIX); ok {
			targetID := string(obj.TargetRef)

			//сохраняем список объектов типа 'relationship' являющихся связующим звеном и которые в последствии необходимо удалить
			listIDRelationshipDel = append(listIDRelationshipDel, obj.ID)
			listIDReporModify = append(listIDReporModify, targetID)

			sl[string(obj.SourceRef)] = struct { // ID объекта типа 'gouping'
				targetRefsID   string                          // объект 'report' на который ссылается какой либо объект из поля SourceRefs
				relationshipID string                          // ID объекта типа 'relationship' который соединяет объекты 'grouping' и 'report' и который так же нужно удалить
				listRefs       []datamodels.IdentifierTypeSTIX // список ID объектов который 'grouping' объединяет в группу и который нужно перенести
				// в объект 'report' к которому принадлежит 'grouping' отмеченный для удаления
			}{
				targetRefsID:   targetID,
				relationshipID: obj.ID,
				listRefs:       sl[string(obj.SourceRef)].listRefs,
			}
		}
	}

	//получаем список ID STIX объектов типа 'report', на которые ссылаются найденные объекты 'relationship'
	cur, err = qp.Find((bson.D{
		bson.E{Key: "commonpropertiesobjectstix.type", Value: "report"},
		bson.E{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: listIDReporModify}}},
	}))
	if err != nil {
		return fmt.Errorf("Error 3: %v", err)
	}

	//обрабатываем полученый список объектов типа 'report' и модифицируем их изменяя свойство ObjectRefs
	for _, v := range interactionmongodb.GetListElementSTIXObject(cur) {
		obj, ok := v.Data.(datamodels.ReportDomainObjectsSTIX)
		if !ok {
			continue
		}

		//ищем ID удаляемого объекта 'grouping' и удаляем из свойства ObjectRefs и в это же свойство добаляем все ссылки которые раньше
		// были в ObjectRefs удаляемого объекта 'grouping'
		for groupingID, v := range sl {
			if v.targetRefsID != obj.ID {
				continue
			}

			//удаляем ID объекта 'grouping' из свойства ObjectRefs объекта 'report' и добавляем туда ссылки на ID объектов находящиеся
			// в свойстве ObjectRefs удаляемого объекта 'grouping'
			listTmp := []datamodels.IdentifierTypeSTIX{}
			for _, v := range obj.ObjectRefs {
				if string(v) == groupingID {
					continue
				}

				listTmp = append(listTmp, v)
			}

			obj.ObjectRefs = append(listTmp, v.listRefs...)
		}

		listObjModiy = append(listObjModiy, &datamodels.ElementSTIXObject{
			DataType: obj.Type,
			Data:     obj,
		})
	}

	//обновляем STIX объекты типа 'report'
	err = interactionmongodb.ReplacementElementsSTIXObject(qp, listObjModiy)
	if err != nil {
		return fmt.Errorf("Error 4: %v", err)
	}

	//удаляем выбранные в списках объекты типа 'relationship' и 'grouping'
	if _, err := qp.DeleteManyData(bson.D{{
		Key:   "commonpropertiesobjectstix.id",
		Value: bson.D{{Key: "$in", Value: append(listIDGroupingDel, listIDRelationshipDel...)}}}}); err != nil {
		return fmt.Errorf("Error 5: %v", err)
	}

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
			Host:     "test-uchet-db.cloud.gcm",
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

	Context("Тест 2. Добавление тестового набора STIX объектов (типы 'report', 'relationship' и 'grouping')", func() {
		It("При добавления тестового набора ошибок быть не должно", func() {
			Expect(errAddListObj).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 3. Выполняем поиск и удаление объектов типа 'grouping' и связующих их объектов типа 'relationship'", func() {
		It("При удалении объектов 'grouping' ошибок быть недолжно", func() {
			Expect(deleteObjTypeGrouping(cdmdb, qp, []string{
				"grouping--911f0f43-1e7c-4f97-7070-a787bdd3b100",
				"grouping--911f0f23-2e7c-4f07-9436-d6e7bdd3b236",
			})).ShouldNot(HaveOccurred())
		})

		It("Должны быть удалены 4 объекта", func() {
			cur, err := qp.FindAllWithLimit(
				interactionmongodb.CreateSearchQueriesSTIXObject(
					&datamodels.SearchThroughCollectionSTIXObjectsType{}),
				&interactionmongodb.FindAllWithLimitOptions{
					Offset:        1,
					LimitMaxSize:  250,
					SortField:     "commonpropertiesdomainobjectstix.created",
					SortAscending: false,
				})
			Expect(err).ShouldNot(HaveOccurred())

			elemSTIXObj := interactionmongodb.GetListElementSTIXObject(cur)
			Expect(len(elemSTIXObj)).Should(Equal(103))
		})
	})
})

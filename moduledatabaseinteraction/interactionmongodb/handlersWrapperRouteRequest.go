package interactionmongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
)

//searchSTIXObject обработчик поисковых запросов, связанных с поиском STIX объектов, по заданным пользователем параметрам
func searchSTIXObject(
	appTaskID string,
	qp QueryParameters,
	taskInfo datamodels.ModAPIRequestProcessingResJSONSearchReqType,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (string, error) {

	var (
		err           error
		fn            string = commonlibs.GetFuncName()
		sortableField string
		sf            = map[string]string{
			"document_type":   "commonpropertiesobjectstix.type",
			"data_created":    "commonpropertiesdomainobjectstix.created",
			"data_modified":   "commonpropertiesdomainobjectstix.modified",
			"data_first_seen": "first_seen",
			"data_last_seen":  "last_seen",
			"ipv4":            "value",
			"ipv6":            "value",
			"country":         "country",
		}
	)

	searchParameters, ok := taskInfo.SearchParameters.(datamodels.SearchThroughCollectionSTIXObjectsType)
	if !ok {
		return fn, fmt.Errorf("type conversion error")
	}

	//получить только общее количество найденных документов
	if (taskInfo.PaginateParameters.CurrentPartNumber <= 0) || (taskInfo.PaginateParameters.MaxPartSize <= 0) {
		resSize, err := qp.CountDocuments(CreateSearchQueriesSTIXObject(&searchParameters))
		if err != nil {
			return fn, err
		}

		//сохраняем общее количество найденных значений во временном хранилище
		err = tst.AddNewFoundInformation(
			appTaskID,
			&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
				Collection:  "stix_object_collection",
				ResultType:  "only_count",
				Information: resSize,
			})
		if err != nil {
			return fn, err
		}

		return fn, nil
	}

	if field, ok := sf[taskInfo.SortableField]; ok {
		sortableField = field
	}

	//получить все найденные документы, с учетом лимита
	cur, err := qp.FindAllWithLimit(CreateSearchQueriesSTIXObject(&searchParameters), &FindAllWithLimitOptions{
		Offset:        int64(taskInfo.PaginateParameters.CurrentPartNumber),
		LimitMaxSize:  int64(taskInfo.PaginateParameters.MaxPartSize),
		SortField:     sortableField,
		SortAscending: false,
	})
	if err != nil {
		return fn, err
	}

	listelm := GetListElementSTIXObject(cur)

	//сохраняем найденные значения во временном хранилище
	err = tst.AddNewFoundInformation(
		appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection:  "stix_object_collection",
			ResultType:  "full_found_info",
			Information: listelm,
		})
	if err != nil {
		return fn, err
	}

	return fn, nil
}

func searchDifferencesObjectsCollection(
	appTaskID string,
	qp QueryParameters,
	taskInfo datamodels.ModAPIRequestProcessingResJSONSearchReqType,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (string, error) {

	var (
		err                  error
		searchDocumentId     bson.E
		searchCollectionName bson.E
		fn                   string = commonlibs.GetFuncName()
		currentPartNumber    int64  = 1
		maxPartSize          int64  = 350
		sortableField        string = "modified_time"
	)

	fmt.Printf("func '%s', START...\n", fn)

	if taskInfo.PaginateParameters.CurrentPartNumber != 0 {
		currentPartNumber = int64(taskInfo.PaginateParameters.CurrentPartNumber)
	}

	if taskInfo.PaginateParameters.MaxPartSize > 15 {
		maxPartSize = int64(taskInfo.PaginateParameters.MaxPartSize)
	}

	if taskInfo.SortableField == "user_name_modified_object" {
		sortableField = taskInfo.SortableField
	}

	sp, ok := taskInfo.SearchParameters.(struct {
		DocumentID     string `json:"document_id"`
		CollectionName string `json:"collection_name"`
	})
	if !ok {
		return fn, fmt.Errorf("type conversion error")
	}

	if len(sp.DocumentID) > 0 {
		searchDocumentId = bson.E{Key: "document_id", Value: sp.DocumentID}
	}

	if len(sp.CollectionName) > 0 {
		searchCollectionName = bson.E{Key: "collection_name", Value: sp.CollectionName}
	}

	//получить все найденные документы, с учетом лимита
	cur, err := qp.FindAllWithLimit(bson.D{
		searchDocumentId,
		searchCollectionName,
	}, &FindAllWithLimitOptions{
		Offset:        currentPartNumber,
		LimitMaxSize:  maxPartSize,
		SortField:     sortableField,
		SortAscending: false,
	})
	if err != nil {
		return fn, err
	}

	documents := []datamodels.DifferentObjectType{}

	for cur.Next(context.Background()) {
		var document datamodels.DifferentObjectType
		if err := cur.Decode(&document); err != nil {
			continue
		}

		documents = append(documents, document)
	}

	fmt.Printf("func '%s', count document found: '%d'\nfound result ==== \n'%v'\n", fn, len(documents), documents)

	//сохраняем найденные значения во временном хранилище
	err = tst.AddNewFoundInformation(
		appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection:  "accounting_differences_objects_collection",
			ResultType:  "full_found_info",
			Information: documents,
		})
	if err != nil {
		return fn, err
	}

	return fn, err
}

//ResultStatisticalInformationSTIXObject содержит результат поиска статистической информации по STIX объектам
// InformationType - тип статистической информации
// ListComputerThreat - список статистической информации по компьютерным угрозам
type ResultStatisticalInformationSTIXObject struct {
	InformationType    string           `json:"information_type"`
	ListComputerThreat map[string]int32 `json:"list_computer_threat"`
}

//statisticalInformationSTIXObject обработчик поиска статистической информации о STIX объектах
func statisticalInformationSTIXObject(
	parameters struct {
		appTaskID                  string
		qp                         QueryParameters
		tst                        *memorytemporarystoragecommoninformation.TemporaryStorageType
		TypeStatisticalInformation string
	}) (string, error) {

	var (
		err                       error
		fn                        string = commonlibs.GetFuncName()
		tmpResults                []bson.M
		outsideSpecificationField = "decisions_made_computer_threat"
		rsiSTIXObject             = ResultStatisticalInformationSTIXObject{
			InformationType:    "decisions_made_computer_threat",
			ListComputerThreat: map[string]int32{},
		}
	)

	if parameters.TypeStatisticalInformation == "types computer threat" {
		outsideSpecificationField = "computer_threat_type"
		rsiSTIXObject.InformationType = "computer_threat_type"
	}

	opts := options.Aggregate().SetAllowDiskUse(true)
	collection := parameters.qp.ConnectDB.Database(parameters.qp.NameDB).Collection(parameters.qp.CollectionName)
	cur, err := collection.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			bson.D{bson.E{Key: "$match", Value: bson.D{
				bson.E{Key: "commonpropertiesobjectstix.type", Value: "report"},
			}}},
			bson.D{
				bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: fmt.Sprintf("$outside_specification.%s", outsideSpecificationField)},
					bson.E{Key: "count", Value: bson.D{
						bson.E{Key: "$sum", Value: 1},
					}},
				}}}},
		opts)
	if err != nil {
		return fn, err
	}

	err = cur.All(context.TODO(), &tmpResults)
	if err != nil {
		return fn, err
	}

	for _, v := range tmpResults {
		name, ok := v["_id"].(string)
		if !ok {
			continue
		}

		if count, ok := v["count"].(int32); ok {
			rsiSTIXObject.ListComputerThreat[name] = count
		}
	}

	//сохраняем найденные значения во временном хранилище
	err = parameters.tst.AddNewFoundInformation(
		parameters.appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection:  "stix_object_collection",
			ResultType:  "handling_statistical_requests",
			Information: rsiSTIXObject,
		})
	if err != nil {
		return fn, err
	}

	return fn, nil
}

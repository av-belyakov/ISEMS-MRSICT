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

//searchSTIXObject обработчик поисковых запросов, связанных с поиском, по заданным параметрам, STIX объектов
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

	fmt.Printf("____ func 'searchSTIXObject' searchParameters: '%v'\n", searchParameters)

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
	fmt.Printf("____ func 'searchSTIXObject' appTaskID: '%s', count ListElementSTIXObject: '%d'\n", appTaskID, len(listelm))
	/*for k, v := range listelm {
		fmt.Printf("%d. %s\n", k, v.Data.GetID())
	}*/

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

	fmt.Printf("func 'statisticalInformationSTIXObject', \nlist 'decisions_made_computer_threat': %v\n", tmpResults)

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

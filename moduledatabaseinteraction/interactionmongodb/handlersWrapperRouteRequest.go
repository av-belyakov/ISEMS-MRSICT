package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"

	"go.mongodb.org/mongo-driver/bson"
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
	if (taskInfo.PaginateParameters.CurrentPartNumber <= 0) || (taskInfo.PaginateParameters.MaxPartNum <= 0) {
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
		LimitMaxSize:  int64(taskInfo.PaginateParameters.MaxPartNum),
		SortField:     sortableField,
		SortAscending: false,
	})
	if err != nil {
		return fn, err
	}

	//сохраняем найденные значения во временном хранилище
	err = tst.AddNewFoundInformation(
		appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection:  "stix_object_collection",
			ResultType:  "full_found_info",
			Information: GetListElementSTIXObject(cur),
		})
	if err != nil {
		return fn, err
	}

	return fn, nil
}

//searchListComputerThreat обработчик запроса, на получения списка id объектов типа 'grouping', с их описанием, относящихся к типам "types decisions
// made computer threat" или "types computer threat"
func searchListComputerThreat(appTaskID string,
	qp QueryParameters,
	taskInfo datamodels.ModAPIRequestProcessingResJSONSearchReqType,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (string, error) {

	var (
		err error
		fn  string = commonlibs.GetFuncName()
		l          = map[string]datamodels.StorageApplicationCommonListType{}
	)

	searchType, ok := taskInfo.SearchParameters.(struct {
		TypeList string `json:"type_list"`
	})
	if !ok {
		return fn, fmt.Errorf("type conversion error")
	}

	switch searchType.TypeList {
	case "types decisions made computer threat":
		l, err = tst.GetListDecisionsMade()

	case "types computer threat":
		l, err = tst.GetListComputerThreat()

	default:
		return fn, fmt.Errorf("undefined type of computer threat list")

	}

	if err != nil {
		return fn, err
	}

	//сохраняем найденные значения во временном хранилище
	if err = tst.AddNewFoundInformation(
		appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection: "stix_object_collection",
			ResultType: "list_computer_threat",
			Information: struct {
				TypeList string                                                 `json:"type_list"`
				List     map[string]datamodels.StorageApplicationCommonListType `json:"list"`
			}{
				TypeList: searchType.TypeList,
				List:     l,
			},
		}); err != nil {
		return fn, err
	}

	return fn, nil
}

//searchSTIXObjectListTypeGrouping обработчик поисковых запросов, связанных с поиском предустановленного набора STIX объектов типа 'Grouping',
// относящихся к спискам 'типы принимаемых решений по компьютерным угрозам' и 'типы компьютерных угроз'
func searchSTIXObjectListTypeGrouping(
	appTaskID string,
	qp QueryParameters,
	taskInfo datamodels.ModAPIRequestProcessingResJSONSearchReqType,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (string, error) {

	var (
		err error
		fn  string = commonlibs.GetFuncName()
	)

	fmt.Println("func 'searchSTIXObjectListTypeGrouping', START...")

	searchType, ok := taskInfo.SearchParameters.(struct {
		TypeList string `json:"type_list"`
	})
	if !ok {
		return fn, fmt.Errorf("type conversion error")
	}

	l := map[string]datamodels.StorageApplicationCommonListType{}

	switch searchType.TypeList {
	case "types decisions made computer threat":

		fmt.Println("func 'searchSTIXObjectListTypeGrouping', types decisions made computer threat")

		l, err = tst.GetListDecisionsMade()

	case "types computer threat":

		fmt.Println("func 'searchSTIXObjectListTypeGrouping', types computer threat")

		l, err = tst.GetListComputerThreat()

	default:
		err = fmt.Errorf("undefined type of computer threat list")

	}

	if err != nil {
		return fn, err
	}

	ls := make([]string, 0, len(l))
	for k := range l {
		ls = append(ls, k)
	}

	//получить все найденные документы, с учетом лимита
	cur, err := qp.FindAllWithLimit(bson.D{{Key: "name", Value: bson.D{{Key: "$in", Value: ls}}}}, &FindAllWithLimitOptions{
		Offset:        1,
		LimitMaxSize:  1000,
		SortField:     taskInfo.SortableField,
		SortAscending: false,
	})
	if err != nil {
		return fn, err
	}

	listComputerThreat := GetListGroupingComputerThreat(cur)

	fmt.Printf("func 'searchSTIXObjectListTypeGrouping', получить все найденные документы, с учетом лимита: '%v'\n", listComputerThreat)

	//сохраняем найденные значения во временном хранилище
	err = tst.AddNewFoundInformation(
		appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection:  "stix_object_collection",
			ResultType:  "found_info_list_computer_threat",
			Information: listComputerThreat,
		})
	if err != nil {
		return fn, err
	}

	return fn, nil
}

package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"

	"go.mongodb.org/mongo-driver/bson"
)

//searchSTIXObjectListTypeGrouping обработчик поисковых запросов связанных с поиском предустановленного набора STIX объектов типа 'Grouping',
// относящихся к спискам 'типы принимаемых решений по компьютерным угрозам' и 'типы компьютерных угроз'
func searchSTIXObjectListTypeGrouping(
	appTaskID string,
	qp QueryParameters,
	taskInfo datamodels.ModAPIRequestProcessingResJSONSearchReqType,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (string, error) {

	var err error
	fn := "searchSTIXObjectListTypeGrouping"

	searchType, ok := taskInfo.SearchParameters.(struct {
		TypeList string `json:"type_list"`
	})
	if !ok {
		return fn, fmt.Errorf("type conversion error")
	}

	l := map[string]datamodels.StorageApplicationCommonListType{}

	switch searchType.TypeList {
	case "types decisions made computer threat":
		l, err = tst.GetListDecisionsMade()

	case "types computer threat":
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

	//сохраняем найденные значения во временном хранилище
	err = tst.AddNewFoundInformation(
		appTaskID,
		&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
			Collection:  "stix_object_collection",
			ResultType:  "found_info_list_computer_threat",
			Information: GetListGroupingObjectSTIX(cur),
		})
	if err != nil {
		return fn, err
	}

	return fn, nil
}

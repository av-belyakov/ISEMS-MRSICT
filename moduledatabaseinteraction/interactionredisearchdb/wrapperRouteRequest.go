package interactionredisearchdb

import (
	"fmt"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"

	"github.com/RediSearch/redisearch-go/redisearch"
)

var errorMessage = datamodels.ModuleDataBaseInteractionChannel{
	CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
		ModuleGeneratorMessage: "module database interaction",
		ModuleReceiverMessage:  "module core application",
		ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
			ModuleAPIRequestProcessingSettingSendTo: true,
		},
	},
}

func getIndex(elem datamodels.ElementSTIXObject) datamodels.RedisearchIndexObject {
	mt := elem.Data.GeneratingDataForIndexing()
	indexObject := datamodels.RedisearchIndexObject{
		Type: elem.DataType,
	}

	for k, v := range mt {
		if k == "id" {
			indexObject.ID = v
		}

		if k == "name" {
			indexObject.Name = v
		}

		if k == "description" {
			indexObject.Description = v
		}

		if k == "street_address" {
			indexObject.StreetAddress = v
		}

		if k == "abstract" {
			indexObject.Abstract = v
		}

		if k == "aliases" {
			indexObject.Aliases = v
		}

		if k == "content" {
			indexObject.Content = v
		}

		if k == "value" {
			indexObject.Value = v
		}
	}

	return indexObject
}

func getListIndex(listElem []*datamodels.ElementSTIXObject) []datamodels.RedisearchIndexObject {
	listIndexObj := []datamodels.RedisearchIndexObject{}

	for _, v := range listElem {
		listIndexObj = append(listIndexObj, getIndex(*v))
	}

	return listIndexObj
}

func getRedisearchDocument(listIndex []datamodels.RedisearchIndexObject) []redisearch.Document {
	redisearchDoc := make([]redisearch.Document, 0, len(listIndex))

	for _, v := range listIndex {
		tmp := redisearch.NewDocument(v.ID, 1.0)
		tmp.Set("type", v.Type)
		tmp.Set("name", v.Name)
		tmp.Set("description", v.Description)
		tmp.Set("street_address", v.StreetAddress)
		tmp.Set("abstract", v.Abstract)
		tmp.Set("content", v.Content)
		tmp.Set("value", v.Value)

		redisearchDoc = append(redisearchDoc, tmp)
	}

	return redisearchDoc
}

func wrapperFuncHandlingInsertIndex(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	dataRequest datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	cdrdb ConnectionDescriptorRedisearchDB) {

	var (
		err error
		fn  = commonlibs.GetFuncName()
	)

	errorMessage.ErrorMessage.FuncName = fn
	errorMessage.Section = "handling insert index"
	errorMessage.AppTaskID = dataRequest.AppTaskID

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(dataRequest.AppTaskID)
	if err != nil {
		errorMessage.ErrorMessage.Error = fmt.Errorf("no information about the task by its id was found in the temporary storage")
		chanOutput <- errorMessage

		return
	}

	listElementSTIX, ok := taskInfo.TaskParameters.([]*datamodels.ElementSTIXObject)
	if !ok {
		errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
		chanOutput <- errorMessage

		return
	}

	listIndex := getListIndex(listElementSTIX)
	if err := cdrdb.Connection.IndexOptions(
		redisearch.IndexingOptions{
			Replace: true,
			Partial: true,
		}, getRedisearchDocument(listIndex)...); err != nil {
		errorMessage.ErrorMessage.Error = err
		chanOutput <- errorMessage

		return
	}
}

func wrapperFuncHandlingSelectIndex(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	cdrdb ConnectionDescriptorRedisearchDB) {

	fmt.Println("func 'wrapperFuncHandlingSelectIndex', START")
}

func wrapperAutoCompleteSuggestions(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	cdrdb ConnectionDescriptorRedisearchDB) {

	fmt.Println("func 'wrapperAutoCompleteSuggestions', START")
}

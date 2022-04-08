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

func wrapperFuncHandlingInsertIndex(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	dataRequest datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	cdrdb ConnectionDescriptorRedisearchDB) {

	var (
		err error
		fn  = commonlibs.GetFuncName()
	)

	fmt.Printf("func '%v', START...\n", fn)

	errorMessage.ErrorMessage.FuncName = fn
	errorMessage.Section = "handling insert index"
	errorMessage.AppTaskID = dataRequest.AppTaskID

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(dataRequest.AppTaskID)
	if err != nil {
		errorMessage.ErrorMessage.Error = fmt.Errorf("no information about the task by its id was found in the temporary storage")
		chanOutput <- errorMessage

		fmt.Printf("func '%v', ERROR: '%v'\n", fn, errorMessage.ErrorMessage.Error)

		return
	}

	fmt.Printf("func '%v', GET TaskInfo '%v'\n", fn, taskInfo)

	listElementSTIX, ok := taskInfo.TaskParameters.([]*datamodels.ElementSTIXObject)
	if !ok {
		errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
		chanOutput <- errorMessage

		fmt.Printf("func '%v', ERROR: '%v'\n", fn, errorMessage.ErrorMessage.Error)

		return
	}

	fmt.Printf("func '%v', GET listElementSTIX, COUNT = '%d'\n", fn, len(listElementSTIX))

	var newDocumentList = make([]redisearch.Document, 0, len(listElementSTIX))
	for _, v := range listElementSTIX {
		if v.DataType == "relationship" || v.DataType == "sighting" {
			continue
		}

		vdata := v.Data.GeneratingDataForIndexing()
		tmp := redisearch.NewDocument(vdata["id"], 1.0)

		for key, value := range vdata {
			if key == "id" {
				continue
			}

			tmp.Set(key, value)
		}

		newDocumentList = append(newDocumentList, tmp)
	}

	fmt.Printf("func '%v', GET newDocumentList, COUNT = '%d'\n", fn, len(newDocumentList))

	if err := cdrdb.Connection.IndexOptions(
		redisearch.IndexingOptions{
			Replace: true,
			Partial: true,
		}, newDocumentList...); err != nil {
		errorMessage.ErrorMessage.Error = err
		chanOutput <- errorMessage

		fmt.Printf("func '%v', ERROR: '%v'\n", fn, errorMessage.ErrorMessage.Error)

		return
	}

	fmt.Printf("func '%v', END\n", fn)
}

func wrapperFuncHandlingSelectIndex(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	cdrdb ConnectionDescriptorRedisearchDB) {

	fmt.Println("func 'wrapperFuncHandlingSelectIndex', START")
}

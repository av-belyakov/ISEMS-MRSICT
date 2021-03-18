package moduledatabaseinteraction

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//ChannelsModuleDataBaseInteraction описание каналов передачи данных между ядром приложения и модулем взаимодействия с базами данных
type ChannelsModuleDataBaseInteraction struct {
	ChannelsMongoDB interactionmongodb.ChannelsMongoDBInteraction
}

var cmdbi ChannelsModuleDataBaseInteraction

func init() {
	cmdbi = ChannelsModuleDataBaseInteraction{}
}

//MainHandlerDataBaseInteraction модуль инициализации обработчиков для взаимодействия с базами данных
func MainHandlerDataBaseInteraction(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	cdb *datamodels.ConnectionsDataBase,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (ChannelsModuleDataBaseInteraction, error) {

	fmt.Println("func 'MainHandlerDataBaseInteraction', START...")

	//инициализируем модуль для взаимодействия с БД MongoDB
	chanMongoDB, err := interactionmongodb.InteractionMongoDB(chanSaveLog, &cdb.MongoDBSettings, tst)
	if err != nil {
		return cmdbi, err
	}
	cmdbi.ChannelsMongoDB = chanMongoDB

	return cmdbi, nil
}

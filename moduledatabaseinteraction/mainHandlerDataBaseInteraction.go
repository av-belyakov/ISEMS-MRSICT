package moduledatabaseinteraction

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
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
func MainHandlerDataBaseInteraction(cdb *datamodels.ConnectionsDataBase) (ChannelsModuleDataBaseInteraction, error) {
	fmt.Println("func 'MainHandlerDataBaseInteraction', START...")

	//инициализируем модуль для взаимодействия с БД MongoDB
	chanMongoDB, err := interactionmongodb.InteractionMongoDB(&cdb.MongoDBSettings)
	if err != nil {
		return cmdbi, err
	}
	cmdbi.ChannelsMongoDB = chanMongoDB

	return cmdbi, nil
}

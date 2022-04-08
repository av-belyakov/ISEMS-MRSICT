package moduledatabaseinteraction

import (
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionredisearchdb"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//ChannelsModuleDataBaseInteraction описание каналов передачи данных между ядром приложения и модулем взаимодействия с базами данных
type ChannelsModuleDataBaseInteraction struct {
	ChannelsMongoDB      interactionmongodb.ChannelsMongoDBInteraction
	ChannelsRedisearchDB interactionredisearchdb.ChannelsRedisearchInteraction
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

	//инициализируем модуль для взаимодействия с БД MongoDB
	chanMongoDB, err := interactionmongodb.InteractionMongoDB(chanSaveLog, &cdb.MongoDBSettings, tst)
	if err != nil {
		return cmdbi, err
	}
	cmdbi.ChannelsMongoDB = chanMongoDB

	//инициализируем модуль для взаимодействия с БД Redisearch
	chanRedisearchDB, err := interactionredisearchdb.InteractionRedisearchDB(chanSaveLog, &cdb.RedisearchDBSettings, tst)
	if err != nil {
		return cmdbi, err
	}
	cmdbi.ChannelsRedisearchDB = chanRedisearchDB

	return cmdbi, nil
}

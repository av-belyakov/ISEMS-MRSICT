package moduledatabaseinteraction

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
)

//ChansDataBaseInteraction описание каналов для взаимодействия с базами данных
type ChansDataBaseInteraction struct {
	ChansMongoDB interactionmongodb.ChansMongoDBInteraction
}

//MainHandlerDataBaseInteraction модуль инициализации обработчиков для взаимодействия с базами данных
func MainHandlerDataBaseInteraction(cdb *datamodels.ConnectionsDataBase) (ChansDataBaseInteraction, error) {
	fmt.Println("func 'MainHandlerDataBaseInteraction', START...")

	cdbi := ChansDataBaseInteraction{}

	chanMongoDB, err := interactionmongodb.InteractionMongoDB(&cdb.MongoDBSettings)
	if err != nil {
		return cdbi, err
	}

	cdbi.ChansMongoDB = chanMongoDB

	return cdbi, nil
}

package interactionredisearchdb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
)

/*type wrappersSetting struct {
	AdditionalRequestParameters interface{}
	NameDB                      string
	ConnectionDB                ConnectionDescriptorMongoDB
}*/

//Routing обеспечивает маршрутизацию информации между базой данных RedisearchDB и ядром приложения
func Routing(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	cdrdb ConnectionDescriptorRedisearchDB,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	chanInput <-chan datamodels.ModuleDataBaseInteractionChannel) {

	fmt.Printf("func 'Routing' Redisearch DB")

	/*
		ws := wrappersSetting{
			NameDB:       nameDb,
			ConnectionDB: cdmdb,
		}

		for data := range chanInput {
			switch data.Section {
			case "handling stix object":
				go ws.wrapperFuncTypeHandlingSTIXObject(chanOutput, data, tst)

			case "handling managing collection stix objects":
				go ws.wrapperFuncTypeHandlingManagingCollectionSTIXObjects(chanOutput, data, tst)

			case "handling managing differences objects collection":
				go ws.wrapperFuncTypeHandlingManagingDifferencesObjectsCollection(chanOutput, data, tst)

			case "handling search requests":
				go ws.wrapperFuncTypeHandlingSearchRequests(chanOutput, data, tst)

			case "handling reference book":
				go ws.wrapperFuncTypeHandlingReferenceBook(chanOutput, data, tst)

			case "handling technical part":
				go ws.wrapperFuncTypeTechnicalPart(chanOutput, data, tst)

			case "handling statistical requests":
				go ws.wrapperFuncTypeHandlingStatisticalRequests(chanOutput, data, tst)

			}
		}
	*/
}

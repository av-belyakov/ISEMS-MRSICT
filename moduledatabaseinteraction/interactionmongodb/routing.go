package interactionmongodb

import (
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
)

type wrappersSetting struct {
	AdditionalRequestParameters interface{}
	NameDB                      string
	ConnectionDB                ConnectionDescriptorMongoDB
}

//Routing обеспечивает маршрутизацию информации между базой данных MongoDB и ядром приложения
func Routing(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	nameDb string,
	cdmdb ConnectionDescriptorMongoDB,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	chanInput <-chan datamodels.ModuleDataBaseInteractionChannel) {

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
}

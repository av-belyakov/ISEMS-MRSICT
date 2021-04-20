package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
)

type wrappersSetting struct {
	DataRequest                 datamodels.ModuleDataBaseInteractionChannel
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

	fmt.Println("ModuleDataBaseInteraction - func 'Routing', START...")

	ws := wrappersSetting{
		NameDB:       nameDb,
		ConnectionDB: cdmdb,
	}

	for data := range chanInput {
		fmt.Printf("func 'Routing', received data from chan: '%v'\n", data)

		ws.DataRequest = data

		switch data.Section {
		case "handling stix object":
			go ws.wrapperFuncTypeHandlingSTIXObject(chanOutput, tst)

		case "handling search requests":
			go ws.wrapperFuncTypeHandlingSearchRequests(tst)

		case "handling reference book":
			go ws.wrapperFuncTypeHandlingReferenceBook(tst)

		}
	}
}

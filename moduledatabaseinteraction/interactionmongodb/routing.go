package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

type wrappersSetting struct {
	DataRequest                 datamodels.ModuleDataBaseInteractionChannel
	AdditionalRequestParameters interface{}
	NameDB                      string
	ChanSaveLog                 chan<- modulelogginginformationerrors.LogMessageType
	ConnectionDB                ConnectionDescriptorMongoDB
}

//Routing обеспечивает маршрутизацию информации между базой данных MongoDB и ядром приложения
func Routing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	nameDb string,
	cdmdb ConnectionDescriptorMongoDB,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	chanInput <-chan datamodels.ModuleDataBaseInteractionChannel) {

	fmt.Println("ModuleDataBaseInteraction - func 'Routing', START...")

	ws := wrappersSetting{
		NameDB:       nameDb,
		ChanSaveLog:  chanSaveLog,
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

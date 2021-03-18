package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

type wrappersSetting struct {
	Command                     string
	AppTaskID                   string
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

		ws.Command = data.Command
		ws.AppTaskID = data.AppTaskID

		switch data.Section {
		case "handling stix object":
			go ws.wrapperFuncTypeHandlingSTIXObject()

		case "handling search requests":
			go ws.wrapperFuncTypeHandlingSearchRequests()

		case "handling reference book":
			go ws.wrapperFuncTypeHandlingReferenceBook()

		}
	}
}

package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
)

//Routing обеспечивает маршрутизацию информации между базой данных MongoDB и ядром приложения
func Routing(cdmdb ConnectionDescriptorMongoDB, chanInput <-chan datamodels.ModuleDataBaseInteractionChannel) {
	fmt.Println("ModuleDataBaseInteraction - func 'Routing', START...")

	for data := range chanInput {
		fmt.Printf("func 'Routing', received data from chan: '%v'\n", data)

		switch data.Section {
		case "handling stix object":

		case "handling search requests":

		case "handling reference book":

		}
	}
}

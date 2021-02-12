package modulecoreapplication

import (
	"fmt"
	"log"

	"ISEMS-MRSICT/datamodels"
)

//RoutingCoreApp обеспечивает маршрутизацию всех данных циркулирующих внутри приложения
func RoutingCoreApp(appConfig *datamodels.AppConfig, clim *ChannelsListInteractingModules) {
	fmt.Println("ModuleCoreApplication - func 'RoutingCoreApp', START...")

	log.Printf("Start application ISEMS-MRSICT, version '%q'\n", appConfig.VersionApp)

	for {
		select {
		case data := <-clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.OutputModule:
			fmt.Printf("func 'Routing', Input data from 'moduleDataBaseInteraction', data base MongoDB. Reseived data: '%v'\n", data)
		case data := <-clim.ChannelsModuleAPIRequestProcessing.OutputModule:
			fmt.Printf("func 'Routing', Input data from 'moduleAPIRequestProcessing'. Reseived data: '%v'\n", data)
		}
	}
}

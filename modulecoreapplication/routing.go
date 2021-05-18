package modulecoreapplication

import (
	"fmt"
	"log"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/routingflowsmoduleapirequestprocessing"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//RoutingCoreApp обеспечивает маршрутизацию всех данных циркулирующих внутри приложения
func RoutingCoreApp(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	appConfig *datamodels.AppConfig,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clim *moddatamodels.ChannelsListInteractingModules) {

	log.Printf("Start application ISEMS-MRSICT, version '%q'\n", appConfig.VersionApp)

	for {
		select {
		case data := <-clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.OutputModule:

			fmt.Printf("func 'Routing', Input data from 'moduleDataBaseInteraction', data base MongoDB. Reseived data: '%v'\n", data)

			go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleDataBaseInteraction(chanSaveLog, &data, tst, clim)

		case data := <-clim.ChannelsModuleAPIRequestProcessing.OutputModule:

			fmt.Printf("func 'Routing', Input data from 'moduleAPIRequestProcessing'. Reseived data: '%v'\n", data)

			go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, &data, tst, clim)
		}
	}
}

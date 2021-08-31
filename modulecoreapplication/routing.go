package modulecoreapplication

import (
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
			// Обработка запросов к хранилищу на MongoDB хранящему SDO,SOO-объектов, справочников, истиории изменений объектов
			go routingflowsmoduleapirequestprocessing.HandlerAssignmentsModuleDataBaseInteraction(chanSaveLog, data, tst, clim)

		case data := <-clim.ChannelsModuleAPIRequestProcessing.OutputModule:
			// Обработка JSON сообщений приходящих от API
			go routingflowsmoduleapirequestprocessing.HandlerAssignmentsModuleAPIRequestProcessing(chanSaveLog, data, tst, clim)
		}
	}
}

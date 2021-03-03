package modulecoreapplication

import (
	"fmt"
	"log"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/modulecoreapplication/requesthandlers"
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

			commonMsgReq, err := requesthandlers.UnmarshalJSONCommonReq(data.Data)
			if err != nil {
				//здесь отправляем информационное сообщение клиенту API

				continue
			}

			switch commonMsgReq.Section {
			case "handling stix object":
				err := requesthandlers.UnmarshalJSONObjectSTIXReq(*commonMsgReq)
				if err != nil {
					//здесь отправляем информационное сообщение клиенту API

					continue
				}

			case "handling search requests":

			case "":

			}
		}
	}
}

package modulecoreapplication

import (
	"fmt"
	"log"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduleapirequestprocessing/auxiliaryfunctions"
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

			//обработка ошибок получаемых от БД MongoDB
			if data.ErrorMessage.Error != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					FuncName:    data.ErrorMessage.FuncName,
					Description: fmt.Sprint(data.ErrorMessage.Error),
				}

				ClientID := ""

				/*
										!!!!!!!!!!!!!!
					Тут проверяем наличие ID задачи в data.AppTaskID, если ID существует, то,
					делаем запрос к временному хранилищу на получение по AppTaskID ID клиета (ClientID) модуля
					moduleapirequestprocessing

				*/

				if err := auxiliaryfunctions.SendCriticalErrorMessageJSON(&auxiliaryfunctions.ErrorMessageType{
					ClientID: ClientID,
					Error:    data.ErrorMessage.Error,
					C:        clim.ChannelsModuleAPIRequestProcessing.InputModule,
				}); err != nil {
					//запись информации в лог-файл
					chanSaveLog <- modulelogginginformationerrors.LogMessageType{
						TypeMessage: "error",
						Description: fmt.Sprint(err),
						FuncName:    "unmarshalJSONCommonReq",
					}

					return
				}
			}

			//обработка информационных сообщений получаемых от БД MongoDB
			if data.InformationMessage != "" {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "info",
					Description: data.InformationMessage,
				}
			}

		case data := <-clim.ChannelsModuleAPIRequestProcessing.OutputModule:
			fmt.Printf("func 'Routing', Input data from 'moduleAPIRequestProcessing'. Reseived data: '%v'\n", data)

			go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, &data, tst, clim)

		}
	}
}

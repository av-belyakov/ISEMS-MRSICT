package modulecoreapplication

import (
	"fmt"
	"log"

	"ISEMS-MRSICT/commonlibs"
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

	var (
		err                       error
		ti                        *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType
		section, taskType, taskID string
	)

	for {
		select {
		case data := <-clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.OutputModule:

			fmt.Printf("func 'Routing', Input data from 'moduleDataBaseInteraction', data base MongoDB. Reseived data: '%v'\n", data)

			if data.Section == "handling stix object" {
				section = "обработка структурированных данных"
				taskType = "добавление или обновление структурированных данных"
			}

			//получаем всю информацию о задаче
			taskID, ti, err = tst.GetTaskByID(data.AppTaskID)
			if err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					FuncName:    data.ErrorMessage.FuncName,
					Description: fmt.Sprint(data.ErrorMessage.Error),
				}
			}

			//обработка ошибок получаемых от БД MongoDB
			if data.ErrorMessage.Error != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					FuncName:    data.ErrorMessage.FuncName,
					Description: fmt.Sprint(data.ErrorMessage.Error),
				}

				//если сообщение об ошибки не надо отправлять клиенту модуля 'moduleapirequestprocessing'
				if !data.ErrorMessage.ModuleAPIRequestProcessingSettingSendTo {
					return
				}

				if taskID == "" {
					return
				}

				if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
					ClientID:         ti.ClientID,
					TaskID:           taskID,
					Section:          data.Section,
					TypeNotification: "danger",
					Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
						Section:     section,
						TaskType:    taskType,
						FinalResult: "задача отклонена",
						Message:     "при выполнении задачи возникла ошибка базы данных",
					}),
					C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
				}); err != nil {
					chanSaveLog <- modulelogginginformationerrors.LogMessageType{
						TypeMessage: "error",
						FuncName:    data.ErrorMessage.FuncName,
						Description: fmt.Sprint(data.ErrorMessage.Error),
					}
				}

				return
			}

			//обработка информационных сообщений получаемых от БД MongoDB
			if data.InformationMessage.Type != "" {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "info",
					Description: data.InformationMessage.Message,
				}

				if taskID == "" {
					return
				}

				if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
					ClientID:         ti.ClientID,
					TaskID:           taskID,
					Section:          data.Section,
					TypeNotification: data.InformationMessage.Type,
					Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
						Section:     section,
						TaskType:    taskType,
						FinalResult: "задача выполнена",
						Message:     data.InformationMessage.Message,
					}),
					C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
				}); err != nil {
					chanSaveLog <- modulelogginginformationerrors.LogMessageType{
						TypeMessage: "error",
						FuncName:    data.ErrorMessage.FuncName,
						Description: fmt.Sprint(data.ErrorMessage.Error),
					}
				}

				return
			}

			/*
					Здесь обработка ответов от БД не относящихся к ошибкам или информационным сообщениям

				//определяем тип задачи по команде, если в результате обработки запроса к БД должны быть получены какие либо данные

				//делаем запрос к временному хранилищу для получении информации о задаче и результатов ее обработки
				//по секции задачи и команде определить выполняемые в результате нее действия
				//отправлять результаты обработки единым целом или кусочками

			*/

		case data := <-clim.ChannelsModuleAPIRequestProcessing.OutputModule:

			fmt.Printf("func 'Routing', Input data from 'moduleAPIRequestProcessing'. Reseived data: '%v'\n", data)

			go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, &data, tst, clim)
		}
	}
}

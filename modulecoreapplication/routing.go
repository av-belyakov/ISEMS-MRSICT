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
		section string
	)

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

					одна из ошибок получаемая из модуля БД это ошибка он не найденной по указанной ID задаче
					Пример:
						chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
							CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
							ModuleGeneratorMessage: "module database interaction",
							ModuleReceiverMessage:  "module core application",
							ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
								FuncName:                                fn,
								ModuleAPIRequestProcessingSettingSendTo: true,
								Error:                                   fmt.Errorf("no information about the task by its id was found in the temporary storage"),
								},
							},
							Section:   "handling stix object",
							AppTaskID: data.AppTaskID,
						}

				*/

				if data.Section == "handling stix object" {
					section = "обработка структурированных данных"
				}

				if err := auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
					ClientID: ClientID,
					Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
						Section:     section,
						TaskType:    "изменение статуса задачи на 'завершена'",
						FinalResult: "задача успешно выполнена",
						Message:     data.InformationMessage.Message,
					}),
					C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
				}); err != nil {
					//запись информации в лог-файл
					chanSaveLog <- modulelogginginformationerrors.LogMessageType{
						TypeMessage: "error",
						Description: fmt.Sprint(err),
						FuncName:    "unmarshalJSONCommonReq",
					}

					return
				}

				/*if err := auxiliaryfunctions.SendCriticalErrorMessageJSON(&auxiliaryfunctions.ErrorMessageType{
					ClientID: ClientID,
					Error:    data.ErrorMessage.Error,
					C:        clim.ChannelsModuleAPIRequestProcessing.InputModule,
				})*/
			}

			//обработка информационных сообщений получаемых от БД MongoDB
			if data.InformationMessage.Type != "" {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "info",
					Description: data.InformationMessage.Message,
				}

				/*
														!!!!!!
										Здесь нужно отправить информационное сообщение клиенту API
														!!!!!!

					common.PatternUserMessage(&common.TypePatternUserMessage{
													TaskType:   "изменение статуса задачи на 'завершена'",
													TaskAction: "задача отклонена",
													Message:    "внутренняя ошибка приложения",
												})

														notifications.SendNotificationToClientAPI(
											outCoreChans.OutCoreChanAPI,
											notifications.NotificationSettingsToClientAPI{
												MsgType: "danger",
												MsgDescription: common.PatternUserMessage(&common.TypePatternUserMessage{
													TaskType:   "изменение статуса задачи на 'завершена'",
													TaskAction: "задача отклонена",
													Message:    "внутренняя ошибка приложения",
												}),
											},
											res.TaskIDClientAPI,
											res.IDClientAPI)
				*/
			}

			//определяем тип задачи по команде, если в результате обработки запроса к БД должны быть получены какие либо данные

			//делаем запрос к временному хранилищу для получении информации о задаче и результатов ее обработки
			//по секции задачи и команде определить выполняемые в результате нее действия
			//отправлять результаты обработки единым целом или кусочками

		case data := <-clim.ChannelsModuleAPIRequestProcessing.OutputModule:
			fmt.Printf("func 'Routing', Input data from 'moduleAPIRequestProcessing'. Reseived data: '%v'\n", data)

			go routingflowsmoduleapirequestprocessing.HandlerAssigmentsModuleAPIRequestProcessing(chanSaveLog, &data, tst, clim)

		}
	}
}

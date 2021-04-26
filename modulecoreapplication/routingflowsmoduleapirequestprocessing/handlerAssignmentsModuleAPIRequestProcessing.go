package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduleapirequestprocessing/auxiliaryfunctions"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//HandlerAssigmentsModuleAPIRequestProcessing является обработчиком приходящих JSON сообщений
func HandlerAssigmentsModuleAPIRequestProcessing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data *datamodels.ModuleReguestProcessingChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clim *moddatamodels.ChannelsListInteractingModules) {

	commonMsgReq, err := unmarshalJSONCommonReq(data.Data)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "unmarshalJSONCommonReq",
		}

		if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
			ClientID:         data.ClientID,
			TypeNotification: "danger",
			Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
				FinalResult: "задача отклонена",
				Message:     "ошибка при декодировании JSON документа",
			}),
			C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
		}); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "SendNotificationModuleAPI",
			}
		}

		return
	}

	switch commonMsgReq.Section {
	case "handling stix object":

		/* *** обработчик JSON сообщений со STIX объектами *** */

		section := "обработка структурированных данных"
		taskType := "добавление или обновление структурированных данных"

		l, err := UnmarshalJSONObjectSTIXReq(*commonMsgReq)
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "UnmarshalJSONObjectSTIXReq",
			}

			if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
				ClientID:         data.ClientID,
				TaskID:           commonMsgReq.TaskID,
				Section:          commonMsgReq.Section,
				TypeNotification: "danger",
				Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
					Section:     section,
					TaskType:    taskType,
					FinalResult: "задача отклонена",
					Message:     "ошибка при декодировании JSON документа",
				}),
				C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "SendNotificationModuleAPI",
				}
			}

			return
		}

		//выполняем валидацию полученных STIX объектов
		if err := CheckSTIXObjects(l); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "CheckSTIXObjects",
			}

			if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
				ClientID:         data.ClientID,
				TaskID:           commonMsgReq.TaskID,
				Section:          commonMsgReq.Section,
				TypeNotification: "danger",
				Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
					Section:     section,
					TaskType:    taskType,
					FinalResult: "задача отклонена",
					Message:     "получения невалидный JSON документ",
				}),
				C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "SendNotificationModuleAPI",
				}
			}

			return
		}

		//добавляем информацию о задаче в хранилище задач
		appTaskID, err := tst.AddNewTask(&memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
			TaskGenerator:        data.ModuleGeneratorMessage,
			ClientID:             data.ClientID,
			ClientName:           data.ClientName,
			ClientTaskID:         commonMsgReq.TaskID,
			AdditionalClientName: commonMsgReq.UserNameGeneratedTask,
			Section:              commonMsgReq.Section,
			Command:              "", //в случае с объектами STIX команда не указывается (автоматически подразумевается добавление или обновление объектов STIX)
			TaskParameters:       SanitizeSTIXObject(l),
		})
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "AddNewTask",
			}

			return
		}

		fmt.Printf("List STIX object:'%v'\n", l)
		fmt.Printf("Application task ID: '%s'\n", appTaskID)

		clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module core application",
				ModuleReceiverMessage:  "module database interaction",
			},
			Section:   "handling stix object",
			AppTaskID: appTaskID,
		}

	case "handling search requests":

		/* *** обработчик JSON сообщений с запросами к поисковой машине приложения *** */

		section := "обработка поискового запроса"
		taskType := "поиск структурированных данных"

		l, err := UnmarshalJSONObjectReqSearchParameters(commonMsgReq.RequestDetails)
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "UnmarshalJSONObjectReqSearchParameters",
			}

			if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
				ClientID:         data.ClientID,
				TaskID:           commonMsgReq.TaskID,
				Section:          commonMsgReq.Section,
				TypeNotification: "danger",
				Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
					Section:     section,
					TaskType:    taskType,
					FinalResult: "задача отклонена",
					Message:     "ошибка при декодировании JSON документа",
				}),
				C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "SendNotificationModuleAPI",
				}
			}

			return
		}

		fmt.Printf("Search data in STIX object:'%v'\n", l)

		//выполняем валидацию и санитаризацию поискового запроса
		l, err = CheckSearchSTIXObject(&l)
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "CheckSearchSTIXObject",
			}

			if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
				ClientID:         data.ClientID,
				TaskID:           commonMsgReq.TaskID,
				Section:          commonMsgReq.Section,
				TypeNotification: "danger",
				Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
					Section:     section,
					TaskType:    taskType,
					FinalResult: "задача отклонена",
					Message:     "получены невалидные параметры поискового запроса",
				}),
				C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "SendNotificationModuleAPI",
				}
			}

			return
		}

		//добавляем информацию о задаче в хранилище задач
		appTaskID, err := tst.AddNewTask(&memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
			TaskGenerator:        data.ModuleGeneratorMessage,
			ClientID:             data.ClientID,
			ClientName:           data.ClientName,
			ClientTaskID:         commonMsgReq.TaskID,
			AdditionalClientName: commonMsgReq.UserNameGeneratedTask,
			Section:              commonMsgReq.Section,
			Command:              "", //в случае с запросом к поисковой машине, команда не указывается
			TaskParameters:       l,
		})
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "AddNewTask",
			}

			return
		}

		clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module core application",
				ModuleReceiverMessage:  "module database interaction",
			},
			Section:   "handling search requests",
			AppTaskID: appTaskID,
		}

	case "handling reference book":

		/* *** обработчик JSON сообщений с параметрами связанными со справочниками *** */

	case "":

		/* *** обработчик JSON сообщений с иными запросами  *** */

	}
}

//unmarshalJSONCommonReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', только по общим полям
func unmarshalJSONCommonReq(msgReq *[]byte) (*datamodels.ModAPIRequestProcessingReqJSON, error) {
	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
	err := json.Unmarshal(*msgReq, &modAPIRequestProcessingReqJSON)

	return &modAPIRequestProcessingReqJSON, err
}

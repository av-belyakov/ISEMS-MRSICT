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

//handlingStatisticalRequests обработчик JSON сообщений со статистическими запросами
func handlingStatisticalRequests(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data *datamodels.ModuleReguestProcessingChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	commonMsgReq *datamodels.ModAPIRequestProcessingReqJSON,
	clim *moddatamodels.ChannelsListInteractingModules) {

	var (
		fn              = commonlibs.GetFuncName()
		section  string = "обработка статистического запроса"
		taskType string = "статистика по структурированным данным"
		statReq  struct {
			CollectionName             string `json:"collection_name"`
			TypeStatisticalInformation string `json:"type_statistical_information"`
		}
	)

	if err := json.Unmarshal(*commonMsgReq.RequestDetails, &statReq); err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    fn,
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

	//добавляем информацию о запросе клиента в лог-файл
	chanSaveLog <- modulelogginginformationerrors.LogMessageType{
		TypeMessage: "requests",
		Description: fmt.Sprintf("Client ID: '%s' (%s), task ID: '%s', section: '%s'", data.ClientID, data.ClientName, commonMsgReq.TaskID, commonMsgReq.Section),
		FuncName:    "SendNotificationModuleAPI",
	}

	//добавляем информацию о задаче в хранилище задач
	appTaskID, err := tst.AddNewTask(&memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
		TaskGenerator:        data.ModuleGeneratorMessage,
		ClientID:             data.ClientID,
		ClientName:           data.ClientName,
		ClientTaskID:         commonMsgReq.TaskID,
		AdditionalClientName: commonMsgReq.UserNameGeneratedTask,
		Section:              commonMsgReq.Section,
		Command:              "",
		TaskParameters:       statReq,
	})
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "AddNewTask",
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
				Message:     "невозможно сохранить параметры запроса во временном хранилище",
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

	clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module database interaction",
		},
		Section:   "handling statistical requests",
		AppTaskID: appTaskID,
	}
}

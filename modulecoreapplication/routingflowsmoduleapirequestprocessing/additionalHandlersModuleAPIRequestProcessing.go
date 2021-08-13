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

//getListComputerThreat обработчик JSON сообщений с запросами списков "types decisions made computer threat" ('типы принимаемых решений по компьютерным
// угрозам') и "types computer threat" ('типы компьютерных угроз')
func getListComputerThreat(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	req *datamodels.ModAPIRequestProcessingResJSONSearchReqType,
	data *datamodels.ModuleReguestProcessingChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	commonMsgReq *datamodels.ModAPIRequestProcessingReqJSON,
	clim *moddatamodels.ChannelsListInteractingModules) {

	var (
		err                error
		section            = "обработка запроса списков типов и решений по компьютерным угрозам"
		taskType           = "поиск информации о типов и решений по компьютерным угрозам"
		fn                 = commonlibs.GetFuncName()
		listComputerThreat map[string]datamodels.StorageApplicationCommonListType
		/*msgRes = datamodels.ModAPIRequestProcessingResJSON{
			ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
				TaskID:  ti.ClientTaskID,
				Section: data.Section,
			},
			IsSuccessful: true,
		}*/
	)

	sp, ok := req.SearchParameters.(struct {
		TypeList string `json:"type_list"`
	})
	if !ok {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: "type conversion error",
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

	switch sp.TypeList {
	case "types decisions made computer threat":
		listComputerThreat, err = tst.GetListDecisionsMade()

	case "types computer threat":
		listComputerThreat, err = tst.GetListComputerThreat()

	default:
		err = fmt.Errorf("undefined type of computer threat list")

	}

	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: "undefined type of computer threat list",
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

	fmt.Printf("func '%s', list computer threat '%v'\n", fn, listComputerThreat)

}

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

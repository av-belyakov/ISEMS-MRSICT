package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"
	"regexp"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduleapirequestprocessing/auxiliaryfunctions"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//handlingManagingCollectionSTIXObjects обработчик JSON сообщений связанных с управлением STIX объектами
func handlingManagingCollectionSTIXObjects(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data datamodels.ModuleReguestProcessingChannel,
	commonMsgReq *datamodels.ModAPIRequestProcessingReqJSON,
	clim *moddatamodels.ChannelsListInteractingModules,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	var (
		err            error
		fn             = commonlibs.GetFuncName()
		isElementExist bool
		section        string = "обработка запросов, связанных с управлением STIX объектами"
		taskType       string = "удаление выбранных STIX объектов"
	)
	at := struct {
		ActionType string `json:"action_type"`
	}{}

	if err := json.Unmarshal(*commonMsgReq.RequestDetails, &at); err != nil {
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

	switch at.ActionType {
	case "delete":
		managementType := struct {
			ActionType   string   `json:"action_type"`
			ListElements []string `json:"list_elements"`
		}{}

		if err := json.Unmarshal(*commonMsgReq.RequestDetails, &managementType); err != nil {
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
			Description: fmt.Sprintf("Client ID: '%s' (%s), task ID: '%s', section: '%s', command: '%s'", data.ClientID, data.ClientName, commonMsgReq.TaskID, commonMsgReq.Section, at.ActionType),
			FuncName:    "SendNotificationModuleAPI",
		}

		listDecisionsMade, errldm := tst.GetListDecisionsMade()
		listComputerThreat, errct := tst.GetListComputerThreat()
		if (errldm != nil) || (errct != nil) {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: "it is not possible to get a list of decisions made on computer threats or a list of types of computer threats",
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
					Message:     "невозможно получить список принимаемых решений по компьютерным угрозам или список типов компьютерных угроз",
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

		//проверяем ID STIX объекта, разрешаем удаление только типов 'grouping' и 'relationship'
		for _, v := range managementType.ListElements {
			isGrouping := (regexp.MustCompile(`^(grouping--)[0-9a-f|-]+$`).MatchString(v))
			isRelationship := (regexp.MustCompile(`^(relationship--)[0-9a-f|-]+$`).MatchString(v))

			if !isGrouping && !isRelationship {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: "it is not possible to delete objects that are not of the 'grouping' or 'relationship' types",
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
						Message:     "невозможно удалить объекты не являющиеся типами 'grouping' или 'relationship'",
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

			//проверяем относится ли удаляемый STIX объект типа 'grouping' к специальным объектам списка принимаемых решений по компьютерным угрозам
			// или списка типов компьютерных угроз
			if !isGrouping {
				continue
			}

		DONELDM:
			for k := range listDecisionsMade {
				if k == v {
					isElementExist = true

					break DONELDM
				}
			}

		DONELCT:
			for k := range listComputerThreat {
				if k == v {
					isElementExist = true

					break DONELCT
				}
			}

			if isElementExist {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: "you cannot delete the STIX element of an object because it is an element of a preset list",
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
						Message:     "нельзя удалить элемент STIX объекта так как он является элементом предустановленного списка ",
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
		}

		//добавляем информацию о задаче в хранилище задач
		appTaskID, err := tst.AddNewTask(&memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
			TaskGenerator:        data.ModuleGeneratorMessage,
			ClientID:             data.ClientID,
			ClientName:           data.ClientName,
			ClientTaskID:         commonMsgReq.TaskID,
			AdditionalClientName: commonMsgReq.UserNameGeneratedTask,
			Section:              commonMsgReq.Section,
			Command:              managementType.ActionType,
			TaskParameters:       managementType.ListElements,
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
			Section:   "handling managing collection stix objects",
			AppTaskID: appTaskID,
		}

	case "":

	}
}

//getListComputerThreat обработчик JSON сообщений с запросами списков "types decisions made computer threat" ('типы принимаемых решений по компьютерным
// угрозам') и "types computer threat" ('типы компьютерных угроз')
func getListComputerThreat(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	searchParameters interface{},
	clientID string,
	commonMsgReq *datamodels.ModAPIRequestProcessingReqJSON,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (string, error) {

	var (
		err                error
		fn                 = commonlibs.GetFuncName()
		listComputerThreat map[string]datamodels.StorageApplicationCommonListType
		msgRes             = datamodels.ModAPIRequestProcessingResJSON{
			ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
				TaskID:  commonMsgReq.TaskID,
				Section: commonMsgReq.Section,
			},
			IsSuccessful: true,
		}
	)

	sp, ok := searchParameters.(struct {
		TypeList string `json:"type_list"`
	})
	if !ok {
		return fn, fmt.Errorf("type conversion error, line 308")
	}

	switch sp.TypeList {
	case "types decisions made computer threat":
		listComputerThreat, err = tst.GetListDecisionsMade()
		if err != nil {
			return fn, err
		}

	case "types computer threat":
		listComputerThreat, err = tst.GetListComputerThreat()
		if err != nil {
			return fn, err
		}

	default:
		return fn, fmt.Errorf("undefined type of computer threat list")

	}

	msgRes.AdditionalParameters = struct {
		TypeList string                                                 `json:"type_list"`
		List     map[string]datamodels.StorageApplicationCommonListType `json:"list"`
	}{
		TypeList: sp.TypeList,
		List:     listComputerThreat,
	}

	msg, err := json.Marshal(msgRes)
	if err != nil {
		return fn, err
	}

	chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module api request processing",
		},
		ClientID: clientID,
		DataType: 1,
		Data:     &msg,
	}

	return fn, nil
}

//handlingStatisticalRequests обработчик JSON сообщений со статистическими запросами
func handlingStatisticalRequests(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data datamodels.ModuleReguestProcessingChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	commonMsgReq *datamodels.ModAPIRequestProcessingReqJSON,
	clim *moddatamodels.ChannelsListInteractingModules) {

	var (
		fn              = commonlibs.GetFuncName()
		section  string = "обработка статистического запроса"
		taskType string = "статистика по структурированным данным"
		statReq         = datamodels.CommonStatisticalRequest{}
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

	fmt.Printf("func 'additionalHandlersModuleAPIRequestProcessing', Section: '%s', appTaskID: '%s'\n", commonMsgReq.Section, appTaskID)

	clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module database interaction",
		},
		Section:   "handling statistical requests",
		AppTaskID: appTaskID,
	}
}

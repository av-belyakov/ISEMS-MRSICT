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
func HandlerAssignmentsModuleAPIRequestProcessing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data datamodels.ModuleReguestProcessingChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clim *moddatamodels.ChannelsListInteractingModules) {

	commonMsgReq, err := unmarshalJSONCommonReq(data.Data)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprintln(err),
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
					Message:     "получен невалидный JSON документ",
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

		//обрабатываем содержимое полей которые не относятся к спецификации STIX 2.0
		VerifyOutsideSpecificationFields(l, tst, data.ClientName)

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

		clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module core application",
				ModuleReceiverMessage:  "module database interaction",
			},
			Section:   "handling stix object",
			AppTaskID: appTaskID,
		}

	case "managing collection stix objects":

		/* *** обработчик JSON сообщений связанных с управлением STIX объектами *** */

		handlingManagingCollectionSTIXObjects(chanSaveLog, data, commonMsgReq, clim, tst)

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

		switch l.CollectionName {
		case "stix object":
			/* *** обработчик JSON сообщений с общими запросами поиска по коллекции STIX объектов *** */

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

		case "get list computer threat":
			/* ***
				обработчик JSON запросов, выполняемых с целью получения списков "types decisions made computer threat" ('типы принимаемых решений
				по компьютерным угрозам') или "types computer threat" ('типы компьютерных угроз'). Эта информация нужна для построения выпадающих
				списков в ISEMS-UI
			*** */

			section = "обработка запроса списков типов и решений по компьютерным угрозам"
			taskType = "поиск информации о типов и решений по компьютерным угрозам"

			if funcName, err := getListComputerThreat(clim.ChannelsModuleAPIRequestProcessing.InputModule, l.SearchParameters, data.ClientID, commonMsgReq, tst); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprintln(err),
					FuncName:    funcName,
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
						Message:     "получены некорректные параметры запроса",
					}),
					C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
				}); err != nil {
					chanSaveLog <- modulelogginginformationerrors.LogMessageType{
						TypeMessage: "error",
						Description: fmt.Sprint(err),
						FuncName:    "SendNotificationModuleAPI",
					}
				}
			}

			return

		case "differences objects collection":
			/* ***
				обработчик JSON сообщений с запросами поиска информации о предыдущем состоянии STIX объектов,
				находящейся в коллекции "accounting_differences_objects_collection"
			*** */

			fmt.Println("func 'HandlerAssignmentsModuleAPIRequestProcessing', START...")
			fmt.Printf("commonMsgReq.Section: '%s', l.CollectionName: '%s'\n", commonMsgReq.Section, l.CollectionName)
			fmt.Println(l.SearchParameters)

			/*
							search_parameters: {
				            	document_id: STRING // идентификатор документа в котором выполнялись модификации
				            	collection_name: STRING // наименование коллекции в которой выполнялись модификации
				        	}
			*/

		default:
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: "the 'collection_name' parameter is not defined or has an invalid value",
				FuncName:    "HandlerAssigmentsModuleAPIRequestProcessing",
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
					//Message:     "получено невалидное название коллекции в которой должен был быть выполнен поиск",
					Message: fmt.Sprintf("получено невалидное название коллекции в которой должен был быть выполнен поиск, collection name: (%s)", l.CollectionName),
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
			Command:              "", //в случае с запросом к поисковой машине, команда не указывается
			TaskParameters:       l,
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
			Section:   "handling search requests",
			AppTaskID: appTaskID,
		}

	case "handling statistical requests":

		/* *** обработчик JSON сообщений со статистическими запросами *** */

		handlingStatisticalRequests(chanSaveLog, data, tst, commonMsgReq, clim)

	case "handling reference book":

		section := "обработка справочной информации"
		taskType := "выполнение действий над данными"

		/* *** обработчик JSON сообщений с параметрами связанными со справочниками *** */
		//l, err := UnmarshalJSONReferenceBookReq(*commonMsgReq)
		l, err := UnmarshalJSONRBookReq(commonMsgReq.RequestDetails)
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "UnmarshalJSONRBookReq",
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

		//выполняем валидацию полученных запросов к справочной информации
		if _, err = l.IsValid(); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "IsValid",
			}

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
			Command:              "", //в случае с объектами ReferenceBook команда не указывается (Для каждого отдельного элемента применяется своя командакоманда присутствует в каждом элементе среза)
			TaskParameters:       SanitizeReqRBObject(l),
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
		}

		clim.ChannelsModuleDataBaseInteraction.ChannelsMongoDB.InputModule <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module core application",
				ModuleReceiverMessage:  "module database interaction",
			},
			Section:   "handling reference book",
			AppTaskID: appTaskID,
		}

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

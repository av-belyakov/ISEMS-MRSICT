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

//HandlerAssigmentsModuleAPIRequestProcessing является обработчиком приходящих JSON сообщений
func HandlerAssigmentsModuleAPIRequestProcessing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data *datamodels.ModuleReguestProcessingChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clim *moddatamodels.ChannelsListInteractingModules) {

	var listTypesComputerThreat = []string{
		"types decisions made computer threat",
		"types computer threat",
	}

	commonMsgReq, err := unmarshalJSONCommonReq(data.Data)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
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

		var (
			isElementExist bool
			section        string = "обработка запросов, связанных с управлением STIX объектами"
			taskType       string = "удаление выбранных STIX объектов"
			fn             string = commonlibs.GetFuncName()
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
			/*
				общий поиск по коллекции STIX объектов
			*/
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

		case "get list computer threat", "stix object list type grouping":
			/*
				получить список id объектов типа 'grouping' с их наименованием, относящихся или к "types decisions made computer threat"
				или к "types computer threat", данный список нужен для построения выпадающих списков в ISEMS-UI, для первого совпадения или
				получить СПИСОК STIX объектов типа 'Grouping' относящихся, к заранее определенному в приложении, списку, 'типы принимаемых решений по
				компьютерным угрозам' (types decisions made computer threat) или 'типы компьютерных угроз' (types computer threat), для второго совпадения
			*/

			if l.SortableField != "name" {
				l.SortableField = "name"
			}

			sp, ok := l.SearchParameters.(struct {
				TypeList string `json:"type_list"`
			})
			if !ok {
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

			var isExist bool
			for _, v := range listTypesComputerThreat {
				if v == sp.TypeList {
					isExist = true

					break
				}
			}

			if !isExist {
				if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
					ClientID:         data.ClientID,
					TaskID:           commonMsgReq.TaskID,
					Section:          commonMsgReq.Section,
					TypeNotification: "danger",
					Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
						Section:     section,
						TaskType:    taskType,
						FinalResult: "задача отклонена",
						Message:     "получено невалидное значение поискового запроса",
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

			l.SearchParameters = sp

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
					Message: fmt.Sprintf("получено невалидное название коллекции в которой должен был быть выполнен поиск, collection name: '%s'", l.CollectionName),
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

	case "handling reference book":

		section := "обработка справочной информации"
		taskType := "выполнение действий над данными"

		/* *** обработчик JSON сообщений с параметрами связанными со справочниками *** */
		//l, err := UnmarshalJSONReferenceBookReq(*commonMsgReq)
		/* *** обработчик JSON сообщений с параметрами связанными со справочниками **** */
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
				FuncName:    "UnmarshalJSONReferenceBookReq",
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

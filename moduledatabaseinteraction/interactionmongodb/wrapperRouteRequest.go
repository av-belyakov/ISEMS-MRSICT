package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/commonhandlers"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
)

//WrapperFuncTypeHandlingSTIXObject набор обработчиков для работы с запросами связанными со STIX объектами
func (ws *wrappersSetting) wrapperFuncTypeHandlingSTIXObject(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	var (
		err error
		fn  = "wrapperFuncTypeHandlingSTIXObject"
		qp  = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "stix_object_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
	)

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
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
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	ti, ok := taskInfo.TaskParameters.([]*datamodels.ElementSTIXObject)
	if !ok {
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   fmt.Errorf("type conversion error"),
				},
			},
			Section:   "handling stix object",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	//получаем список ID STIX объектов предназначенных для добавление в БД
	listID := commonhandlers.GetListIDFromListSTIXObjects(ti)

	//выполняем запрос к БД, для получения полной информации об STIX объектах по их ID
	listElemetSTIXObject, err := FindSTIXObjectByID(qp, listID)
	if err != nil {
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   err,
				},
			},
			Section:   "handling stix object",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	//выполняем сравнение объектов и ищем внесенные изменения для каждого из STIX объектов
	listDifferentObject := ComparasionListSTIXObject(ComparasionListTypeSTIXObject{
		CollectionType: qp.CollectionName,
		OldList:        listElemetSTIXObject,
		NewList:        ti,
	})

	//логируем изменения в STIX объектах в отдельную коллекцию 'accounting_differences_objects_collection'
	if len(listDifferentObject) > 0 {
		qp.CollectionName = "accounting_differences_objects_collection"

		_, err := qp.InsertData([]interface{}{listDifferentObject})
		if err != nil {
			chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
				CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
					ModuleGeneratorMessage: "module database interaction",
					ModuleReceiverMessage:  "module core application",
					ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
						FuncName:                                fn,
						ModuleAPIRequestProcessingSettingSendTo: true,
						Error:                                   err,
					},
				},
				Section:   "handling stix object",
				AppTaskID: ws.DataRequest.AppTaskID,
			}

			return
		}
	}

	//добавляем или обновляем STIX объекты в БД
	err = ReplacementElementsSTIXObject(qp, ti)
	if err != nil {
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   err,
				},
			},
			Section:   "handling stix object",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module database interaction",
			ModuleReceiverMessage:  "module core application",
			InformationMessage: datamodels.InformationDataTypePassedThroughChannels{
				Type:    "success",
				Message: "информация о структурированных данных успешно добавлена",
			},
		},
		Section:   "handling stix object",
		AppTaskID: ws.DataRequest.AppTaskID,
	}
}

//wrapperFuncTypeHandlingSearchRequests набор обработчиков для работы с запросами направленными на обработку поисковой машине
func (ws *wrappersSetting) wrapperFuncTypeHandlingSearchRequests(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	var (
		err           error
		fn            = "wrapperFuncTypeHandlingSearchRequests"
		sortableField string
		qp            = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "stix_object_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
		sf = map[string]string{
			"document_type":   "commonpropertiesobjectstix.type",
			"data_created":    "commonpropertiesdomainobjectstix.created",
			"data_modified":   "commonpropertiesdomainobjectstix.modified",
			"data_first_seen": "first_seen",
			"data_last_seen":  "last_seen",
			"ipv4":            "value",
			"ipv6":            "value",
			"country":         "country",
		}
	)

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
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
			Section:   "handling search requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	psr, ok := taskInfo.TaskParameters.(datamodels.ModAPIRequestProcessingResJSONSearchReqType)
	if !ok {
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   fmt.Errorf("type conversion error"),
				},
			},
			Section:   "handling search requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	//изменяем время модификации информации о задаче
	_ = tst.ChangeDateTaskModification(ws.DataRequest.AppTaskID)

	//изменяем статус выполняемой задачи на 'in progress'
	if err := tst.ChangeTaskStatus(ws.DataRequest.AppTaskID, "in progress"); err != nil {
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   err,
				},
			},
			Section:   "handling search requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}

	switch psr.CollectionName {
	case "stix object":
		searchParameters, ok := psr.SearchParameters.(datamodels.SearchThroughCollectionSTIXObjectsType)
		if !ok {
			chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
				CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
					ModuleGeneratorMessage: "module database interaction",
					ModuleReceiverMessage:  "module core application",
					ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
						FuncName:                                fn,
						ModuleAPIRequestProcessingSettingSendTo: true,
						Error:                                   fmt.Errorf("type conversion error"),
					},
				},
				Section:   "handling search requests",
				AppTaskID: ws.DataRequest.AppTaskID,
			}

			return
		}

		//получить только общее количество найденных документов
		if (psr.PaginateParameters.CurrentPartNumber <= 0) || (psr.PaginateParameters.MaxPartNum <= 0) {
			resSize, err := qp.CountDocuments(CreateSearchQueriesSTIXObject(&searchParameters))
			if err != nil {
				chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
					CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
						ModuleGeneratorMessage: "module database interaction",
						ModuleReceiverMessage:  "module core application",
						ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
							FuncName:                                fn,
							ModuleAPIRequestProcessingSettingSendTo: true,
							Error:                                   err,
						},
					},
					Section:   "handling search requests",
					AppTaskID: ws.DataRequest.AppTaskID,
				}

				return
			}

			//сохраняем общее количество найденных значений во временном хранилище
			err = tst.AddNewFoundInformation(
				ws.DataRequest.AppTaskID,
				&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
					Collection:  "stix_object_collection",
					ResultType:  "only_count",
					Information: resSize,
				})
			if err != nil {
				chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
					CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
						ModuleGeneratorMessage: "module database interaction",
						ModuleReceiverMessage:  "module core application",
						ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
							FuncName:                                fn,
							ModuleAPIRequestProcessingSettingSendTo: true,
							Error:                                   err,
						},
					},
					Section:   "handling search requests",
					AppTaskID: ws.DataRequest.AppTaskID,
				}

				return
			}
		} else {
			if field, ok := sf[psr.SortableField]; ok {
				sortableField = field
			}

			//получить все найденные документы, с учетом лимита
			cur, err := qp.FindAllWithLimit(CreateSearchQueriesSTIXObject(&searchParameters), &FindAllWithLimitOptions{
				Offset:        int64(psr.PaginateParameters.CurrentPartNumber),
				LimitMaxSize:  int64(psr.PaginateParameters.MaxPartNum),
				SortField:     sortableField,
				SortAscending: false,
			})
			if err != nil {
				chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
					CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
						ModuleGeneratorMessage: "module database interaction",
						ModuleReceiverMessage:  "module core application",
						ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
							FuncName:                                fn,
							ModuleAPIRequestProcessingSettingSendTo: true,
							Error:                                   err,
						},
					},
					Section:   "handling search requests",
					AppTaskID: ws.DataRequest.AppTaskID,
				}

				return
			}

			//сохраняем найденные значения во временном хранилище
			err = tst.AddNewFoundInformation(
				ws.DataRequest.AppTaskID,
				&memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation{
					Collection:  "stix_object_collection",
					ResultType:  "full_found_info",
					Information: GetListElementSTIXObject(cur),
				})
			if err != nil {
				chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
					CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
						ModuleGeneratorMessage: "module database interaction",
						ModuleReceiverMessage:  "module core application",
						ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
							FuncName:                                fn,
							ModuleAPIRequestProcessingSettingSendTo: true,
							Error:                                   err,
						},
					},
					Section:   "handling search requests",
					AppTaskID: ws.DataRequest.AppTaskID,
				}

				return
			}
		}

		_ = tst.ChangeDateTaskModification(ws.DataRequest.AppTaskID)

		//изменяем состояние задачи на 'completed'
		if err := tst.ChangeTaskStatus(ws.DataRequest.AppTaskID, "completed"); err != nil {
			chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
				CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
					ModuleGeneratorMessage: "module database interaction",
					ModuleReceiverMessage:  "module core application",
					ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
						FuncName:                                fn,
						ModuleAPIRequestProcessingSettingSendTo: true,
						Error:                                   err,
					},
				},
				Section:   "handling search requests",
				AppTaskID: ws.DataRequest.AppTaskID,
			}

			return
		}

		//отправляем в канал идентификатор задачи и специальные параметры которые информируют что задача была выполненна
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
			},
			Section:   "handling search requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

	case "":

	default:
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   fmt.Errorf("the name of the database collection is not defined"),
				},
			},
			Section:   "handling search requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}
	}
}

//wrapperFuncTypeHandlingReferenceBook набор обработчиков для работы с запросами к справочнику
func (ws *wrappersSetting) wrapperFuncTypeHandlingReferenceBook(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {
	var (
		err error
		fn  = " wrapperFuncTypeHandlingReferenceBook"
		qp  = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "reference_book_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
	)
	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
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
			Section:   "reference book requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}
	pbr, ok := taskInfo.TaskParameters.(datamodels.RBookReq)
	if !ok {
		chanOutput <- datamodels.ModuleDataBaseInteractionChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module database interaction",
				ModuleReceiverMessage:  "module core application",
				ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
					FuncName:                                fn,
					ModuleAPIRequestProcessingSettingSendTo: true,
					Error:                                   fmt.Errorf("type conversion error"),
				},
			},
			Section:   "reference book requests",
			AppTaskID: ws.DataRequest.AppTaskID,
		}

		return
	}
	/*switch wt.command {
	case "find_all":

	case "find_all_for_client_API":

	case "":

	}*/
}

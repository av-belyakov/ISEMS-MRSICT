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

	fmt.Println("func 'wrapperFuncTypeHandlingSTIXObject', START...")
	fmt.Printf("func 'wrapperFuncTypeHandlingSTIXObject', received message: '%v'\n", ws)

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

	fmt.Printf("func 'wrapperFuncTypeHandlingSTIXObject', task info: '%v'\n", taskInfo)

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

	fmt.Printf("func 'wrapperFuncTypeHandlingSTIXObject', list ID: '%v'\n", listID)

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

	fmt.Println("func 'wrapperFuncTypeHandlingSearchRequests', START...")

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

	fmt.Printf("func 'wrapperFuncTypeHandlingSearchRequests', task info: '%v'\n", taskInfo)

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

			fmt.Printf("func '%s', search for collection name 'stix object', RESULT COUNT ELEMENTS: '%d'\n", fn, resSize)

			//сохраняем общее количество найденных значений во временном хранилище

			//отправляем в канал идентификатор задачи и специальные параметры которые информируют что задача была выполненна

			return
		}

		if field, ok := sf[psr.SortableField]; ok {
			sortableField = field
		}

		fmt.Printf("func '%s', search for collection name 'stix object', SORTABLE FIELD: '%s'\n", fn, sortableField)

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

		result := GetListElementSTIXObject(cur)

		fmt.Printf("func '%s', search for collection name 'stix object', RESULT: '%v'\n", fn, result)

		/*
		   Надо проверить сортировку по полям которые, возможно, окажутся не во всех найденных документах.
		   Как при этом поведет себя MongoDB, не будет ли ошибки?

		*/

		//сохраняем найденные значения во временном хранилище

		//отправляем в канал идентификатор задачи и специальные параметры которые информируют что задача была выполненна

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
	/*switch wt.command {
	case "find_all":

	case "find_all_for_client_API":

	case "":

	}*/
}

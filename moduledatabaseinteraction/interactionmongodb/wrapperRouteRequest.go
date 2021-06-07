package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/commonhandlers"
	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"

	"go.mongodb.org/mongo-driver/mongo"
)

var errorMessage = datamodels.ModuleDataBaseInteractionChannel{
	CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
		ModuleGeneratorMessage: "module database interaction",
		ModuleReceiverMessage:  "module core application",
		ErrorMessage: datamodels.ErrorDataTypePassedThroughChannels{
			ModuleAPIRequestProcessingSettingSendTo: true,
		},
	},
}

//wrapperFuncTypeHandlingSTIXObject набор обработчиков для работы с запросами связанными со STIX объектами
func (ws *wrappersSetting) wrapperFuncTypeHandlingSTIXObject(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	var (
		err error
		fn  = commonlibs.GetFuncName() //"wrapperFuncTypeHandlingSTIXObject"
		qp  = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "stix_object_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
	)

	errorMessage.ErrorMessage.FuncName = fn
	errorMessage.Section = "handling stix object"
	errorMessage.AppTaskID = ws.DataRequest.AppTaskID

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
		errorMessage.ErrorMessage.Error = fmt.Errorf("no information about the task by its id was found in the temporary storage")
		chanOutput <- errorMessage

		return
	}

	ti, ok := taskInfo.TaskParameters.([]*datamodels.ElementSTIXObject)
	if !ok {
		errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
		chanOutput <- errorMessage

		return
	}

	//получаем список ID STIX объектов предназначенных для добавление в БД
	listID := commonhandlers.GetListIDFromListSTIXObjects(ti)

	//выполняем запрос к БД, для получения полной информации об STIX объектах по их ID
	listElemetSTIXObject, err := FindSTIXObjectByID(qp, listID)
	if err != nil {
		errorMessage.ErrorMessage.Error = err
		chanOutput <- errorMessage

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

		_, err := qp.InsertData([]interface{}{listDifferentObject}, []mongo.IndexModel{})
		if err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}
	}

	//добавляем или обновляем STIX объекты в БД
	err = ReplacementElementsSTIXObject(qp, ti)
	if err != nil {
		errorMessage.ErrorMessage.Error = err
		chanOutput <- errorMessage

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
		fn            = commonlibs.GetFuncName() //"wrapperFuncTypeHandlingSearchRequests"
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

	errorMessage.ErrorMessage.FuncName = fn
	errorMessage.Section = "handling search requests"
	errorMessage.AppTaskID = ws.DataRequest.AppTaskID

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
		errorMessage.ErrorMessage.Error = err
		chanOutput <- errorMessage

		return
	}

	psr, ok := taskInfo.TaskParameters.(datamodels.ModAPIRequestProcessingResJSONSearchReqType)
	if !ok {
		errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
		chanOutput <- errorMessage

		return
	}

	//изменяем время модификации информации о задаче
	_ = tst.ChangeDateTaskModification(ws.DataRequest.AppTaskID)

	//изменяем статус выполняемой задачи на 'in progress'
	if err := tst.ChangeTaskStatus(ws.DataRequest.AppTaskID, "in progress"); err != nil {
		errorMessage.ErrorMessage.Error = err
		chanOutput <- errorMessage

		return
	}

	switch psr.CollectionName {
	case "stix object":
		searchParameters, ok := psr.SearchParameters.(datamodels.SearchThroughCollectionSTIXObjectsType)
		if !ok {
			errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
			chanOutput <- errorMessage

			return
		}

		//получить только общее количество найденных документов
		if (psr.PaginateParameters.CurrentPartNumber <= 0) || (psr.PaginateParameters.MaxPartNum <= 0) {
			resSize, err := qp.CountDocuments(CreateSearchQueriesSTIXObject(&searchParameters))
			if err != nil {
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

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
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

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
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

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
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

				return
			}
		}

		_ = tst.ChangeDateTaskModification(ws.DataRequest.AppTaskID)

		//изменяем состояние задачи на 'completed'
		if err := tst.ChangeTaskStatus(ws.DataRequest.AppTaskID, "completed"); err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

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

	case "stix object list type grouping":
		if fn, err := searchSTIXObjectListTypeGrouping(ws.DataRequest.AppTaskID, qp, psr, tst); err != nil {
			errorMessage.ErrorMessage.FuncName = fn
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

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

	default:
		errorMessage.CommanDataTypePassedThroughChannels.ErrorMessage.Error = fmt.Errorf("the name of the database collection is not defined")
		chanOutput <- errorMessage

	}
}

//switchMSGType - функция заполняющая одно из информационных полей cообщения
// распознавая тип объекта передаваемого в нее
func switchMSGType(msg *datamodels.ModuleDataBaseInteractionChannel, m interface{}) bool {
	msg.ErrorMessage = datamodels.ErrorDataTypePassedThroughChannels{}
	msg.InformationMessage = datamodels.InformationDataTypePassedThroughChannels{}
	switch m.(type) {
	case datamodels.ErrorDataTypePassedThroughChannels:
		msg.ErrorMessage = m.(datamodels.ErrorDataTypePassedThroughChannels)
		return true
	case datamodels.InformationDataTypePassedThroughChannels:
		msg.InformationMessage = m.(datamodels.InformationDataTypePassedThroughChannels)
		return true
	default:
		return false
	}
}

//wrapperFuncTypeTechnicalPart набор обработчиков для осуществления задач, связанных с технической частью приложения: формирование документов БД
// связанных с хранением технической информации или документов, учавствующих в посторении иерархии объектов типа STIX. Запись идентификаторов таких
// объектов во временное хранилище и т.д.
func (ws *wrappersSetting) wrapperFuncTypeTechnicalPart(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	fmt.Println("func 'wrapperFuncTypeTechnicalPart', START...")

	var (
		fn string = "wrapperFuncTypeTechnicalPart"
		qp        = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "stix_object_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
	)

	errorMessage.ErrorMessage.FuncName = fn
	errorMessage.Section = "handling technical part"

	switch ws.DataRequest.Command {
	case "create STIX DO type 'grouping'":
		/*
			проверяем наличие объектов STIX DO типа 'grouping', содержащих списки 'подтвержденных' или 'отклоненных' объектов STIX DO типа 'report'
			и при необходимости создаем новые STIX DO объекты типа 'grouping'
		*/
		go func() {
			ldm, err := tst.GetListDecisionsMade()
			if err != nil {
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

				return
			}

			listID, err := GetIDGroupingObjectSTIX(qp, ldm)
			if err != nil {
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

				return
			}

			//добавляем список ID во временное хранилище
			tst.SetListDecisionsMade(listID)
		}()

		/*
			проверяем наличие объектов STIX DO типа 'grouping', содержащих списки объектов STIX DO типа 'report', относящихся к какому то определенному
			виду компьютерного воздействия
		*/
		go func() {
			lct, err := tst.GetListComputerThreat()
			if err != nil {
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

				return
			}

			listID, err := GetIDGroupingObjectSTIX(qp, lct)
			if err != nil {
				errorMessage.ErrorMessage.Error = err
				chanOutput <- errorMessage

				return
			}

			//добавляем список ID во временное хранилище
			tst.SetListComputerThreat(listID)
		}()
	}
}

//wrapperFuncTypeHandlingReferenceBook набор обработчиков для работы с запросами к справочнику
func (ws *wrappersSetting) wrapperFuncTypeHandlingReferenceBook(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {
	var (
		err error
		fn  = commonlibs.GetFuncName()
		qp  = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "reference_book_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
		sortVocs map[string]datamodels.Vocabularys = make(map[string]datamodels.Vocabularys,
			len(datamodels.CommandSet))
	)
	message := datamodels.ModuleDataBaseInteractionChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module database interaction",
			ModuleReceiverMessage:  "module core application",
		},
		Section:   "handling reference book request",
		AppTaskID: ws.DataRequest.AppTaskID,
	}

	errorMessage := datamodels.ErrorDataTypePassedThroughChannels{
		FuncName:                                fn,
		ModuleAPIRequestProcessingSettingSendTo: true,
	}

	infoMessage := datamodels.InformationDataTypePassedThroughChannels{
		Type: "",
	}
	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
		errorMessage.Error = err
		switchMSGType(&message, errorMessage)
		chanOutput <- message
		return
	}

	rbrps, ok := taskInfo.TaskParameters.(datamodels.RBookReqParameters)
	if !ok {
		errorMessage.Error = fmt.Errorf("type conversion error")
		switchMSGType(&message, errorMessage)
		chanOutput <- message
		return
	}
	//Сортируем в разные срезы по действиям над объектами
	for _, rbp := range rbrps {
		sortVocs[rbp.OP] = append(sortVocs[rbp.OP], rbp.Vocabulary)
	}

	// Пока что реализуем только добавление

	if listAddVocs, ok := sortVocs["add"]; ok {
		//Отфильтровываем только те объекты RB которые являются редактируемыми
		listAddVocs, listNotEditable := FilterEditabelRB(listAddVocs)
		//О наличии в срезе не редактируемых объектов сообщаем куда следует
		if len(listNotEditable) != 0 {
			strNames := listNotEditable.GetListIDtoStr()
			infoMessage.Message = fmt.Sprintf("Создание и редактирование следующих объектов: %s запрещено", strNames)
			switchMSGType(&message, infoMessage)
			chanOutput <- message
		}
		if len(listAddVocs) == 0 {
			infoMessage.Message = ""
			switchMSGType(&message, infoMessage)
			chanOutput <- message
		}
		//получаем список имен RB
		listID := listAddVocs.GetListID()

		//выполняем запрос к БД, для получения полной информации об уже существующих в коллекции объектах справочниках по их ID
		listFoundRB, err := FindRBObjectsByNames(qp, listID)
		if err != nil {
			errorMessage.Error = err
			switchMSGType(&message, errorMessage)
			chanOutput <- message

			return
		}

		//сравниваем объекты RB хранящиеся в БД и добавляемые и фиксируем их различия
		listDifferentObject := ComparasionListRBbject(listAddVocs, listFoundRB)

		//логируем отличия RB-объектов в отдельную коллекцию 'accounting_differences_objects_collection'
		if len(listDifferentObject) > 0 {
			qp.CollectionName = "accounting_differences_objects_collection"

			_, err := qp.InsertData([]interface{}{listDifferentObject})
			if err != nil {
				errorMessage.Error = err
				switchMSGType(&message, errorMessage)
				chanOutput <- message

				return
			}
		}

		//добавляем или обновляем STIX объекты в БД
		//	err = ReplacementElementsSTIXObject(qp, ti)
		/*	if err != nil {
			errorMessage.Error = err
			switchMSGType(&message, errorMessage)
			chanOutput <- message

			return

		}*/

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

}

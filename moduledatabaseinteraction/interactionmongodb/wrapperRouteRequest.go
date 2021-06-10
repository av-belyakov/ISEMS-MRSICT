package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/commonhandlers"
	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"

	"go.mongodb.org/mongo-driver/bson"
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

//wrapperFuncTypeHandlingSTIXObject набор обработчиков для работы с запросами, связанными со STIX объектами
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

	/*
		СДЕЛАЛ 1. Удаление объектов типа 'grouping' и 'relationship' я написал, НЕОБХОДИМО ТЕСТИРОВАНИЕ, начал писать тест.
		2. Нужно сделать автоматическое установление ОБРАТНЫХ связей между STIX объектами содержащими свойство ObjectRefs, такими как объекты типов:
		- 'grouping'
		- 'report'
		- 'note'
		- 'observed'
		- 'opinion'
		и с любыми другими объектами по средствам объектов типа 'relationship'.
		3. Нужно сделать автоматическое удаление объектов типа 'relationship' обеспечивающие обратную связь между объектами типов 'grouping'
		и 'report' и еще три выше перечисленных и другими объектами при удалении ссылок на объекты из поля ObjectRefs объектов типов
		'grouping' и 'report' и т.д.
	*/

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

//wrapperFuncTypeHandlingManagingCollectionSTIXObjects набор обработчиков для работы с запросами, связанными с управлением STIX объектами
func (ws *wrappersSetting) wrapperFuncTypeHandlingManagingCollectionSTIXObjects(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	var (
		err error
		fn  = commonlibs.GetFuncName()
		qp  = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "stix_object_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
		}
	)

	errorMessage.ErrorMessage.FuncName = fn
	errorMessage.Section = "handling managing collection stix objects"
	errorMessage.AppTaskID = ws.DataRequest.AppTaskID

	//получаем всю информацию о выполняемой задаче из временного хранилища задач
	_, taskInfo, err := tst.GetTaskByID(ws.DataRequest.AppTaskID)
	if err != nil {
		errorMessage.ErrorMessage.Error = err
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

	switch taskInfo.Command {
	case "delete":
		var (
			listIDGroupingDel     []string // список объектов типа 'grouping' предназначеных для удаления
			listIDRelationshipDel []string // список объектов типа 'relationship' предназначеных для удаления
			listIDReporModify     []string // список объектов типа 'report' предназначенных для модификации
			listObjModiy          []*datamodels.ElementSTIXObject
		)
		sl := map[string]struct {
			targetRefsID   string
			relationshipID string
			listRefs       []datamodels.IdentifierTypeSTIX
		}{}

		listID, ok := taskInfo.TaskParameters.([]string)
		if !ok {
			errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
			chanOutput <- errorMessage

			return
		}

		//получаем все объекты предназначенные для удаления (проверка типа объекта, удаление возможно только объектов типа 'grouping' или
		//'relationship', осуществляется на этапе валидации входных параметров)
		listElementSTIXObject, err := FindSTIXObjectByID(qp, listID)
		if err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}

		//обрабатываем все объекты типа 'grouping' и 'relationship' из полученной задачи и отмеченные для удаления
		for _, v := range listElementSTIXObject {
			if v.DataType == "grouping" {
				listIDGroupingDel = append(listIDGroupingDel, v.Data.GetID())

				element, ok := v.Data.(datamodels.GroupingDomainObjectsSTIX)
				if !ok {
					errorMessage.ErrorMessage.Error = fmt.Errorf("type conversion error")
					chanOutput <- errorMessage

					return
				}

				if len(element.ObjectRefs) == 0 {
					continue
				}

				sl[v.Data.GetID()] = struct {
					targetRefsID   string
					relationshipID string
					listRefs       []datamodels.IdentifierTypeSTIX
				}{listRefs: element.ObjectRefs}
			}

			if v.DataType == "relationship" {
				listIDRelationshipDel = append(listIDRelationshipDel, v.Data.GetID())
			}
		}

		//ищем объекты типа 'relationship' являющиеся связующим звеном между объектами 'grouping' и другими объектами, чаще всего 'report'
		cur, err := qp.Find((bson.D{
			bson.E{Key: "commonpropertiesobjectstix.type", Value: "relationship"},
			bson.E{Key: "source_ref", Value: bson.D{{Key: "$in", Value: listIDGroupingDel}}},
		}))
		if err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}

		//обрабатываем полученый список объектов типа 'relationship'
		for _, v := range GetListElementSTIXObject(cur) {
			if obj, ok := v.Data.(datamodels.RelationshipObjectSTIX); ok {
				targetID := string(obj.TargetRef)

				//сохраняем список объектов типа 'relationship' являющихся связующим звеном и которые в последствии необходимо удалить
				listIDRelationshipDel = append(listIDRelationshipDel, obj.ID)
				listIDReporModify = append(listIDReporModify, targetID)

				sl[string(obj.SourceRef)] = struct { // ID объекта типа 'gouping'
					targetRefsID   string                          // объект 'report' на который ссылается какой либо объект из поля SourceRefs
					relationshipID string                          // ID объекта типа 'relationship' который соединяет объекты 'grouping' и 'report' и который так же нужно удалить
					listRefs       []datamodels.IdentifierTypeSTIX // список ID объектов который 'grouping' объединяет в группу и который нужно перенести
					// в объект 'report' к которому принадлежит 'grouping' отмеченный для удаления
				}{
					targetRefsID:   targetID,
					relationshipID: obj.ID,
					listRefs:       sl[string(obj.SourceRef)].listRefs,
				}
			}
		}

		//получаем список ID STIX объектов типа 'report', на которые ссылаются найденные объекты 'relationship'
		cur, err = qp.Find((bson.D{
			bson.E{Key: "commonpropertiesobjectstix.type", Value: "report"},
			bson.E{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: listIDReporModify}}},
		}))
		if err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}

		//обрабатываем полученый список объектов типа 'report' и модифицируем их изменяя свойство ObjectRefs
		for _, v := range GetListElementSTIXObject(cur) {
			obj, ok := v.Data.(datamodels.ReportDomainObjectsSTIX)
			if !ok {
				continue
			}

			//ищем ID удаляемого объекта 'grouping' и удаляем из свойства ObjectRefs и в это же свойство добаляем все ссылки которые раньше
			// были в ObjectRefs удаляемого объекта 'grouping'
			for groupingID, v := range sl {
				if v.targetRefsID != obj.ID {
					continue
				}

				//удаляем ID объекта 'grouping' из свойства ObjectRefs объекта 'report' и добавляем туда ссылки на ID объектов находящиеся
				// в свойстве ObjectRefs удаляемого объекта 'grouping'
				listTmp := []datamodels.IdentifierTypeSTIX{}
				for _, v := range obj.ObjectRefs {
					if string(v) == groupingID {
						continue
					}

					listTmp = append(listTmp, v)
				}

				obj.ObjectRefs = append(listTmp, v.listRefs...)
			}

			listObjModiy = append(listObjModiy, &datamodels.ElementSTIXObject{
				DataType: obj.Type,
				Data:     obj,
			})
		}

		//обновляем STIX объекты типа 'report'
		err = ReplacementElementsSTIXObject(qp, listObjModiy)
		if err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}

		//удаляем выбранные в списках объекты типа 'relationship' и 'grouping'
		if _, err := qp.DeleteManyData(bson.D{{
			Key:   "commonpropertiesobjectstix.id",
			Value: bson.D{{Key: "$in", Value: append(listIDGroupingDel, listIDRelationshipDel...)}}}}); err != nil {
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}
	}
}

//wrapperFuncTypeHandlingSearchRequests набор обработчиков для работы с запросами, связанными с поиском информации
func (ws *wrappersSetting) wrapperFuncTypeHandlingSearchRequests(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) {

	var (
		err error
		fn  = commonlibs.GetFuncName()
		qp  = QueryParameters{
			NameDB:         ws.NameDB,
			CollectionName: "stix_object_collection",
			ConnectDB:      ws.ConnectionDB.Connection,
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
		if fn, err := searchSTIXObject(ws.DataRequest.AppTaskID, qp, psr, tst); err != nil {
			errorMessage.ErrorMessage.FuncName = fn
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}

	case "stix object list type grouping":
		if fn, err := searchSTIXObjectListTypeGrouping(ws.DataRequest.AppTaskID, qp, psr, tst); err != nil {
			errorMessage.ErrorMessage.FuncName = fn
			errorMessage.ErrorMessage.Error = err
			chanOutput <- errorMessage

			return
		}

	default:
		errorMessage.CommanDataTypePassedThroughChannels.ErrorMessage.Error = fmt.Errorf("the name of the database collection is not defined")
		chanOutput <- errorMessage

		return
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

			/*_, err := qp.InsertData([]interface{}{listDifferentObject})
			if err != nil {
				errorMessage.Error = err
				switchMSGType(&message, errorMessage)
				chanOutput <- message

				return
			}*/
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

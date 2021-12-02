package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/moduleapirequestprocessing/auxiliaryfunctions"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/moduledatabaseinteraction/interactionmongodb"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//HandlerAssignmentsModuleDataBaseInteraction - обработчик ответов от модуля хранилища данных
// (описывает общие для всех типов обрабатываемх запросов действия)
func HandlerAssignmentsModuleDataBaseInteraction(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clim *moddatamodels.ChannelsListInteractingModules) {

	var (
		err      error
		ti       *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType
		section  string = "не определена"
		taskType string = "не определен"
		taskID   string
	)

	if data.Section == "handling stix object" {
		section = "обработка структурированных данных"
		taskType = "добавление или обновление структурированных данных"
	} else if data.Section == "handling search requests" {
		section = "обработка поискового запроса"
		taskType = "осуществление поиска информации"
	} else if data.Section == "handling reference book" {
		section = "обработка справочной информации"
		taskType = "добавление или обновление справочной информации"
	} else if data.Section == "handling statistical requests" {
		section = "обработка статистического запроса"
		taskType = "статистика по структурированным данным"
	}

	fmt.Printf("func 'HandlerAssignmentsModuleDataBaseInteraction', section: '%s' - received appTaskID: '%s'\n", data.Section, data.AppTaskID)

	//получаем всю информацию о задаче
	taskID, ti, err = tst.GetTaskByID(data.AppTaskID)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			FuncName:    data.ErrorMessage.FuncName,
			Description: fmt.Sprint(data.ErrorMessage.Error),
		}
	}

	//обработка ошибок получаемых от БД MongoDB
	if data.ErrorMessage.Error != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			FuncName:    data.ErrorMessage.FuncName,
			Description: fmt.Sprint(data.ErrorMessage.Error),
		}

		fmt.Println("func 'HandlerAssignmentsModuleDataBaseInteraction', received Error Message")
		fmt.Println(data.ErrorMessage)
		fmt.Printf("func 'HandlerAssignmentsModuleDataBaseInteraction', data.AppTaskID: '%v'\n", data.AppTaskID)

		//удаляем задачу и результаты поиска информации, если они есть
		tst.DeletingTaskByID(data.AppTaskID)
		tst.DeletingFoundInformationByID(data.AppTaskID)

		//В случае если сообщение об ошибки не надо отправлять клиенту модуля 'moduleapirequestprocessing' просто завершаем обработчик
		if !data.ErrorMessage.ModuleAPIRequestProcessingSettingSendTo {
			return
		}

		if taskID == "" {
			return
		}

		if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
			ClientID:         "",
			TaskID:           taskID,
			Section:          data.Section,
			TypeNotification: "danger",
			Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
				Section:     section,
				TaskType:    taskType,
				FinalResult: "задача отклонена",
				Message:     "при выполнении задачи возникла ошибка базы данных",
			}),
			C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
		}); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				FuncName:    data.ErrorMessage.FuncName,
				Description: fmt.Sprint(data.ErrorMessage.Error),
			}
		}

		return
	}

	//обработка информационных сообщений получаемых от БД MongoDB
	if data.InformationMessage.Type != "" {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "info",
			Description: data.InformationMessage.Message,
		}

		if taskID == "" {
			return
		}

		if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
			ClientID:         ti.ClientID,
			TaskID:           taskID,
			Section:          data.Section,
			TypeNotification: data.InformationMessage.Type,
			Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
				Section:     section,
				TaskType:    taskType,
				FinalResult: "задача выполнена",
				Message:     data.InformationMessage.Message,
			}),
			C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
		}); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				FuncName:    "HandlerAssigmentsModuleDataBaseInteraction",
				Description: fmt.Sprint(data.ErrorMessage.Error),
			}
		}

		return
	}

	//непосредственно обработка ответа хранилища с учетом типа объектов
	if err := handlerDataBaseResponse(clim.ChannelsModuleAPIRequestProcessing.InputModule, data, tst, ti); err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			FuncName:    data.ErrorMessage.FuncName,
			Description: fmt.Sprint(err),
		}

		if err = auxiliaryfunctions.SendNotificationModuleAPI(&auxiliaryfunctions.SendNotificationTypeModuleAPI{
			ClientID:         ti.ClientID,
			TaskID:           taskID,
			Section:          data.Section,
			TypeNotification: "danger",
			Notification: commonlibs.PatternUserMessage(&commonlibs.PatternUserMessageType{
				Section:     section,
				TaskType:    taskType,
				FinalResult: "задача отклонена",
				Message:     "при выполнении задачи возникла ошибка",
			}),
			C: clim.ChannelsModuleAPIRequestProcessing.InputModule,
		}); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				FuncName:    data.ErrorMessage.FuncName,
				Description: fmt.Sprint(err),
			}
		}
	}
}

//handlerDataBaseResponse - непосредственно обработчик ответов на запросы к сущиностям различных типов находящихся в хранилище,
// приходят от модуля хранилища
func handlerDataBaseResponse(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {

	//размер части сообщения
	const _maxChunkSize int = 100
	if ti.TaskStatus == "completed" {
		switch data.Section {
		case "handling search requests": //обработка ответов на поисковый запрос
			//делаем запрос к временному хранилищу информации
			result, err := tst.GetFoundInformationByID(data.AppTaskID)
			if err != nil {
				return err
			}

			if result.Collection == "stix_object_collection" {
				if err := handlingSearchRequestsSTIXObject(chanResModAPI, _maxChunkSize, data, result, ti); err != nil {
					return err
				}
			}

			if result.Collection == "accounting_differences_objects_collection" {
				if err := handlingSearchDifferencesObjectsCollection(chanResModAPI, _maxChunkSize, data, result, ti); err != nil {
					return err
				}
			}

		case "handling statistical requests": //обработка ответов на запрос статистической информации
			if err := handlingStatisticalRequestsSTIXObject(chanResModAPI, data, tst, ti); err != nil {
				return err
			}

		case "handling reference book": //обработка ответов на операции со стправочниками
			if err := handlingRBRequests(chanResModAPI, _maxChunkSize, data, tst, ti); err != nil {
				return err
			}

		}
	}

	//удаляем задачу и результаты поиска информации, если они есть
	tst.DeletingTaskByID(data.AppTaskID)
	tst.DeletingFoundInformationByID(data.AppTaskID)

	return nil
}

//handlingRBRequest - обработчик запросов на операции с объектами справочниками
func handlingRBRequests(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	maxChunkSize int,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {

	result, err := tst.GetFoundInformationByID(data.AppTaskID)
	if err != nil {
		return err
	}
	msgRes := datamodels.ModAPIRequestProcessingResJSON{
		ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
			TaskID:  data.AppTaskID,
			Section: data.Section,
		},
		IsSuccessful: true,
	}

	//Непосредственное действие по отправке - используем замыкание для доступа в область видимости на уровень выше
	ActionRBObjs := func(obj interface{}, chunkTotalNumbers, chunkNumber, chunkSize int) error {
		//кастомизируем к нужному типу
		RBObjects, ok := obj.([]*datamodels.RBookRespParameter)
		if !ok {
			return fmt.Errorf("type conversion error, line 206")
		}
		//формируем ответ
		msgRes.AdditionalParameters = datamodels.ResJSONParts{
			TotalNumberParts:      chunkTotalNumbers,
			GivenSizePart:         chunkSize,
			NumberTransmittedPart: chunkNumber,
			TransmittedData:       RBObjects,
		}
		//преобразуем результат в байты
		msg, err := json.Marshal(msgRes)
		if err != nil {
			return err
		}
		//отправляем результат в канал
		chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module core application",
				ModuleReceiverMessage:  "module api request processing",
			},
			ClientID: ti.ClientID,
			DataType: 1,
			Data:     &msg,
		}
		return nil
	}

	//обрабатываем полученный список RB-объектов, в том числе если он превышает размер в maxChunkSize
	//разбиваем его на части и отправляем в канал для дальнейшей обработки
	commonlibs.СhunkSplitting(result.Information, ActionRBObjs, maxChunkSize)

	return nil
}

//handlingSearchRequestsSTIXObject - обработчик ответов на запросы поиска по STIX объектам
func handlingSearchRequestsSTIXObject(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	maxChunkSize int,
	data datamodels.ModuleDataBaseInteractionChannel,
	result *memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {

	msgRes := datamodels.ModAPIRequestProcessingResJSON{
		ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
			TaskID:  ti.ClientTaskID,
			Section: data.Section,
		},
		IsSuccessful: true,
	}

	switch result.ResultType {
	case "only_count":
		//для КРАТКОЙ информации, только количество, по найденным STIX объектам

		numFound, ok := result.Information.(int64)
		if !ok {
			return fmt.Errorf("type conversion error, line 298")
		}

		msgRes.AdditionalParameters = struct {
			NumberDocumentsFound int64 `json:"number_documents_found"`
		}{
			NumberDocumentsFound: numFound,
		}

	case "full_found_info":
		//для ПОЛНОЙ информации по найденным STIX объектам

		listElemSTIXObj, ok := result.Information.([]*datamodels.ElementSTIXObject)
		if !ok {
			return fmt.Errorf("type conversion error, line 316")
		}

		sestixo := len(listElemSTIXObj)
		listMsgRes := make([]interface{}, 0, sestixo)
		for _, v := range listElemSTIXObj {
			listMsgRes = append(listMsgRes, v.Data)
		}

		//обрабатываем полученный список STIX объектов, в том числе если он превышает размер в 100 объектов
		if sestixo < maxChunkSize {
			msgRes.AdditionalParameters = datamodels.ResJSONParts{
				TotalNumberParts:      1,
				GivenSizePart:         maxChunkSize,
				NumberTransmittedPart: 1,
				TransmittedData:       listMsgRes,
			}
		} else {
			num := commonlibs.GetCountChunk(int64(sestixo), maxChunkSize)
			min := 0
			max := maxChunkSize
			for i := 0; i < num; i++ {
				data := datamodels.ResJSONParts{
					TotalNumberParts:      num,
					GivenSizePart:         maxChunkSize,
					NumberTransmittedPart: i + 1,
				}

				if i == 0 {
					data.TransmittedData = listMsgRes[:max]
				} else if i == num-1 {
					data.TransmittedData = listMsgRes[min:]
				} else {
					data.TransmittedData = listMsgRes[min:max]
				}

				min = min + maxChunkSize
				max = max + maxChunkSize
				msgRes.AdditionalParameters = data
			}
		}
	}

	msg, err := json.Marshal(msgRes)
	if err != nil {
		return err
	}

	chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module api request processing",
		},
		ClientID: ti.ClientID,
		DataType: 1,
		Data:     &msg,
	}

	return nil
}

//handlingSearchDifferencesObjectsCollection - обработчик ответов на запросы поиска по коллекции отслеживающей изменение объектов
func handlingSearchDifferencesObjectsCollection(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	maxChunkSize int,
	data datamodels.ModuleDataBaseInteractionChannel,
	result *memorytemporarystoragecommoninformation.TemporaryStorageFoundInformation,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {

	fmt.Println("func 'handlingSearchDifferencesObjectsCollection', START...")

	msgRes := datamodels.ModAPIRequestProcessingResJSON{
		ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
			TaskID:  ti.ClientTaskID,
			Section: data.Section,
		},
		IsSuccessful: true,
	}

	listDifferentObject, ok := result.Information.([]datamodels.DifferentObjectType)
	if !ok {
		return fmt.Errorf("type conversion error, line 398")
	}

	numldo := len(listDifferentObject)
	if numldo < maxChunkSize {
		msgRes.AdditionalParameters = datamodels.ResJSONParts{
			TotalNumberParts:      1,
			GivenSizePart:         maxChunkSize,
			NumberTransmittedPart: 1,
			TransmittedData:       listDifferentObject,
		}
	} else {
		num := commonlibs.GetCountChunk(int64(numldo), maxChunkSize)
		min := 0
		max := maxChunkSize
		for i := 0; i < num; i++ {
			data := datamodels.ResJSONParts{
				TotalNumberParts:      num,
				GivenSizePart:         maxChunkSize,
				NumberTransmittedPart: i + 1,
			}

			if i == 0 {
				data.TransmittedData = listDifferentObject[:max]
			} else if i == num-1 {
				data.TransmittedData = listDifferentObject[min:]
			} else {
				data.TransmittedData = listDifferentObject[min:max]
			}

			min = min + maxChunkSize
			max = max + maxChunkSize
			msgRes.AdditionalParameters = data
		}

		/*
			Тут надо посмотреть внимательнее, отправки похоже нет!!!
		*/
	}

	fmt.Printf("func 'handlingSearchDifferencesObjectsCollection', MsgRes: '%v'\n", msgRes)

	/*
	   по тестам msgRes содержит всю необходимую информацию
	   надо дальше смотреть
	*/

	msg, err := json.Marshal(msgRes)
	if err != nil {
		return err
	}

	chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module api request processing",
		},
		ClientID: ti.ClientID,
		DataType: 1,
		Data:     &msg,
	}

	return nil
}

//На всякий случай оставляю тот код который был в func handlingSearchRequestsSTIXObject
/*	listElemSTIXObj, ok := result.Information.([]*datamodels.ElementSTIXObject)
	if !ok {
		return fmt.Errorf("type conversion error, line 220")
	}

	sestixo := len(listElemSTIXObj)
	listMsgRes := make([]interface{}, 0, sestixo)
	for _, v := range listElemSTIXObj {
		listMsgRes = append(listMsgRes, v.Data)
	}*/

//обрабатываем полученный список STIX объектов, в том числе если он превышает размер в 100 объектов
/*	if sestixo < maxChunkSize {
	msgRes.AdditionalParameters = datamodels.ResJSONParts{
		TotalNumberParts:      1,
		GivenSizePart:         maxChunkSize,
		NumberTransmittedPart: 1,
		TransmittedData:       listMsgRes,
	}
*/
//преобразуем результат в байты
/*			msg, err := json.Marshal(msgRes)
			if err != nil {
				return err
			}

			chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
				CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
					ModuleGeneratorMessage: "module core application",
					ModuleReceiverMessage:  "module api request processing",
				},
				ClientID: ti.ClientID,
				DataType: 1,
				Data:     &msg,
			}
		} else {
			num := commonlibs.GetCountChunk(int64(sestixo), maxChunkSize)

			min := 0
			max := maxChunkSize
			for i := 0; i < num; i++ {
				data := datamodels.ResJSONParts{
					TotalNumberParts:      num,
					GivenSizePart:         maxChunkSize,
					NumberTransmittedPart: i + 1,
				}

				if i == 0 {
					data.TransmittedData = listMsgRes[:max]
				} else if i == num-1 {
					data.TransmittedData = listMsgRes[min:]
				} else {
					data.TransmittedData = listMsgRes[min:max]
				}

				min = min + maxChunkSize
				max = max + maxChunkSize

				msgRes.AdditionalParameters = data
				msg, err := json.Marshal(msgRes)
				if err != nil {
					return err
				}

				chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
					CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
						ModuleGeneratorMessage: "module core application",
						ModuleReceiverMessage:  "module api request processing",
					},
					ClientID: ti.ClientID,
					DataType: 1,
					Data:     &msg,
				}
			}
		}

	}
}*/

//return nil
//}

//handlingStatisticalRequestsSTIXObject - обработчик ответов на запросы статистичеких данных
func handlingStatisticalRequestsSTIXObject(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	data datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {

	//делаем запрос к временному хранилищу информации
	result, err := tst.GetFoundInformationByID(data.AppTaskID)
	if err != nil {
		return err
	}

	msgRes := datamodels.ModAPIRequestProcessingResJSON{
		ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
			TaskID:  ti.ClientTaskID,
			Section: data.Section,
		},
		IsSuccessful: true,
	}

	if result.ResultType == "handling_statistical_requests" {
		info, ok := result.Information.(interactionmongodb.ResultStatisticalInformationSTIXObject)
		if !ok {
			return fmt.Errorf("type conversion error, line 323")
		}

		msgRes.AdditionalParameters = info
	}

	msg, err := json.Marshal(msgRes)
	if err != nil {
		return err
	}

	chanResModAPI <- datamodels.ModuleReguestProcessingChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module api request processing",
		},
		ClientID: ti.ClientID,
		DataType: 1,
		Data:     &msg,
	}

	return nil
}

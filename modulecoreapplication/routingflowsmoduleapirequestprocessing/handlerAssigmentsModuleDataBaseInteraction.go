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

func HandlerAssigmentsModuleDataBaseInteraction(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data *datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clim *moddatamodels.ChannelsListInteractingModules) {

	var (
		err                       error
		ti                        *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType
		section, taskType, taskID string
	)

	if data.Section == "handling stix object" {
		section = "обработка структурированных данных"
		taskType = "добавление или обновление структурированных данных"
	} else if data.Section == "handling search requests" {
		section = "обработка поискового запроса"
		taskType = "осуществление поиска информации"
	} else {
		section = "не определена"
		taskType = "не определен"
	}

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

		//удаляем задачу и результаты поиска информации, если они есть
		tst.DeletingTaskByID(data.AppTaskID)
		tst.DeletingFoundInformationByID(data.AppTaskID)

		//если сообщение об ошибки не надо отправлять клиенту модуля 'moduleapirequestprocessing'
		if !data.ErrorMessage.ModuleAPIRequestProcessingSettingSendTo {
			return
		}

		if taskID == "" {
			return
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

	if err := handlerDataBaseResponse(clim.ChannelsModuleAPIRequestProcessing.OutputModule, data, tst, ti); err != nil {
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

func handlerDataBaseResponse(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	data *datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {

	//размер части сообщения
	const _maxChunkSize int = 100

	switch data.Section {
	case "handling search requests":
		if err := handlingSearchRequestsSTIXObject(chanResModAPI, _maxChunkSize, data, tst, ti); err != nil {
			return err
		}

	case "":
		//пока заглушка, будет использоватся для обработки ответов от БД не связанных с секцией 'handling search requests'
	}

	//удаляем задачу и результаты поиска информации, если они есть
	tst.DeletingTaskByID(data.AppTaskID)
	tst.DeletingFoundInformationByID(data.AppTaskID)

	return nil
}

func handlingSearchRequestsSTIXObject(
	chanResModAPI chan<- datamodels.ModuleReguestProcessingChannel,
	maxChunkSize int,
	data *datamodels.ModuleDataBaseInteractionChannel,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	ti *memorytemporarystoragecommoninformation.TemporaryStorageTaskInDetailType) error {
	if ti.TaskStatus != "completed" {
		return nil
	}

	tp, ok := ti.TaskParameters.(datamodels.ModAPIRequestProcessingResJSONSearchReqType)
	if !ok {
		return fmt.Errorf("type conversion error, line 166")
	}

	//обрабатываем результаты опираясь на типы коллекций
	if tp.CollectionName == "stix object" {
		//делаем запрос к временному хранилищу информации
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

		//для КРАТКОЙ информации, только колличество, по найденным STIX объектам
		if result.Collection == "stix_object_collection" && result.ResultType == "only_count" {
			numFound, ok := result.Information.(int64)
			if !ok {
				return fmt.Errorf("type conversion error, line 191")
			}

			msgRes.AdditionalParameters = struct {
				NumberDocumentsFound int64 `json:"number_documents_found"`
			}{
				NumberDocumentsFound: numFound,
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
		}

		//для ПОЛНОЙ информации по найденным STIX объектам
		if result.Collection == "stix_object_collection" && result.ResultType == "full_found_info" {
			listElemSTIXObj, ok := result.Information.([]*datamodels.ElementSTIXObject)
			if !ok {
				return fmt.Errorf("type conversion error, line 220")
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
	}

	return nil
}

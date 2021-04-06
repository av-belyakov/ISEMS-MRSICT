package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"

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

	commonMsgReq, err := unmarshalJSONCommonReq(data.Data)
	if err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "unmarshalJSONCommonReq",
		}

		if err := auxiliaryfunctions.SendCriticalErrorMessageJSON(&auxiliaryfunctions.ErrorMessageType{
			ClientID: data.ClientID,
			Error:    fmt.Errorf("Error: error when decoding a JSON document"),
			C:        clim.ChannelsModuleAPIRequestProcessing.InputModule,
		}); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "unmarshalJSONCommonReq",
			}

			return
		}

		return
	}

	switch commonMsgReq.Section {
	case "handling stix object":

		/* *** обработчик JSON сообщений со STIX объектами *** */

		l, err := UnmarshalJSONObjectSTIXReq(*commonMsgReq)
		//если полностью не возможно декодировать список STIX объектов
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "unmarshalJSONObjectSTIXReq",
			}

			if err := auxiliaryfunctions.SendCriticalErrorMessageJSON(&auxiliaryfunctions.ErrorMessageType{
				ClientID: data.ClientID,
				Error:    fmt.Errorf("Error: error when decoding a JSON document. Section: '%v'", commonMsgReq.Section),
				C:        clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "unmarshalJSONObjectSTIXReq",
				}
			}

			return
		}

		/*
							!!!!!!!
			   Функция CheckSTIXObjects(l) пока всего лишь заглушка, она пустая внутри и ничего не делает.
			   Необходимо ее дописать для валидации STIX объектов
							!!!!!!!
		*/

		//выполняем валидацию полученных STIX объектов
		if err := CheckSTIXObjects(l); err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "unmarshalJSONObjectSTIXReq",
			}

			if err := auxiliaryfunctions.SendCriticalErrorMessageJSON(&auxiliaryfunctions.ErrorMessageType{
				ClientID: data.ClientID,
				Error:    fmt.Errorf("Error: non-valid STIX objects were received. Section: '%v'", commonMsgReq.Section),
				C:        clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "unmarshalJSONObjectSTIXReq",
				}
			}

			return
		}

		//добавляем информацию о задаче в хранилище задач
		appTaskID, err := tst.AddNewTask(&memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
			TaskGenerator:        data.ModuleGeneratorMessage,
			ClientID:             data.ClientID,
			ClientName:           data.ClientName,
			ClientTaskID:         commonMsgReq.TaskID,
			AdditionalClientName: commonMsgReq.UserNameGeneratedTask,
			Section:              commonMsgReq.Section,
			Command:              "", //в случае с объектами STIX команда не указывается (автоматически подразумевается добавление или обновление объектов STIX)
			TaskParameters:       l,
		})
		if err != nil {
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "unmarshalJSONObjectSTIXReq",
			}
		}
		fmt.Println(l)
		fmt.Printf("Application task ID: '%s'\n", appTaskID)

	case "handling search requests":

		/* *** обработчик JSON сообщений с запросами к поисковой машине приложения *** */

	case "handling reference book":

		/* *** обработчик JSON сообщений с параметрами связанными со справочниками *** */

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

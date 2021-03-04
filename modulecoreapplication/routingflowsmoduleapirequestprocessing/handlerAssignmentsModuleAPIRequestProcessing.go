package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/moduleapirequestprocessing/auxiliaryfunctions"
	moddatamodels "ISEMS-MRSICT/modulecoreapplication/datamodels"
	"ISEMS-MRSICT/modulelogginginformationerrors"
)

//HandlerAssigmentsModuleAPIRequestProcessing является обработчиком JSON  приходящих
func HandlerAssigmentsModuleAPIRequestProcessing(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	data *datamodels.ModuleReguestProcessingChannel,
	clim *moddatamodels.ChannelsListInteractingModules) {

	commonMsgReq, err := unmarshalJSONCommonReq(data.Data)
	if err != nil {
		//запись информации в лог-файл
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "unmarshalJSONCommonReq",
		}

		em, e := auxiliaryfunctions.CreateCriticalErrorMessageJSON(&auxiliaryfunctions.CriticalErrorMessageType{Error: err})
		if e != nil {
			//запись информации в лог-файл
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "unmarshalJSONCommonReq",
			}

			return
		}

		//здесь отправляем информационное сообщение через канал клиенту API
		clim.ChannelsModuleAPIRequestProcessing.InputModule <- datamodels.ModuleReguestProcessingChannel{
			CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
				ModuleGeneratorMessage: "module core application",
				ModuleReceiverMessage:  "module api request processing",
			},
			ClientID: data.ClientID,
			DataType: 1,
			Data:     em,
		}

		return
	}

	switch commonMsgReq.Section {
	case "handling stix object":
		l, ok, err := unmarshalJSONObjectSTIXReq(*commonMsgReq)
		//если полностью не возможно декодировать список STIX объектов
		if err != nil {
			//запись информации в лог-файл
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: fmt.Sprint(err),
				FuncName:    "unmarshalJSONCommonReq",
			}

			em, e := auxiliaryfunctions.CreateCriticalErrorMessageJSON(&auxiliaryfunctions.CriticalErrorMessageType{
				TaskID:  commonMsgReq.TaskID,
				Section: commonMsgReq.Section,
				Error:   err,
			})
			if e != nil {
				//запись информации в лог-файл
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "unmarshalJSONObjectSTIXReq",
				}

				return
			}

			//здесь отправляем информационное сообщение через канал клиенту API
			clim.ChannelsModuleAPIRequestProcessing.InputModule <- datamodels.ModuleReguestProcessingChannel{
				CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
					ModuleGeneratorMessage: "module core application",
					ModuleReceiverMessage:  "module api request processing",
				},
				ClientID: data.ClientID,
				DataType: 1,
				Data:     em,
			}

			return
		}

		if !ok {
			//не все stix объекты были успешно декодированны
			//отправляем сообщение с флагом PartiallySuccessful true
		}

		//выполняем дальнейшую обработку
		//зоздаем внутренний task ID приложения
		//добавляем информацию о задаче в хранилище задач
		fmt.Println(l)

	case "handling search requests":

	case "":

	}
}

//unmarshalJSONCommonReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', только по общим полям
func unmarshalJSONCommonReq(msgReq *[]byte) (*datamodels.ModAPIRequestProcessingReqJSON, error) {
	var modAPIRequestProcessingReqJSON datamodels.ModAPIRequestProcessingReqJSON
	err := json.Unmarshal(*msgReq, &modAPIRequestProcessingReqJSON)

	return &modAPIRequestProcessingReqJSON, err
}

//unmarshalJSONObjectSTIXReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который содержит список объектов STIX
func unmarshalJSONObjectSTIXReq(msgReq datamodels.ModAPIRequestProcessingReqJSON) ([]*datamodels.ListSTIXObject, bool, error) {
	var (
		isFail                             bool
		listResults                        []*datamodels.ListSTIXObject
		listSTIXObjectJSON                 datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
		commonPropertiesObjectSTIX         datamodels.CommonPropertiesObjectSTIX
		numberSuccessfullyProcessedObjects int
	)

	if err := json.Unmarshal(*msgReq.RequestDetails, &listSTIXObjectJSON); err != nil {
		return nil, isFail, err
	}

	for _, item := range listSTIXObjectJSON {
		err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
		if err != nil {
			numberSuccessfullyProcessedObjects++
		}

		r, t, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
		if err != nil {
			numberSuccessfullyProcessedObjects++
		}

		listResults = append(listResults, &datamodels.ListSTIXObject{
			DataType: t,
			Data:     r,
		})
	}

	if len(listSTIXObjectJSON) == numberSuccessfullyProcessedObjects {
		isFail = true
	}

	return listResults, isFail, nil
}

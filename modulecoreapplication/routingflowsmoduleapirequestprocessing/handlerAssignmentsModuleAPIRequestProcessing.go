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

//HandlerAssigmentsModuleAPIRequestProcessing является обработчиком JSON  приходящих
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
		l, success, err := unmarshalJSONObjectSTIXReq(*commonMsgReq)
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

				return
			}

			return
		}

		if success.Unsuccess != 0 {
			errStr := fmt.Sprintf("Error: error when decoding a JSON document, only %d out of %d elements of the document were successfully decoded. Section: '%s'", (success.All - success.Unsuccess), success.All, commonMsgReq.Section)
			chanSaveLog <- modulelogginginformationerrors.LogMessageType{
				TypeMessage: "error",
				Description: errStr,
				FuncName:    "unmarshalJSONObjectSTIXReq",
			}

			if err := auxiliaryfunctions.SendCriticalErrorMessageJSON(&auxiliaryfunctions.ErrorMessageType{
				ClientID: data.ClientID,
				Error:    fmt.Errorf(errStr),
				C:        clim.ChannelsModuleAPIRequestProcessing.InputModule,
			}); err != nil {
				chanSaveLog <- modulelogginginformationerrors.LogMessageType{
					TypeMessage: "error",
					Description: fmt.Sprint(err),
					FuncName:    "unmarshalJSONObjectSTIXReq",
				}

				return
			}

			return
		}

		fmt.Println(success.All)
		fmt.Println(success.Unsuccess)

		//выполняем дальнейшую обработку
		//зоздаем внутренний task ID приложения
		//добавляем информацию о задаче в хранилище задач
		appTaskID, err := tst.AddNewTask(&memorytemporarystoragecommoninformation.TemporaryStorageTaskType{
			TaskGenerator:        data.ModuleGeneratorMessage,
			ClientID:             data.ClientID,
			ClientName:           data.ClientName,
			ClientTaskID:         commonMsgReq.TaskID,
			AdditionalClientName: commonMsgReq.UserNameGeneratedTask,
			Section:              commonMsgReq.Section,
			Command:              "", //в случае с объектами STIX команда не указывается (автоматически подразумевается добавление или обновление объектов STIX)
			TaskParameters:       commonMsgReq.RequestDetails,
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
func unmarshalJSONObjectSTIXReq(msgReq datamodels.ModAPIRequestProcessingReqJSON) ([]*datamodels.ListSTIXObject, struct{ All, Unsuccess int }, error) {
	var (
		listResults                []*datamodels.ListSTIXObject
		listSTIXObjectJSON         datamodels.ModAPIRequestProcessingReqHandlingSTIXObjectJSON
		commonPropertiesObjectSTIX datamodels.CommonPropertiesObjectSTIX
		numberUnsuccessfullyProc   int
	)

	if err := json.Unmarshal(*msgReq.RequestDetails, &listSTIXObjectJSON); err != nil {
		return nil, struct{ All, Unsuccess int }{0, 0}, err
	}

	for _, item := range listSTIXObjectJSON {
		err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
		if err != nil {
			numberUnsuccessfullyProc++
		}

		r, t, err := commonlibs.DecoderFromJSONToSTIXObject(commonPropertiesObjectSTIX.Type, item)
		if err != nil {
			numberUnsuccessfullyProc++
		}

		listResults = append(listResults, &datamodels.ListSTIXObject{
			DataType: t,
			Data:     r,
		})
	}

	return listResults, struct{ All, Unsuccess int }{len(listSTIXObjectJSON), numberUnsuccessfullyProc}, nil
}

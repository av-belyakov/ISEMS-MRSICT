package auxiliaryfunctions

import (
	"encoding/json"

	"ISEMS-MRSICT/datamodels"
)

//SendNotificationTypeModuleAPI  тип содержащий описание информационного сообщения
// ClientID - идентификатор клиента moduleapirequestprocessing
// TaskID - ID задачи
// Section - секция, в которой возникла ошибка
// Notification - подробное информационное сообщение
type SendNotificationTypeModuleAPI struct {
	ClientID     string
	TaskID       string
	Section      string
	Notification string
	C            chan<- datamodels.ModuleReguestProcessingChannel
}

//SendNotificationModuleAPI формирует информационное сообщение в формате JSON
func SendNotificationModuleAPI(n *SendNotificationTypeModuleAPI) error {
	msg, err := json.Marshal(datamodels.ModAPIRequestProcessingResJSON{Description: n.Notification})
	if err != nil {
		return err
	}

	/*
	   Переделать все сообщения отправляемые клиенту API под это сообщение и с учетом\
	   type ModAPIRequestProcessingResJSON struct {
	   	ModAPIRequestProcessingCommonJSON
	   	IsSuccessful         bool                                      `json:"is_successful"`
	   	Description          string                                    `json:"description"`
	   	InformationMessage   ModAPIRequestProcessingResJSONInfoMsgType `json:"information_message"`
	   	AdditionalParameters interface{}                               `json:"additional_parameters"`
	   }
	*/

	n.C <- datamodels.ModuleReguestProcessingChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module api request processing",
		},
		ClientID: n.ClientID,
		DataType: 1,
		Data:     &msg,
	}

	return nil
}

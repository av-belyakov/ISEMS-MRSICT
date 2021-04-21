package auxiliaryfunctions

import (
	"encoding/json"

	"ISEMS-MRSICT/datamodels"
)

//SendNotificationTypeModuleAPI  тип содержащий описание информационного сообщения
// ClientID - идентификатор клиента moduleapirequestprocessing
// TaskID - ID задачи
// Section - секция, в которой возникла ошибка
// TypeNotification - тип информационного сообщения
// Notification - подробное информационное сообщение
type SendNotificationTypeModuleAPI struct {
	ClientID         string
	TaskID           string
	Section          string
	TypeNotification string
	Notification     string
	C                chan<- datamodels.ModuleReguestProcessingChannel
}

//SendNotificationModuleAPI формирует информационное сообщение в формате JSON
func SendNotificationModuleAPI(n *SendNotificationTypeModuleAPI) error {
	var isSuccess bool

	if n.TypeNotification == "success" || n.TypeNotification == "info" {
		isSuccess = true
	}

	msg, err := json.Marshal(datamodels.ModAPIRequestProcessingResJSON{
		ModAPIRequestProcessingCommonJSON: datamodels.ModAPIRequestProcessingCommonJSON{
			TaskID:  n.TaskID,
			Section: n.Section,
		},
		IsSuccessful: isSuccess,
		InformationMessage: datamodels.ModAPIRequestProcessingResJSONInfoMsgType{
			MsgType: n.TypeNotification,
			Msg:     n.Notification,
		}})
	if err != nil {
		return err
	}

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

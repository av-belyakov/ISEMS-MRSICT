package auxiliaryfunctions

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/datamodels"
)

//ErrorMessageType  тип содержащий описание сообщения, информирующего о возникновении критической ошибки
// ClientID - идентификатор клиента moduleapirequestprocessing
// TaskID - ID задачи
// Section - секция, в которой возникла ошибка
// Error - ошибка
type ErrorMessageType struct {
	ClientID string
	TaskID   string
	Section  string
	Error    error
	C        chan<- datamodels.ModuleReguestProcessingChannel
}

//SendCriticalErrorMessageJSON формирует информационное сообщение в формате JSON о возникшей критической ошибке
func SendCriticalErrorMessageJSON(cem *ErrorMessageType) error {
	msg, err := json.Marshal(datamodels.ModAPIRequestProcessingResJSON{Description: fmt.Sprint(cem.Error)})
	if err != nil {
		return err
	}

	cem.C <- datamodels.ModuleReguestProcessingChannel{
		CommanDataTypePassedThroughChannels: datamodels.CommanDataTypePassedThroughChannels{
			ModuleGeneratorMessage: "module core application",
			ModuleReceiverMessage:  "module api request processing",
		},
		ClientID: cem.ClientID,
		DataType: 1,
		Data:     &msg,
	}

	return nil
}

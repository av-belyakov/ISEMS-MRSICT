package auxiliaryfunctions

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/datamodels"
)

//CriticalErrorMessageType  тип содержащий описание сообщения, информирующего о возникновении критической ошибки
// TaskID - ID задачи
// Section - секция, в которой возникла ошибка
// Error - ошибка
type CriticalErrorMessageType struct {
	TaskID  string
	Section string
	Error   error
}

//CreateCriticalErrorMessageJSON формирует информационное сообщение в формате JSON о возникшей критической ошибке
func CreateCriticalErrorMessageJSON(cem *CriticalErrorMessageType) (*[]byte, error) {
	msg, err := json.Marshal(datamodels.ModAPIRequestProcessingResJSON{
		MessageType: datamodels.ModAPIRequestProcessingResJSONMessageType{
			CriticalError:       true,
			TaskExecutionResult: false,
			DataTransfer:        false,
		},
		DetailedInformation: datamodels.ModAPIRequestProcessingResJSONCriticalError{
			Description: fmt.Sprint(cem.Error),
		},
	})

	return &msg, err
}

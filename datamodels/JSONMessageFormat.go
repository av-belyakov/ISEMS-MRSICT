package datamodels

import "encoding/json"

//ModAPIRequestProcessingCommonJSON содержит описание общих полей формата JSON для запросов и ответов модуля ModuleAPIRequestProcessing
// TaskID - ID задачи
// Section - секция, предназначена для разграничения запросов по типам обработчиков. Например, используются следующие типы обработчиков:
//  - "handling stix object" (обработка объектов STIX)
//  - "handling search requests" (обработка поисковых запросов)
//  - "handling reference book" (обработка запросов связанных со справочниками)
//  - "generating reports" (генерирование отчетов)
//  - "formation final documents" (генерирование итоговых документов)
type ModAPIRequestProcessingCommonJSON struct {
	TaskID  string `json:"task_id"`
	Section string `json:"section"`
}

//ModAPIRequestProcessingReqJSON содержит описание формата JSON запроса получаемого через модуль ModuleAPIRequestProcessing
// TaskWasGeneratedAutomatically - задача была сгенерирована автоматически (true - да)
// UserNameGeneratedTask - имя пользователя сгенерировавшего задачу
// RequestDetails - подробности запроса
type ModAPIRequestProcessingReqJSON struct {
	ModAPIRequestProcessingCommonJSON
	TaskWasGeneratedAutomatically bool             `json:"task_was_generated_automatically"`
	UserNameGeneratedTask         string           `json:"user_name_generated_task"`
	RequestDetails                *json.RawMessage `json:"request_details"`
}

//ModAPIRequestProcessingReqHandlingSTIXObjectJSON содержит список произвольных объектов STIX. Информация из данных списков добавляется в
// базу данных, если ее там нет или обновляется, если она там есть. Данный тип применяется ТОЛЬКО при Section:"handling stix object"
type ModAPIRequestProcessingReqHandlingSTIXObjectJSON []*json.RawMessage

//ModAPIRequestProcessingResJSON содержит описание формата JSON ответа, передаваемого через модуль ModuleAPIRequestProcessing
// Description - дополнительное детальное описание. При возникновении ошибки, сюда пишется максимально подробное описание ошибки.
// IsSuccessful - индикатор успешности выполнения задачи (false - неуспешно, true - успешно).
// InformationMessage - информационное сообщение
// AdditionalParameters - дополнительные параметры связанные с выполняемой задачей.
type ModAPIRequestProcessingResJSON struct {
	ModAPIRequestProcessingCommonJSON
	IsSuccessful         bool                                      `json:"is_successful"`
	Description          string                                    `json:"description"`
	InformationMessage   ModAPIRequestProcessingResJSONInfoMsgType `json:"information_message"`
	AdditionalParameters interface{}                               `json:"additional_parameters"`
}

//ModAPIRequestProcessingResJSONInfoMsgType подробное описание InformationMessage
// MsgType - тип информационного сообщения ('info', 'success', 'warning', 'danger')
// Msg - информационное сообщение
type ModAPIRequestProcessingResJSONInfoMsgType struct {
	MsgType string `json:"msg_type"`
	Msg     string `json:"msg"`
}

/*
	Example!!! для передачи информации частями
//TransmittingRequestedInformationModAPIRequestProcessingResJSON содержит информацию передаваемую в ответ на запрашиваемые данные
// TotalNumberParts - общее количество частей
// GivenSizePart - заданный размер части
// NumberTransmittedPart - номер передаваемой части
// TransmittedData - передаваемые данные
type TransmittingRequestedInformationModAPIRequestProcessingResJSON struct {
	TotalNumberParts      int         `json:"total_number_parts"`
	GivenSizePart         int         `json:"given_size_part"`
	NumberTransmittedPart int         `json:"number_transmitted_part"`
	TransmittedData       interface{} `json:"transmitted_data"` // пока 'interface{}' так как не знаю тип передаваемых данных
}
*/

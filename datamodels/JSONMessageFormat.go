package datamodels

import "encoding/json"

//ModAPIRequestProcessingCommonJSON содержит описание общих полей формата JSON для запросов и ответов модуля ModuleAPIRequestProcessing
// TaskID - ID задачи
type ModAPIRequestProcessingCommonJSON struct {
	TaskID string `json:"task_id"`
}

//ModAPIRequestProcessingReqJSON содержит описание формата JSON запроса получаемого через модуль ModuleAPIRequestProcessing
// Section - секция, предназначена для разграничения запросов по типам обработчиков. Например, используются следующие типы обработчиков:
//  - "handling stix object" (обработка объектов STIX)
//  - "handling search requests" (обработка поисковых запросов)
//  - "generating reports" (генерирование отчетов)
//  - "formation final documents" (генерирование итоговых документов)
// TaskWasGeneratedAutomatically - задача была сгенерирована автоматически (true - да)
// UserNameGeneratedTask - имя пользователя сгенерировавшего задачу
// RequestDetails - подробности запроса
// DetailedInformation - подбробная информация о запросе
type ModAPIRequestProcessingReqJSON struct {
	ModAPIRequestProcessingCommonJSON
	Section                       string           `json:"section"`
	TaskWasGeneratedAutomatically bool             `json:"task_was_generated_automatically"`
	UserNameGeneratedTask         string           `json:"user_name_generated_task"`
	RequestDetails                *json.RawMessage `json:"request_details"`
}

//ModAPIRequestProcessingReqHandlingSTIXObjectJSON содержит список произвольных объектов STIX. Информация из данных списков добавляется в
// базу данных, если ее там нет или обновляется, если она там есть. Данный тип применяется ТОЛЬКО при Section:"handling stix object"
type ModAPIRequestProcessingReqHandlingSTIXObjectJSON []*json.RawMessage

//ModAPIRequestProcessingResJSON содержит описание формата JSON ответа, передаваемого через модуль ModuleAPIRequestProcessing
// MessageType - тип сообщения
//  CriticalError - информационное сообщение о критической ошибке
//  TaskExecutionResult - информационное сообщение о результате выполнения задачи
//  DataTransfer - передача данных, например передача данных в результате успешного поиска информации
// DetailedInformation - подробная информация о запросе
type ModAPIRequestProcessingResJSON struct {
	MessageType         ModAPIRequestProcessingResJSONMessageType `json:"message_type"`
	DetailedInformation interface{}                               `json:"detailed_information"`
}

//ModAPIRequestProcessingResJSONMessageType содержит тип сообщения
// CriticalError - информационное сообщение о критической ошибке
// TaskExecutionResult - информационное сообщение о результате выполнения задачи
// DataTransfer - передача данных, например передача данных в результате успешного поиска информации
type ModAPIRequestProcessingResJSONMessageType struct {
	CriticalError       bool `json:"critical_error"`
	TaskExecutionResult bool `json:"error_during_data_processing"`
	DataTransfer        bool `json:"data_transfer"`
}

//ModAPIRequestProcessingResJSONCriticalError содержит информацию о критической ошибке, возникшей при обработке данных. Как правило
// подобные ошибки связаны с невозможностью корректно декодировать JSON сообщение
// Description - детальное описание ошибки (ОБЯЗАТЕЛЬНОЕ ПОЛЕ)
type ModAPIRequestProcessingResJSONCriticalError struct {
	Description string `json:"description"`
}

//ModAPIRequestProcessingResJSONTaskExecutionResult содержит результат обработки поставленной задачи
// TaskID - ID задачи
// Section - секция выполнения (ОБЯЗАТЕЛЬНОЕ ПОЛЕ)
// ExecutionResult - результат выполнения
//  Successful - задача успешно выполнена
//  Failure - задача не выполнена
//  PartiallySuccessful - задача выполнена частично
// Description - дополнительное детальное описание
// AdditionalParameters - дополнительные параметры связанные с выполняемой задачей
type ModAPIRequestProcessingResJSONTaskExecutionResult struct {
	ModAPIRequestProcessingCommonJSON
	Section         string `json:"section"`
	ExecutionResult struct {
		Successful          bool `json:"successful"`
		Failure             bool `json:"failure"`
		PartiallySuccessful bool `json:"partially_successful"`
	} `json:"execution_result"`
	Description          string      `json:"description"`
	AdditionalParameters interface{} `json:"additional_parameters"`
}

/*
	Example!!!
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

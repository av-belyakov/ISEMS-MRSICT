package datamodels

//ModAPIRequestProcessingCommonJSON содержит описание общих полей формата JSON для запросов и ответов модуля ModuleAPIRequestProcessing
// TaskID - ID задачи
type ModAPIRequestProcessingCommonJSON struct {
	TaskID string `json:"task_id"`
}

//ModAPIRequestProcessingReqJSON содержит описание формата JSON запроса получаемого через модуль ModuleAPIRequestProcessing
// TaskWasGeneratedAutomatically - задача была сгенерирована автоматически (true - да)
// UserNameGeneratedTask - имя пользователя сгенерировавшего задачу
// MessageType - тип сообщения
//  - "set info", добавить информацию
//  - "get info", получить информацию
// DetailedInformation - подбробная информация о запросе
type ModAPIRequestProcessingReqJSON struct {
	ModAPIRequestProcessingCommonJSON
	TaskWasGeneratedAutomatically bool          `json:"task_was_generated_automatically"`
	UserNameGeneratedTask         string        `json:"user_name_generated_task"`
	MessageType                   string        `json:"message_type"`
	DetailedInformation           []interface{} `json:"detailed_information"`
}

//ModAPIRequestProcessingResJSON содержит описание формата JSON ответа, передаваемого через модуль ModuleAPIRequestProcessing
// DetailedInformation - подбробная информация о запросе
// MessageType - тип сообщения
type ModAPIRequestProcessingResJSON struct {
	ModAPIRequestProcessingCommonJSON
	MessageType         MessageTypeModAPIRequestProcessingResJSON `json:"message_type"`
	DetailedInformation interface{}                               `json:"detailed_information"`
}

//MessageTypeModAPIRequestProcessingResJSON содержит описание типа ответа, передаваемого через модуль ModuleAPIRequestProcessing
// UserInformation - сообщения предназначенные пользователю: сообщения об успешном выполнении задачи, об ошибках и т.д.). В этом случае используется
//  тип NotificationModAPIRequestProcessingUserJSON
// PerformanceTasks - информация связанная с ходом выполнения задачи но не предназначенная для пользователя. В этом случае используется
//  тип PerformanceTasksModAPIRequestProcessingJSON
// TransmittingRequestedInformation - передача запрошенной информации. В этом случае используется
//  тип TransmittingRequestedInformationModAPIRequestProcessingResJSON
type MessageTypeModAPIRequestProcessingResJSON struct {
	UserInformation                  bool `json:"user_information"`
	PerformanceTasks                 bool `json:"performance_tasks"`
	TransmittingRequestedInformation bool `json:"transmitting_requested_information"`
}

//NotificationModAPIRequestProcessingUserJSON содержит информационное сообщение предназначенное пользователю и передаваемого
// через модуль ModuleAPIRequestProcessing
// Данное сообщение применяется когда MessageType основного сообщения равно "information message"
// MessageType - тип сообщения (success, warning, info, danger)
// Description - описание сообщения
type NotificationModAPIRequestProcessingUserJSON struct {
	ModAPIRequestProcessingCommonJSON
	MessageType string `json:"message_type"`
	Description string `json:"description"`
}

//PerformanceTasksModAPIRequestProcessingResJSON содержит информацию, связанную с ходом выполнения задачи и передаваемую через
// модуль ModuleAPIRequestProcessing. Данная информация предназначенна исключительно для передачи пользователю.
type PerformanceTasksModAPIRequestProcessingResJSON struct {
}

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

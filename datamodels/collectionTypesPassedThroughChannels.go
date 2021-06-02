package datamodels

/***  Описание типов данных циркулирующих по каналам связи внутри приложения  ***/

//CommanDataTypePassedThroughChannels описание общих параметров для типов данных циркулирующих по каналам связи внутри приложения
// ModuleGeneratorMessage - модуль генератор сообщения
// ModuleReceiverMessage - модуль приемник сообщения
//  Для свойств ModuleGeneratorMessage и ModuleReceiverMessage, доступны следующие предустановленные значения:
// - "module api request processing"
// - "module core application"
// - "module database interaction"
// InformationMessage - информационное сообщение
// ErrorMessage - подробное описание сообщения об ошибке
type CommanDataTypePassedThroughChannels struct {
	ModuleGeneratorMessage string
	ModuleReceiverMessage  string
	InformationMessage     InformationDataTypePassedThroughChannels
	ErrorMessage           ErrorDataTypePassedThroughChannels
}

//InformationDataTypePassedThroughChannels подробное описание информационного сообщения
// Type - тип информационного сообщения ('info', 'success', 'warning', 'danger')
// Message - информационное сообщение
type InformationDataTypePassedThroughChannels struct {
	Type    string
	Message string
}

//ErrorDataTypePassedThroughChannels подробное описание сообщения об ошибке
// FuncName - имя функции где возникла ошибка
// ModuleAPIRequestProcessingSettingSendTo - отправить информацию клиенту модуля moduleapirequestprocessing
// Error - подробное описание ошибки
type ErrorDataTypePassedThroughChannels struct {
	FuncName                                string
	ModuleAPIRequestProcessingSettingSendTo bool
	Error                                   error
}

//ModuleDataBaseInteractionChannel описание типов данных циркулирующих между модулем взаимодействия с БД и Ядром приложения
// Section - секция обработки данных
// Command - команда
// AppTaskID - внутренний идентификатор задачи
type ModuleDataBaseInteractionChannel struct {
	CommanDataTypePassedThroughChannels
	Section   string
	Command   string
	AppTaskID string
}

//ModuleReguestProcessingChannel описание типов данных циркулирующих между модулем обрабатывающем запросы с внешних источников и ядром приложения
// ClientID - уникальный идентификатор клиента (используется для идентификации соединения)
// ClientName - название клиента полученное из конфигурационного файла
// DataType - тип передаваемых данных (1 - text, 2 - binary)
// Data - данные
type ModuleReguestProcessingChannel struct {
	CommanDataTypePassedThroughChannels
	ClientID   string
	ClientName string
	DataType   int
	Data       *[]byte
}

/***  Описание 'общих' типов данных  ***/

//StorageApplicationCommonListType описание общего списка связанного с временной информацией по компьютерным угрозам
// ID - уникальный идентификатор STIX DO объекта связанного с информацией о компьютерной угрозе
// Description - дополнительное описание
type StorageApplicationCommonListType struct {
	ID          string
	Description string
}

package datamodels

/***  Описание типов данных циркулирующих по каналам связи внутри приложения  ***/

//CommanDataTypePassedThroughChannels описание общих параметров для типов данных циркулирующих по каналам связи внутри приложения
// ModuleGeneratorMessage - модуль генератор сообщения
// ModuleReceiverMessage - модуль приемник сообщения
//  Для свойств ModuleGeneratorMessage и ModuleReceiverMessage, доступны следующие предустановленные значения:
// - "module api request processing"
// - "module core application"
type CommanDataTypePassedThroughChannels struct {
	ModuleGeneratorMessage string
	ModuleReceiverMessage  string
}

//ModuleDataBaseInteractionChannel описание типов данных циркулирующих между модулем взаимодействия с БД и ядром приложения
// Section - секция данных
// Command - команда
// TaskID - внутренний идентификатор задачи
type ModuleDataBaseInteractionChannel struct {
	CommanDataTypePassedThroughChannels
	Section string
	Command string
	TaskID  string
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

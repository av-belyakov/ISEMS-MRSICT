package datamodels

/***  Описание типов данных циркулирующих по каналам связи внутри приложения  ***/

//CommanDataTypePassedThroughChannels описание общих параметров для типов данных циркулирующих по каналам связи внутри приложения
// ModuleGeneratorMessage - модуль генератор сообщения
// ModuleReceiverMessage - модуль приемник сообщения
// Command - команда
// TaskID - внутренний идентификатор задачи
type CommanDataTypePassedThroughChannels struct {
	MessageGeneratorModule string
	ModuleReceiverMessage  string
	Command                string
	TaskID                 string
}

//ModuleDataBaseInteractionChannel описание типов данных циркулирующих между модулем взаимодействия с БД и ядром приложения
// Section - секция данных
type ModuleDataBaseInteractionChannel struct {
	CommanDataTypePassedThroughChannels
	Section string
}

//ModuleReguestProcessingChannel описание типов данных циркулирующих между модулем обрабатывающем запросы с внешних источников и ядром приложения
// ClientID - уникальный идентификатор клиента (используется для идентификации соединения)
type ModuleReguestProcessingChannel struct {
	CommanDataTypePassedThroughChannels
	ClientID string
}

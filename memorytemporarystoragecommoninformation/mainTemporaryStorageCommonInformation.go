package memorytemporarystoragecommoninformation

import (
	"fmt"
	"time"
)

//TemporaryStorageType общее хранилище временной информации
// taskStorage - хранилище задач, где ключ является ID задачи
// chanReqTaskStorage - канал доступа к хранилищу задач
// foundInformationStorage - хранилище найденной информации
// chanReqFoundInformationStorage - канал доступа к хранилищу найденной информации
// parameterStorage - хранилище параметров
// chanReqParameterStorage - канал доступа к хранилищу параметров
type TemporaryStorageType struct {
	taskStorage                    map[string]TemporaryStorageTaskType
	chanReqTaskStorage             chan channelRequestTaskStorage
	foundInformationStorage        map[string]interface{}                     //(ЭТО всего лишь предположительный набросок)
	chanReqFoundInformationStorage chan channelRequestFoundInformationStorage //(ЭТО всего лишь предположительный набросок)
	parameterStorage               map[string]interface{}                     //(ЭТО всего лишь предположительный набросок)
	chanReqParameterStorage        chan channelRequestParameterStorage        //(ЭТО всего лишь предположительный набросок)
}

//TemporaryStorageTaskType описание задачи обрабатываемой приложением
// TaskGenerator - источник-генератор задачи (фактически название того модуля что прописано в CommanDataTypePassedThroughChannels.ModuleGeneratorMessage)
// ClientID - уникальный идентификатор клиента, каким либо образом связанного с источником-генератором задачи
// ClientName - имя клиента (возможно логин)
// ClientTaskID - идентификатор задачи, переданный клиентом, если есть
// AdditionalClientName - дополнительное имя клиента (возможно использовать Ф.И.О.)
// DateTaskCreated - дата создания задачи
// DateTaskModification - дата модификации задачи (данный параметр необходим для удаления 'старых' задач)
// TaskStatus - статус выполнения задачи. Данный параметр напрямую влияет на удаление 'старых' задач. Предусмотрены следующие значения:
//  - "wait"
//  - "in progress"
//  - "completed"
// Section - секция обработки данных.
// Command - команда для выполнения задачи (не обязательный параметр)
// TaskParameters - параметры задачи. Напрямую зависит от параметра находящегося в поле Section.
type TemporaryStorageTaskType struct {
	TaskGenerator        string
	ClientID             string
	ClientName           string
	ClientTaskID         string
	AdditionalClientName string
	DateTaskCreated      time.Time
	DateTaskModification time.Time
	TaskStatus           string
	Section              string
	Command              string
	TaskParameters       interface{}
}

//channelRequest канал через который выполняются запросы к внутреннему обработчику хранилища задач приложения
// actionType - тип действия
// appTaskID - внутренний идентификатор задачи
// chanRes - канал через который будет получен ответ от хранилища
type channelRequestTaskStorage struct {
	typeTemporaryStorage string
	appTaskID            string
	actionType           string
	chanRes              chan channelResponseTaskStorage
}

type channelResponseTaskStorage struct {
	appTaskID string
	errMsg    error
}

//канал через который выполняются запросы к внутреннему обработчику хранилища параметров приложения
// actionType - тип действия
// parameterID - внутренний идентификатор параметра
// chanRes - канал через который будет получен ответ от хранилища
type channelRequestParameterStorage struct {
	actionType  string
	parameterID string
	chanRes     chan channelResponseParameterStorage
}

type channelResponseParameterStorage struct{}

type channelRequestFoundInformationStorage struct{}
type channelResponseFoundInformationStorage struct{}

//NewTemporaryStorage конструктор инициализирующий временное хранилище общей информации
func NewTemporaryStorage() *TemporaryStorageType {
	fmt.Println("fun 'NewStorageTemporaryMemoryCommon', START...")

	chanReqTask := make(chan channelRequestTaskStorage)
	chanReqParameter := make(chan channelRequestParameterStorage)
	chanReqFoundInfo := make(chan channelRequestFoundInformationStorage)

	stmc := TemporaryStorageType{
		taskStorage:                    map[string]TemporaryStorageTaskType{},
		chanReqTaskStorage:             chanReqTask,
		foundInformationStorage:        map[string]interface{}{},
		chanReqFoundInformationStorage: chanReqFoundInfo,
		parameterStorage:               map[string]interface{}{},
		chanReqParameterStorage:        chanReqParameter,
	}

	go func() {
		select {
		case msg := <-chanReqTask:
			switch msg.actionType {
			case "":
			}

		case msg := <-chanReqFoundInfo:
			fmt.Println(msg)

		case msg := <-chanReqParameter:
			fmt.Println(msg)

		}
	}()

	return &stmc
}

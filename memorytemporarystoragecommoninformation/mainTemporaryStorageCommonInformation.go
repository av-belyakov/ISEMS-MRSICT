package memorytemporarystoragecommoninformation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//TemporaryStorageType общее хранилище временной информации
// taskStorage - хранилище задач, где ключ является внутренним идентификатором задачи (appTaskID)
// chanReqTaskStorage - канал доступа к хранилищу задач
// foundInformationStorage - хранилище найденной информации
// chanReqFoundInformationStorage - канал доступа к хранилищу найденной информации
// storageApplicationParameters - хранилище параметров
// chanReqParameterStorage - канал доступа к хранилищу параметров
type TemporaryStorageType struct {
	taskStorage                    map[string]TemporaryStorageTaskInDetailType
	chanReqTaskStorage             chan channelRequestTaskStorage
	foundInformationStorage        map[string]interface{}                     //(ЭТО всего лишь предположительный набросок)
	chanReqFoundInformationStorage chan channelRequestFoundInformationStorage //(ЭТО всего лишь предположительный набросок)
	storageApplicationParameters   storageApplicationParametersType
	chanReqParameterStorage        chan channelRequestParameterStorage //(ЭТО всего лишь предположительный набросок)
}

//TemporaryStorageTaskType описание задачи обрабатываемой приложением
// TaskGenerator - источник-генератор задачи (фактически название того модуля что прописано в CommanDataTypePassedThroughChannels.ModuleGeneratorMessage)
// ClientID - уникальный идентификатор клиента, каким либо образом связанного с источником-генератором задачи
// ClientName - имя клиента (возможно логин)
// ClientTaskID - идентификатор задачи, переданный клиентом, если есть
// AdditionalClientName - дополнительное имя клиента (возможно использовать Ф.И.О.)
// Section - секция обработки данных.
// Command - команда для выполнения задачи (не обязательный параметр)
// TaskStatus - статус выполнения задачи. Предусмотрены следующие значения:
//  - "wait"
//  - "in progress"
//  - "completed"
// TaskParameters - параметры задачи. Напрямую зависит от параметра находящегося в поле Section.
type TemporaryStorageTaskType struct {
	TaskGenerator        string
	ClientID             string
	ClientName           string
	ClientTaskID         string
	AdditionalClientName string
	Section              string
	Command              string
	TaskStatus           string
	TaskParameters       interface{}
}

//TemporaryStorageTaskInDetailType подробное описание задачи обрабатываемой приложением
// DateTaskCreated - дата создания задачи
// DateTaskModification - дата модификации задачи (данный параметр необходим для удаления 'старых' задач)
// RemovalRequired - требуется удаление. Считается что данный объект устарел и его необходимо удалить. Данный параметр напрямую влияет на удаление 'старых' задач.
type TemporaryStorageTaskInDetailType struct {
	TemporaryStorageTaskType
	DateTaskCreated      time.Time
	DateTaskModification time.Time
	RemovalRequired      bool
}

//commanChannelTaskStorage общее описание типа для каналов взаимодействия с хранилищем задач
// appTaskID - внутренний идентификатор задачи
// detailedDescriptionTask - подробное описание задачи
type commanChannelTaskStorage struct {
	appTaskID string
}

//channelRequestTaskStorage канал, через который выполняются запросы к внутреннему обработчику хранилища задач приложения
// actionType - тип действия
// chanRes - канал, через который будет получен ответ от хранилища
type channelRequestTaskStorage struct {
	commanChannelTaskStorage
	actionType              string
	detailedDescriptionTask *TemporaryStorageTaskType
	chanRes                 chan channelResponseTaskStorage
}

//channelResponseTaskStorage канал, через который поступает информация от внутреннего обработчика хранилища задач
// listAppTasksID - список внутренних идентификаторов задач
// errMsg - сообщение об ошибке
type channelResponseTaskStorage struct {
	commanChannelTaskStorage
	listAppTasksID          []string
	detailedDescriptionTask *TemporaryStorageTaskInDetailType
	errMsg                  error
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

//storageApplicationParametersType хранилище параметров приложения
type storageApplicationParametersType struct{}

type channelResponseParameterStorage struct {
	parameterType string
	actionType    string
}

type channelRequestFoundInformationStorage struct{}
type channelResponseFoundInformationStorage struct{}

//NewTemporaryStorage конструктор инициализирующий временное хранилище общей информации
func NewTemporaryStorage() *TemporaryStorageType {
	fmt.Println("fun 'NewStorageTemporaryMemoryCommon', START...")

	chanReqTask := make(chan channelRequestTaskStorage)
	chanReqParameter := make(chan channelRequestParameterStorage)
	chanReqFoundInfo := make(chan channelRequestFoundInformationStorage)

	stmc := TemporaryStorageType{
		taskStorage:                    map[string]TemporaryStorageTaskInDetailType{},
		chanReqTaskStorage:             chanReqTask,
		foundInformationStorage:        map[string]interface{}{},
		chanReqFoundInformationStorage: chanReqFoundInfo,
		storageApplicationParameters:   storageApplicationParametersType{},
		chanReqParameterStorage:        chanReqParameter,
	}

	go func() {
		for {
			select {
			case msg := <-chanReqTask:
				switch msg.actionType {
				case "add new task":
					uuid := uuid.NewString()
					err := stmc.addNewTask(uuid, msg.detailedDescriptionTask)
					msg.chanRes <- channelResponseTaskStorage{
						commanChannelTaskStorage: commanChannelTaskStorage{
							appTaskID: uuid,
						},
						errMsg: err,
					}

				case "get task by id":
					taskInfo, err := stmc.getTaskByID(msg.appTaskID)

					fmt.Println("sqsqsqsqqsqsqsqsqsqsq")
					fmt.Println(taskInfo)

					msg.chanRes <- channelResponseTaskStorage{
						detailedDescriptionTask: taskInfo,
						errMsg:                  err,
					}

					fmt.Println("vvvvvvvvvvvvvvv")

				case "get tasks by client id":
					msg.chanRes <- channelResponseTaskStorage{listAppTasksID: stmc.getTasksByClientID(msg.detailedDescriptionTask.ClientID)}

				case "change task status":
					msg.chanRes <- channelResponseTaskStorage{errMsg: stmc.changeTaskStatus(msg.appTaskID, msg.detailedDescriptionTask.TaskStatus)}

				case "change removal required parameter":
					msg.chanRes <- channelResponseTaskStorage{errMsg: stmc.changeRemovalRequiredParameter(msg.appTaskID)}

				case "change date task modification":
					msg.chanRes <- channelResponseTaskStorage{errMsg: stmc.changeDateTaskModification(msg.appTaskID)}

				case "deleting task by id":
					stmc.deletingTaskByID(msg.appTaskID)

					msg.chanRes <- channelResponseTaskStorage{}

				}

			case msg := <-chanReqFoundInfo:
				fmt.Println(msg)

			case msg := <-chanReqParameter:
				fmt.Println(msg)

			}
		}
	}()

	return &stmc
}

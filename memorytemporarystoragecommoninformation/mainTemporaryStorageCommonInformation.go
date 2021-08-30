package memorytemporarystoragecommoninformation

import (
	"sync"
	"time"

	"ISEMS-MRSICT/datamodels"

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
	foundInformationStorage        map[string]TemporaryStorageFoundInformation
	chanReqFoundInformationStorage chan channelRequestFoundInformationStorage
	storageApplicationParameters   StorageApplicationParametersType
	chanReqParameterStorage        chan channelRequestStorageApplicationParameters
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

//commonChannelTaskStorage общее описание типа для каналов взаимодействия с хранилищем задач
// appTaskID - внутренний идентификатор задачи
// detailedDescriptionTask - подробное описание задачи
type commonChannelTaskStorage struct {
	appTaskID string
}

//channelRequestTaskStorage канал, через который выполняются запросы к внутреннему обработчику хранилища задач приложения
// actionType - тип действия
// detailedDescriptionTask - детальное описание задач
// chanRes - канал, через который будет получен ответ от хранилища
type channelRequestTaskStorage struct {
	commonChannelTaskStorage
	actionType              string
	detailedDescriptionTask *TemporaryStorageTaskType
	chanRes                 chan channelResponseTaskStorage
}

//channelResponseTaskStorage канал, через который поступает информация от внутреннего обработчика хранилища задач
// listAppTasksID - список внутренних идентификаторов задач
// detailedDescriptionTask - детальное описание найденных задач
// errMsg - сообщение об ошибке
type channelResponseTaskStorage struct {
	commonChannelTaskStorage
	listAppTasksID          []string
	detailedDescriptionTask *TemporaryStorageTaskInDetailType
	errMsg                  error
}

//TemporaryStorageFoundInformation подробное описание информации найденной в результате выполнения задачи поиска
// Collection - коллекция в которой выполнялся поиск информации ('stix_object_collection', 'reference_book_collection' и т.д.)
// ResultType - тип результата ('only_count', 'full_found_info')
// Information interface{}
type TemporaryStorageFoundInformation struct {
	Collection  string
	ResultType  string
	Information interface{}
}

//StorageApplicationParametersType хранилище параметров приложения
// ListTypesDecisionsMadeComputerThreat - типы принимаемых решений по компьютерным угрозам
// ListTypeComputerThreat - список типов компьютерных угроз
type StorageApplicationParametersType struct {
	ListTypesDecisionsMadeComputerThreat map[string]datamodels.StorageApplicationCommonListType
	ListTypesComputerThreat              map[string]datamodels.StorageApplicationCommonListType
}

//channelRequestFoundInformationStorage канал через который передаются запросы к хранилищу найденной информации
// appTaskID - идентификатор задачи
// actionType - тип действия
// description - описание найденной информации
// chanRes - канал для передачи ответа
type channelRequestFoundInformationStorage struct {
	commonChannelTaskStorage
	actionType  string
	description *TemporaryStorageFoundInformation
	chanRes     chan channelResponseFoundInformationStorage
}

//channelResponseFoundInformationStorage канал через который передаются ответы от хранилища найденной информации
// appTaskID - идентификатор задачи
// description - описание найденной информации
// errMsg - сообщение об ошибке
type channelResponseFoundInformationStorage struct {
	commonChannelTaskStorage
	description *TemporaryStorageFoundInformation
	errMsg      error
}

//channelRequestStorageApplicationParameters канал через который выполняются запросы к внутреннему обработчику хранилища параметров приложения
// actionType - тип действия
// parameterStorage - параметр
// chanRes - канал через который будет получен ответ от хранилища
type channelRequestStorageApplicationParameters struct {
	actionType       string
	parameterStorage interface{}
	chanRes          chan channelResponseStorageApplicationParameters
}

//channelResponseParameterStorage канал через который передаются ответы с параметрами приложения (ПОКА ЗАГЛУШКА)
// dataParameterStorage - передаваемые параметры приложения
// errMsg - сообщение об ошибке
type channelResponseStorageApplicationParameters struct {
	dataParameterStorage interface{}
	errMsg               error
}

var once sync.Once
var stmc TemporaryStorageType

//NewTemporaryStorage конструктор инициализирующий временное хранилище общей информации
func NewTemporaryStorage() *TemporaryStorageType {
	once.Do(func() {
		chanReqTask := make(chan channelRequestTaskStorage)
		chanReqFoundInfo := make(chan channelRequestFoundInformationStorage)
		chanReqParameter := make(chan channelRequestStorageApplicationParameters)

		stmc = TemporaryStorageType{
			taskStorage:                    map[string]TemporaryStorageTaskInDetailType{},
			chanReqTaskStorage:             chanReqTask,
			foundInformationStorage:        map[string]TemporaryStorageFoundInformation{},
			chanReqFoundInformationStorage: chanReqFoundInfo,
			storageApplicationParameters:   StorageApplicationParametersType{},
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
							commonChannelTaskStorage: commonChannelTaskStorage{
								appTaskID: uuid,
							},
							errMsg: err,
						}

					case "get task by id":
						taskInfo, err := stmc.getTaskByID(msg.appTaskID)

						msg.chanRes <- channelResponseTaskStorage{
							detailedDescriptionTask: taskInfo,
							errMsg:                  err,
						}

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
					switch msg.actionType {
					case "add new information":
						msg.chanRes <- channelResponseFoundInformationStorage{
							errMsg: stmc.addNewFoundInformation(msg.appTaskID, msg.description),
						}

					case "get information by id":
						info, err := stmc.getFoundInformationByID(msg.appTaskID)

						msg.chanRes <- channelResponseFoundInformationStorage{
							description: info,
							errMsg:      err,
						}

					case "delete information by id":
						stmc.deletingFoundInformationByID(msg.appTaskID)

						msg.chanRes <- channelResponseFoundInformationStorage{}

					}

				case msg := <-chanReqParameter:
					switch msg.actionType {
					case "set list decisions made":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							errMsg: stmc.setListDecisionsMade(msg.parameterStorage),
						}

					case "get list decisions made":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							dataParameterStorage: stmc.getListDecisionsMade(),
						}

					case "get id decisions made type: successfully implemented computer threat":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							dataParameterStorage: stmc.getIDDecisionsMadeSuccessfully(),
						}

					case "get id decisions made type: unsuccessfully computer threat":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							dataParameterStorage: stmc.getIDDecisionsMadeUnsuccessfully(),
						}

					case "get id decisions made type: false positive":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							dataParameterStorage: stmc.getIDDecisionsMadeFalsePositive(),
						}

					case "set type computer threat":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							errMsg: stmc.setListComputerThreat(msg.parameterStorage),
						}

					case "get type computer threat":
						msg.chanRes <- channelResponseStorageApplicationParameters{
							dataParameterStorage: stmc.getListComputerThreat(),
						}

					}
				}
			}
		}()
	})

	return &stmc
}

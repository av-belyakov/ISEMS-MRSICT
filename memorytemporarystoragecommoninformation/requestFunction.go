package memorytemporarystoragecommoninformation

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
)

/*** МЕТОДЫ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ЗАДАЧ ***/

//AddNewTask метод добавляющий новую задачу в хранилище задач
func (tst *TemporaryStorageType) AddNewTask(taskInfo *TemporaryStorageTaskType) (string, error) {
	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType:              "add new task",
		detailedDescriptionTask: taskInfo,
		chanRes:                 chanRes,
	}
	result := <-chanRes

	return result.appTaskID, result.errMsg
}

//GetTaskByID метод возвращающий всю информацию о задаче, по ее ID
func (tst *TemporaryStorageType) GetTaskByID(appTaskID string) (string, *TemporaryStorageTaskInDetailType, error) {
	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType:               "get task by id",
		commonChannelTaskStorage: commonChannelTaskStorage{appTaskID: appTaskID},
		chanRes:                  chanRes,
	}
	result := <-chanRes

	return appTaskID, result.detailedDescriptionTask, result.errMsg
}

//GetTasksByClientID метод возвращающий список задач принадлежащих клиенту с определенным ID
func (tst *TemporaryStorageType) GetTasksByClientID(clientID string) []string {
	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType:              "get tasks by client id",
		detailedDescriptionTask: &TemporaryStorageTaskType{ClientID: clientID},
		chanRes:                 chanRes,
	}

	return (<-chanRes).listAppTasksID
}

//ChangeRemovalRequiredParameter метод устанавливает параметр RemovalRequired в TRUE
func (tst *TemporaryStorageType) ChangeRemovalRequiredParameter(appTaskID string) error {
	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType: "change removal required parameter",
		commonChannelTaskStorage: commonChannelTaskStorage{
			appTaskID: appTaskID,
		},
		chanRes: chanRes,
	}

	return (<-chanRes).errMsg
}

//ChangeDateTaskModification метод меняющий время модификации задачи
func (tst *TemporaryStorageType) ChangeDateTaskModification(appTaskID string) error {
	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType: "change date task modification",
		commonChannelTaskStorage: commonChannelTaskStorage{
			appTaskID: appTaskID,
		},
		chanRes: chanRes,
	}

	return (<-chanRes).errMsg
}

//ChangeTaskStatus метод меняющий статус выполнения задачи
func (tst *TemporaryStorageType) ChangeTaskStatus(appTaskID, taskStatus string) error {
	var isExist bool
	statuses := []string{"wait", "in progress", "completed"}
	for key := range statuses {
		if statuses[key] == taskStatus {
			isExist = true

			break
		}
	}

	if !isExist {
		return fmt.Errorf("the undefined status of the task is accepted")
	}

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType: "change task status",
		commonChannelTaskStorage: commonChannelTaskStorage{
			appTaskID: appTaskID,
		},
		detailedDescriptionTask: &TemporaryStorageTaskType{TaskStatus: taskStatus},
		chanRes:                 chanRes,
	}

	return (<-chanRes).errMsg
}

//DeletingTaskByID удаление задачи по ее ID
func (tst *TemporaryStorageType) DeletingTaskByID(appTaskID string) {
	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType:               "deleting task by id",
		commonChannelTaskStorage: commonChannelTaskStorage{appTaskID: appTaskID},
		chanRes:                  chanRes,
	}

	<-chanRes
}

/*** МЕТОДЫ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ НАЙДЕННОЙ ИНФОРМАЦИИ ***/

//AddNewFoundInformation метод добавляющий новую информацию, полученную в ходе выполнения поиска, в хранилище найденной информации
func (tst *TemporaryStorageType) AddNewFoundInformation(appTaskID string, info *TemporaryStorageFoundInformation) error {
	chanRes := make(chan channelResponseFoundInformationStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqFoundInformationStorage <- channelRequestFoundInformationStorage{
		commonChannelTaskStorage: commonChannelTaskStorage{appTaskID: appTaskID},
		actionType:               "add new information",
		description:              info,
		chanRes:                  chanRes,
	}

	return (<-chanRes).errMsg
}

//GetFoundInformationByID метод возвращающий всю информацию найденную, в результате поиска, информацию по ее ID
func (tst *TemporaryStorageType) GetFoundInformationByID(appTaskID string) (*TemporaryStorageFoundInformation, error) {
	chanRes := make(chan channelResponseFoundInformationStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqFoundInformationStorage <- channelRequestFoundInformationStorage{
		commonChannelTaskStorage: commonChannelTaskStorage{appTaskID: appTaskID},
		actionType:               "get information by id",
		chanRes:                  chanRes,
	}

	result := <-chanRes

	return result.description, result.errMsg
}

//DeletingFoundInformationByID удаление задачи по ее ID
func (tst *TemporaryStorageType) DeletingFoundInformationByID(appTaskID string) {
	chanRes := make(chan channelResponseFoundInformationStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqFoundInformationStorage <- channelRequestFoundInformationStorage{
		commonChannelTaskStorage: commonChannelTaskStorage{appTaskID: appTaskID},
		actionType:               "delete information by id",
		chanRes:                  chanRes,
	}

	<-chanRes
}

/*** МЕТОДЫ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ПАРАМЕТРОВ ПРИЛОЖЕНИЯ ***/

//SetListDecisionsMade добавление списка принимаемых решений по компьютерным угрозам
func (tst *TemporaryStorageType) SetListDecisionsMade(l map[string]datamodels.StorageApplicationCommonListType) error {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType:       "set list decisions made",
		parameterStorage: l,
		chanRes:          chanRes,
	}

	return (<-chanRes).errMsg
}

//GetListDecisionsMade возвращение списка принимаемых решений по компьютерным угрозам
func (tst *TemporaryStorageType) GetListDecisionsMade() (map[string]datamodels.StorageApplicationCommonListType, error) {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType: "get list decisions made",
		chanRes:    chanRes,
	}

	result := <-chanRes
	l, ok := result.dataParameterStorage.(map[string]datamodels.StorageApplicationCommonListType)
	if !ok {
		return l, fmt.Errorf("type conversion error")
	}

	return l, result.errMsg
}

//GetIDDecisionsMadeSuccessfully возвращение ID решения 'successfully implemented computer threat' по компьютерным угрозам
func (tst *TemporaryStorageType) GetIDDecisionsMadeSuccessfully() (datamodels.StorageApplicationCommonListType, error) {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType: "get id decisions made type: successfully implemented computer threat",
		chanRes:    chanRes,
	}

	result := <-chanRes
	l, ok := result.dataParameterStorage.(datamodels.StorageApplicationCommonListType)
	if !ok {
		return l, fmt.Errorf("type conversion error")
	}

	return l, result.errMsg
}

//GetIDDecisionsMadeUnsuccessfully возвращение ID решения 'unsuccessfully computer threat' по компьютерным угрозам
func (tst *TemporaryStorageType) GetIDDecisionsMadeUnsuccessfully() (datamodels.StorageApplicationCommonListType, error) {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType: "get id decisions made type: unsuccessfully computer threat",
		chanRes:    chanRes,
	}

	result := <-chanRes
	l, ok := result.dataParameterStorage.(datamodels.StorageApplicationCommonListType)
	if !ok {
		return l, fmt.Errorf("type conversion error")
	}

	return l, result.errMsg
}

//GetIDDecisionsMadeFalsePositive возвращение ID решения 'false positive' по компьютерным угрозам
func (tst *TemporaryStorageType) GetIDDecisionsMadeFalsePositive() (datamodels.StorageApplicationCommonListType, error) {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType: "get id decisions made type: false positive",
		chanRes:    chanRes,
	}

	result := <-chanRes
	l, ok := result.dataParameterStorage.(datamodels.StorageApplicationCommonListType)
	if !ok {
		return l, fmt.Errorf("type conversion error")
	}

	return l, result.errMsg
}

//SetListComputerThreat добавление списка типов компьютерных угроз
func (tst *TemporaryStorageType) SetListComputerThreat(l map[string]datamodels.StorageApplicationCommonListType) error {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType:       "set type computer threat",
		parameterStorage: l,
		chanRes:          chanRes,
	}

	return (<-chanRes).errMsg
}

//GetListComputerThreat возвращение списка типов компьютерных угроз
func (tst *TemporaryStorageType) GetListComputerThreat() (map[string]datamodels.StorageApplicationCommonListType, error) {
	chanRes := make(chan channelResponseStorageApplicationParameters)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqParameterStorage <- channelRequestStorageApplicationParameters{
		actionType: "get type computer threat",
		chanRes:    chanRes,
	}

	result := <-chanRes
	l, ok := result.dataParameterStorage.(map[string]datamodels.StorageApplicationCommonListType)
	if !ok {
		return l, fmt.Errorf("type conversion error")
	}

	return l, result.errMsg
}

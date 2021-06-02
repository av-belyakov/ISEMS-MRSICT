package memorytemporarystoragecommoninformation

import (
	"fmt"
	"time"

	"ISEMS-MRSICT/datamodels"
)

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ЗАДАЧ ***/

//addNewTask добавляет новую задачу в хранилище задач
func (tst *TemporaryStorageType) addNewTask(appTaskID string, taskInfo *TemporaryStorageTaskType) error {
	if ok := tst.taskIsExist(appTaskID); ok {
		return fmt.Errorf("the task with the ID '%s' already exists in the repository 'taskStorage'", appTaskID)
	}

	tst.taskStorage[appTaskID] = TemporaryStorageTaskInDetailType{
		TemporaryStorageTaskType: TemporaryStorageTaskType{
			TaskGenerator:        taskInfo.TaskGenerator,
			ClientID:             taskInfo.ClientID,
			ClientName:           taskInfo.ClientName,
			ClientTaskID:         taskInfo.ClientTaskID,
			AdditionalClientName: taskInfo.AdditionalClientName,
			Section:              taskInfo.Section,
			Command:              taskInfo.Command,
			TaskStatus:           "wait",
			TaskParameters:       taskInfo.TaskParameters,
		},
		DateTaskCreated:      time.Now(),
		DateTaskModification: time.Now(),
		RemovalRequired:      false,
	}

	return nil
}

//getTaskByID возвращает информацию об одной задаче найденной по её ID
func (tst *TemporaryStorageType) getTaskByID(appTaskID string) (*TemporaryStorageTaskInDetailType, error) {
	if ok := tst.taskIsExist(appTaskID); !ok {
		return nil, fmt.Errorf("no tasks with an ID '%s' were found in the repository 'taskStorage'", appTaskID)
	}
	taskInfo, _ := tst.taskStorage[appTaskID]

	return &taskInfo, nil
}

//getTasksByClientID метод возвращающий список задач принадлежащих клиенту с определенным ID
func (tst *TemporaryStorageType) getTasksByClientID(clientID string) []string {
	result := []string{}

	for appTaskID, item := range tst.taskStorage {
		if item.ClientID == clientID {
			result = append(result, appTaskID)
		}
	}

	return result
}

//changeRemovalRequiredParameter метод устанавливает параметр RemovalRequired в TRUE
func (tst *TemporaryStorageType) changeRemovalRequiredParameter(appTaskID string) error {
	if ok := tst.taskIsExist(appTaskID); !ok {
		return fmt.Errorf("no tasks with an ID '%s' were found in the repository 'taskStorage'", appTaskID)
	}

	ts, _ := tst.taskStorage[appTaskID]
	ts.RemovalRequired = true
	tst.taskStorage[appTaskID] = ts

	return nil
}

//changeDateTaskModification изменяет параметр DateTaskModification
func (tst *TemporaryStorageType) changeDateTaskModification(appTaskID string) error {
	if ok := tst.taskIsExist(appTaskID); !ok {
		return fmt.Errorf("no tasks with an ID '%s' were found in the repository 'taskStorage'", appTaskID)
	}

	ts, _ := tst.taskStorage[appTaskID]
	ts.DateTaskModification = time.Now()
	tst.taskStorage[appTaskID] = ts

	return nil
}

//changeTaskStatus меняет статус задачи и параметр DateTaskModification
func (tst *TemporaryStorageType) changeTaskStatus(appTaskID, taskStatus string) error {
	if ok := tst.taskIsExist(appTaskID); !ok {
		return fmt.Errorf("no tasks with an ID '%s' were found in the repository 'taskStorage'", appTaskID)
	}

	ts, _ := tst.taskStorage[appTaskID]
	ts.TaskStatus = taskStatus
	ts.DateTaskModification = time.Now()
	tst.taskStorage[appTaskID] = ts

	return nil
}

//deletingTaskByID удаляет задачу по ее ID
func (tst *TemporaryStorageType) deletingTaskByID(appTaskID string) {
	delete(tst.taskStorage, appTaskID)
}

//taskIsExist проверяет наличие задачи
func (tst *TemporaryStorageType) taskIsExist(appTaskID string) bool {
	_, ok := (*tst).taskStorage[appTaskID]

	return ok
}

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ НАЙДЕННОЙ ИНФОРМАЦИИ ***/

//addNewFoundInformation добавляет новую найденную информацию
func (tst *TemporaryStorageType) addNewFoundInformation(appTaskID string, info *TemporaryStorageFoundInformation) error {
	if _, ok := tst.foundInformationStorage[appTaskID]; ok {
		return fmt.Errorf("the found information with the ID '%s' already exists in the repository 'foundInformationStorage'", appTaskID)
	}

	tst.foundInformationStorage[appTaskID] = *info

	return nil
}

//getFoundInformationByID возвращает информацию об одной задаче найденной по её ID
func (tst *TemporaryStorageType) getFoundInformationByID(appTaskID string) (*TemporaryStorageFoundInformation, error) {
	info, ok := tst.foundInformationStorage[appTaskID]
	if !ok {
		return nil, fmt.Errorf("no found information with an ID '%s' were found in the repository 'foundInformationStorage'", appTaskID)
	}

	return &info, nil
}

//deletingFoundInformationByID удаляет найденную информацию по ее ID
func (tst *TemporaryStorageType) deletingFoundInformationByID(appTaskID string) {
	delete(tst.foundInformationStorage, appTaskID)
}

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ПАРАМЕТРОВ ПРИЛОЖЕНИЯ ***/

//setListDecisionsMade добавление списка принимаемых решений по компьютерным угрозам
func (tst *TemporaryStorageType) setListDecisionsMade(l interface{}) error {
	list, ok := l.(map[string]datamodels.StorageApplicationCommonListType)
	if !ok {
		return fmt.Errorf("type conversion error")
	}

	tst.storageApplicationParameters.ListTypesDecisionsMadeComputerThreat = list

	return nil
}

//getListDecisionsMade возвращение списка принимаемых решений по компьютерным угрозам
func (tst *TemporaryStorageType) getListDecisionsMade() map[string]datamodels.StorageApplicationCommonListType {
	return tst.storageApplicationParameters.ListTypesDecisionsMadeComputerThreat
}

//getIDDecisionsMadeSuccessfully возвращение ID решения 'successfully implemented computer threat' по компьютерным угрозам
func (tst *TemporaryStorageType) getIDDecisionsMadeSuccessfully() datamodels.StorageApplicationCommonListType {
	return tst.storageApplicationParameters.ListTypesDecisionsMadeComputerThreat["successfully implemented computer threat"]
}

//getIDDecisionsMadeUnsuccessfully возвращение ID решения 'unsuccessfully computer threat' по компьютерным угрозам
func (tst *TemporaryStorageType) getIDDecisionsMadeUnsuccessfully() datamodels.StorageApplicationCommonListType {
	return tst.storageApplicationParameters.ListTypesDecisionsMadeComputerThreat["unsuccessfully computer threat"]
}

//getIDDecisionsMadeFalsePositive возвращение ID решения 'false positive' по компьютерным угрозам
func (tst *TemporaryStorageType) getIDDecisionsMadeFalsePositive() datamodels.StorageApplicationCommonListType {
	return tst.storageApplicationParameters.ListTypesDecisionsMadeComputerThreat["false positive"]
}

//setListComputerThreat добавление списка типов компьютерных угроз
func (tst *TemporaryStorageType) setListComputerThreat(l interface{}) error {
	list, ok := l.(map[string]datamodels.StorageApplicationCommonListType)
	if !ok {
		return fmt.Errorf("type conversion error")
	}

	tst.storageApplicationParameters.ListTypesComputerThreat = list

	return nil
}

//getListComputerThreat возвращение списка типов компьютерных угроз
func (tst *TemporaryStorageType) getListComputerThreat() map[string]datamodels.StorageApplicationCommonListType {
	return tst.storageApplicationParameters.ListTypesComputerThreat
}

package memorytemporarystoragecommoninformation

import (
	"fmt"
	"time"
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

//getOneTask возвращает информацию об одной задаче найденной по её ID
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

func (tst *TemporaryStorageType) deletingTaskByID(appTaskID string) {
	delete(tst.taskStorage, appTaskID)
}

func (tst *TemporaryStorageType) taskIsExist(appTaskID string) bool {
	_, ok := (*tst).taskStorage[appTaskID]

	return ok
}

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ НАЙДЕННОЙ ИНФОРМАЦИИ ***/

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ПАРАМЕТРОВ ПРИЛОЖЕНИЯ ***/

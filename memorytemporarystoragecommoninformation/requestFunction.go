package memorytemporarystoragecommoninformation

import (
	"fmt"
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
		commanChannelTaskStorage: commanChannelTaskStorage{appTaskID: appTaskID},
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
		commanChannelTaskStorage: commanChannelTaskStorage{
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
		commanChannelTaskStorage: commanChannelTaskStorage{
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
		commanChannelTaskStorage: commanChannelTaskStorage{
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
		commanChannelTaskStorage: commanChannelTaskStorage{appTaskID: appTaskID},
		chanRes:                  chanRes,
	}
	<-chanRes
}

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ НАЙДЕННОЙ ИНФОРМАЦИИ ***/

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ПАРАМЕТРОВ ПРИЛОЖЕНИЯ ***/

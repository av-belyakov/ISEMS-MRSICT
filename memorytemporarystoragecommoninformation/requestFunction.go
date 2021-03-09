package memorytemporarystoragecommoninformation

import (
	"fmt"
)

/*** МЕТОДЫ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ЗАДАЧ ***/

//AddNewTask метод добавляющий новую задачу в хранилище задач
func (tst *TemporaryStorageType) AddNewTask(taskInfo *TemporaryStorageTaskType) (string, error) {
	fmt.Println("func 'AddNewTask', START...")

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
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType:               "get task by id",
		commanChannelTaskStorage: commanChannelTaskStorage{appTaskID: appTaskID},
	}
	result := <-chanRes

	return appTaskID, result.detailedDescriptionTask, result.errMsg
}

//GetTasksByClientID метод возвращающий список задач принадлежащих клиенту с определенным ID
func (tst *TemporaryStorageType) GetTasksByClientID(clientID string) []string {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType:              "get tasks by client id",
		detailedDescriptionTask: &TemporaryStorageTaskType{ClientID: clientID},
	}

	return (<-chanRes).listAppTasksID
}

//ChangeRemovalRequiredParameter метод устанавливает параметр RemovalRequired в TRUE
func (tst *TemporaryStorageType) ChangeRemovalRequiredParameter(appTaskID string) error {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType: "change removal required parameter",
		commanChannelTaskStorage: commanChannelTaskStorage{
			appTaskID: appTaskID,
		},
	}

	return (<-chanRes).errMsg
}

//ChangeDateTaskModification метод меняющий время модификации задачи
func (tst *TemporaryStorageType) ChangeDateTaskModification(appTaskID string) error {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{
		actionType: "change date task modification",
		commanChannelTaskStorage: commanChannelTaskStorage{
			appTaskID: appTaskID,
		},
	}

	return (<-chanRes).errMsg
}

//ChangeTaskStatus метод меняющий статус выполнения задачи
func (tst *TemporaryStorageType) ChangeTaskStatus(appTaskID, taskStatus string) error {
	fmt.Println("func 'GetTaskByID', START...")

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
	}

	return (<-chanRes).errMsg
}

//DeletingTaskByID удаление задачи по ее ID
func (tst *TemporaryStorageType) DeletingTaskByID(appTaskID string) {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{commanChannelTaskStorage: commanChannelTaskStorage{appTaskID: appTaskID}}
	<-chanRes
}

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ НАЙДЕННОЙ ИНФОРМАЦИИ ***/

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ПАРАМЕТРОВ ПРИЛОЖЕНИЯ ***/

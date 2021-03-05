package memorytemporarystoragecommoninformation

import (
	"fmt"
)

/*** МЕТОДЫ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ЗАДАЧ ***/

/*
type TemporaryStorageTaskType struct {
	TaskGenerator        string
	ClientID             string
	ClientName           string
	AdditionalClientName string
	DateTaskCreated      time.Time
	DateTaskModification time.Time
	TaskStatus           string
	Section              string
	TaskParameters       interface{}
}
*/

//AddNewTask метод добавляющий новую задачу в хранилище задач
func (tst *TemporaryStorageType) AddNewTask(taskInfo *TemporaryStorageTaskType) (string, error) {
	fmt.Println("func 'AddNewTask', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{ /*

					НЕДОДЕЛАЛ
			надо дополнительно описать тип канала для передачи информации
		*/}
	result := <-chanRes

	return result.appTaskID, result.errMsg
}

//GetTaskByID получить информацию о задаче по ее ID
func (tst *TemporaryStorageType) GetTaskByID(taskID string) error {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{}

	return (<-chanRes).errMsg
}

//ChangeDateTaskModification изменяет время модификации задачи
func (tst *TemporaryStorageType) ChangeDateTaskModification(taskID string) error {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{}

	return (<-chanRes).errMsg
}

//ChangeTaskStatus изменить статус выполнения задачи
func (tst *TemporaryStorageType) ChangeTaskStatus(taskID string) error {
	fmt.Println("func 'GetTaskByID', START...")

	chanRes := make(chan channelResponseTaskStorage)
	defer func() {
		close(chanRes)
	}()

	tst.chanReqTaskStorage <- channelRequestTaskStorage{}

	return (<-chanRes).errMsg
}

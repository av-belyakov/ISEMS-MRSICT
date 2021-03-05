package memorytemporarystoragecommoninformation

import (
	"fmt"
)

/*** ФУНКЦИИ ОТНОСЯЩИЕСЯ К ХРАНИЛИЩУ ЗАДАЧ ***/

//addNewTask добавляет новую задачу в хранилище задач
func (tst *TemporaryStorageType) addNewTask() {
	fmt.Println("addNewTask добавляет новую задачу в хранилище задач")

	//генерируем новый UUID и добавляем задачу в хранилище
	//UUID возвращаем пользователю
}

//getOneTask возвращает информацию об одной задаче найденной по её ID
func (tst *TemporaryStorageType) getOneTask() {
	fmt.Println("getOneTask возвращает информацию об одной задаче найденной по её ID")
}

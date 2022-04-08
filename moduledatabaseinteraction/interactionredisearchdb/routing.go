package interactionredisearchdb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
)

//Routing обеспечивает маршрутизацию информации между базой данных RedisearchDB и ядром приложения
// chanOutput - канал для ПЕРЕДАЧИ данных ядру приложения
// cdrdb - дескриптор соединения с БД Redishearch
// tst - общее хранилище временной информации
// chanInput - канал для ПРИЕМА данных, приходящих от ядра приложения
func Routing(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	cdrdb ConnectionDescriptorRedisearchDB,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	chanInput <-chan datamodels.ModuleDataBaseInteractionChannel) {

	fmt.Printf("func 'Routing' Redisearch DB")

	for data := range chanInput {
		switch data.Section {
		case "handling insert index":
			go wrapperFuncHandlingInsertIndex(chanOutput, data, tst, cdrdb)
		case "handling select index":
			go wrapperFuncHandlingSelectIndex(chanOutput, data, tst, cdrdb)
		}
	}
}

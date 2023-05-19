package interactionredisearchdb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulelogginginformationerrors"

	"github.com/RediSearch/redisearch-go/redisearch"
)

// ChannelsRedisearchInteraction содержит описание каналов для взаимодействия с БД Redisearch
// InputModule - канал для ПРИЕМА данных, приходящих от ядра приложения
// OutputModule - канал для ПЕРЕДАЧИ данных ядру приложения
type ChannelsRedisearchInteraction struct {
	InputModule, OutputModule chan datamodels.ModuleDataBaseInteractionChannel
}

// ConnectionDescriptorRedisearchDB дескриптор соединения с БД RedisearchB
// Connection - дескриптор соединения
// CTX - контекст переносит крайний срок, сигнал отмены и другие значения через границы API
type ConnectionDescriptorRedisearchDB struct {
	Connection *redisearch.Client
}

var crdbi ChannelsRedisearchInteraction
var cdrdb ConnectionDescriptorRedisearchDB

func init() {
	crdbi = ChannelsRedisearchInteraction{
		InputModule:  make(chan datamodels.ModuleDataBaseInteractionChannel),
		OutputModule: make(chan datamodels.ModuleDataBaseInteractionChannel),
	}

	//	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	cdrdb = ConnectionDescriptorRedisearchDB{}
}

func InteractionRedisearchDB(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	rdbs *datamodels.RedisearchDBSettings,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (ChannelsRedisearchInteraction, error) {

	fmt.Println("func 'InteractionRedisearchDB', START...")

	if err := cdrdb.CreateConnection(rdbs); err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "InteractionRedisearchDB",
		}

		return crdbi, err
	}

	/*
		Сделал:
		1. добавил в HandlerAssignmentsModuleAPIRequestProcessing Section "handling stix object" на
		 ряду с отправкой на добавление списка STIX объектов к БД MongoDB еще и отправку запроса
		 на создание индексов в БД RedisearchDB. При чем отправка запроса к RedisearchDB будет
		 выполнятся первой, а обработка запросов в обоих БД будет вестись параллельно. Но я думаю
		 что RedisearchDB отработает быстрее так как ей не надо выполнят допонительные запросы из
		 БД и сравнение изменений в объектах. Так что за очередность обработки не стоит беспокоится.
		2. Добавил функцию handlerDataBaseInteraction в mailHandlerModuleScheduler.
		3. Написал функцию Routing состоящую из двух оберток wrapperFuncHandlingInsertIndex и написал
		функцию wrapperFuncHandlingInsertIndex в wrapperRouteRequest. Там выполняется полная обработка
		списка ElementSTIXObject и добавление индексов а БД RedisearchDB. Добавление индексов в БД
		RedisearchDB проверялось в тестах (успешно), однако в совокупности вся сепочка обработки STIX
		объектов не проверялась.
		4. Изменил версию приложения так как добавился новый функционал.

		Что нужно сделать:
		+ 1. Обновить версию MRSICa но не в ДОКЕРАХ, а просто на хостовой тестовой системе и проверить
		всю цепочку обновления хотя бы одного STIX объекта. (проверил обнавление одного объекта, все работает)
		2. Продумать и написать методы поиска индексов по заданным параметрам
		в БД Redisearch, а также методы и алгоритмы передачи информации в
		модуль который работает с БД Redisearch и приема из него
		3. Продумать и написать метод индексации тех STIX объектов которые,
		возможно ранее не были проиндексированны или оказались не проиндексированны
		по причине перезапуска БД Rediserch
	*/

	go Routing(crdbi.OutputModule, cdrdb, tst, crdbi.InputModule)

	return crdbi, nil
}

func (cdrdb *ConnectionDescriptorRedisearchDB) CreateConnection(mdbs *datamodels.RedisearchDBSettings) error {
	fmt.Println("func 'CreateConnection', Redisearch, START...")
	fmt.Printf("RedisearchDBSettings: %v\n", mdbs)

	cdrdb.Connection = redisearch.NewClient(fmt.Sprintf("%v:%v", mdbs.Host, mdbs.Port), "isems-index")
	if _, err := cdrdb.Connection.Info(); err == nil {
		return nil
	}

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("type")).
		AddField(redisearch.NewTextField("name")).
		AddField(redisearch.NewTextField("description")).
		//физический адрес
		AddField(redisearch.NewTextField("street_address")).
		//результат классификации или имя, присвоенное экземпляру вредоносного ПО инструментом анализа (сканером)
		// используется в STIX объектах MalwareAnalysis
		AddField(redisearch.NewTextField("result_name")).
		//краткое изложение содержания записки используется в STIX объектах Node
		AddField(redisearch.NewTextField("abstract")).
		//основное содержание записки используется в STIX объектах Node
		AddField(redisearch.NewTextField("content")).
		AddField(redisearch.NewTextField("url")).
		//параметр value может содержать в себе сетевое доменное имя,
		// email адрес, ip адрес, url в STIX объектах DomainName, EmailAddress,
		// IPv4Address, IPv6Address, URL
		AddField(redisearch.NewTextField("value"))

	if err := cdrdb.Connection.CreateIndex(sc); err == nil {
		return nil
	}

	return fmt.Errorf("error connecting to the Research database or error creating indexes")
}

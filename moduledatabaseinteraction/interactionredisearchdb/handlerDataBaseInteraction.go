package interactionredisearchdb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulelogginginformationerrors"

	"github.com/RediSearch/redisearch-go/redisearch"
)

//ChannelsRedisearchInteraction содержит описание каналов для взаимодействия с БД Redisearch
// InputModule - канал для ПРИЕМА данных, приходящих от ядра приложения
// OutputModule - канал для ПЕРЕДАЧИ данных ядру приложения
type ChannelsRedisearchInteraction struct {
	InputModule, OutputModule chan datamodels.ModuleDataBaseInteractionChannel
}

//ConnectionDescriptorRedisearchDB дескриптор соединения с БД RedisearchB
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
	mdbs *datamodels.RedisearchDBSettings,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (ChannelsRedisearchInteraction, error) {

	fmt.Println("func 'InteractionRedisearchDB', START...")

	if err := cdrdb.CreateConnection(mdbs); err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "InteractionRedisearchDB",
		}

		return crdbi, err
	}

	/*
	   1. Добавить функцию handlerDataBaseInteraction в
	   mailHandlerModuleScheduler
	   2. Дописать функцию Routing
	   3. Написать метод для каждого STIX объекта который соответствовал бы
	   определенному интерфейсу и обеспечивал индексацию данных из STIX
	   объекта для их последующего сохранения в БД Redisearch. Кстати,
	   надо обращать внимание на пустые поля STIX объекта или поля
	   содержащие значение null. Если все индексируемые поля для
	   данного объекта остаются пустыми или содержат значение null
	   то такой объект не должен индексироватся. А также не должны
	   индексироватся поля содержащие такие значения.
	   4. Продумать очередность обработки данных по STIX объектам при
	   их сохранении в БД MongoDB и Redisearch
	   5. Продумать и написать методы поиска индексов по заданным параметрам
	   в БД Redisearch, а также методы и алгоритмы передачи информации в
	   модуль который работает с БД Redisearch и приема из него

	   Все STIX объекты соответствуют интерфейсу IndexingSTIXObject то есть
	   имеют метод GeneratingDataForIndexing, теперь нужно сделать сначало
	   тестовую функцию для перебора среза типа ElementSTIXObject в котором
	   свойство Data соответствует интерфейсу HandlerSTIXObject в составе которого
	   есть и интерфейс IndexingSTIXObject
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

package interactionredisearchdb

import (
	"fmt"
	"log"

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

	if err := cdrdb.CreateConnection(rdbs); err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "InteractionRedisearchDB",
		}

		return crdbi, err
	}

	go Routing(crdbi.OutputModule, cdrdb, tst, crdbi.InputModule)

	return crdbi, nil
}

func (cdrdb *ConnectionDescriptorRedisearchDB) CreateConnection(mdbs *datamodels.RedisearchDBSettings) error {
	cdrdb.Connection = redisearch.NewClient(fmt.Sprintf("%v:%v", mdbs.Host, mdbs.Port), "isems-index")
	if _, err := cdrdb.Connection.Info(); err != nil {
		return err
	}

	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("type")).
		AddField(redisearch.NewTextField("name")).
		AddField(redisearch.NewTextField("description")).
		//физический адрес
		AddField(redisearch.NewTextField("street_address")).
		//результат классификации или имя, присвоенное экземпляру вредоносного ПО инструментом анализа (сканером)
		// используется в STIX объектах MalwareAnalysis
		//AddField(redisearch.NewTextField("result_name")).
		//краткое изложение содержания записки используется в STIX объектах Node
		AddField(redisearch.NewTextField("abstract")).
		//основное содержание записки используется в STIX объектах Node
		AddField(redisearch.NewTextField("content")).
		//AddField(redisearch.NewTextField("url")).
		//параметр value может содержать в себе сетевое доменное имя,
		// email адрес, ip адрес, url в STIX объектах DomainName, EmailAddress,
		// IPv4Address, IPv6Address, URL
		AddField(redisearch.NewTextField("value"))

	log.Printf("Create connection with Redisearch (%s:%d)\n", mdbs.Host, mdbs.Port)

	if err := cdrdb.Connection.CreateIndex(sc); err == nil {
		return err
	}

	return nil
}

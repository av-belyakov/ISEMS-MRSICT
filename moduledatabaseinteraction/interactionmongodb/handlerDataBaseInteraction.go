package interactionmongodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"
	"ISEMS-MRSICT/modulelogginginformationerrors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//ChannelsMongoDBInteraction содержит описание каналов для взаимодействия с БД MongoDB
// InputModule - канал для ПРИЕМА данных, приходящих от ядра приложения
// OutputModule - канал для ПЕРЕДАЧИ данных ядру приложения
type ChannelsMongoDBInteraction struct {
	InputModule, OutputModule chan datamodels.ModuleDataBaseInteractionChannel
}

//ConnectionDescriptorMongoDB дескриптор соединения с БД MongoDB
// Connection - дескриптор соединения
// CTX - контекст переносит крайний срок, сигнал отмены и другие значения через границы API
type ConnectionDescriptorMongoDB struct {
	connection *mongo.Client
	ctx        context.Context
	ctxCancel  context.CancelFunc
}

var cmdbi ChannelsMongoDBInteraction
var cdmdb ConnectionDescriptorMongoDB

func init() {
	cmdbi = ChannelsMongoDBInteraction{
		InputModule:  make(chan datamodels.ModuleDataBaseInteractionChannel),
		OutputModule: make(chan datamodels.ModuleDataBaseInteractionChannel),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)

	cdmdb = ConnectionDescriptorMongoDB{
		ctx:       ctx,
		ctxCancel: cancel,
	}
}

//InteractionMongoDB модуль взаимодействия с БД MongoDB
func InteractionMongoDB(
	chanSaveLog chan<- modulelogginginformationerrors.LogMessageType,
	mdbs *datamodels.MongoDBSettings,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType) (ChannelsMongoDBInteraction, error) {

	fmt.Println("func 'InteractionMongoDB', START...")
	fmt.Printf("func 'InteractionMongoDB', settings db: '%v'\n", mdbs)

	defer func() {
		cdmdb.ctxCancel()
	}()

	//подключаемся к базе данных MongoDB
	if err := cdmdb.createConnection(mdbs); err != nil {
		chanSaveLog <- modulelogginginformationerrors.LogMessageType{
			TypeMessage: "error",
			Description: fmt.Sprint(err),
			FuncName:    "InteractionMongoDB",
		}

		return cmdbi, err
	}

	//инициализируем маршрутизатор запросов
	go Routing(chanSaveLog, cmdbi.OutputModule, mdbs.NameDB, cdmdb, tst, cmdbi.InputModule)

	return cmdbi, nil
}

func (cdmongodb ConnectionDescriptorMongoDB) createConnection(mdbs *datamodels.MongoDBSettings) error {

	fmt.Println("func 'createConnection' ----------- START ----------")

	clientOption := options.Client()
	clientOption.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    mdbs.NameDB,
		Username:      mdbs.User,
		Password:      mdbs.Password,
	})

	client, err := mongo.NewClient(clientOption.ApplyURI("mongodb://" + mdbs.Host + ":" + strconv.Itoa(mdbs.Port) + "/" + mdbs.NameDB))
	if err != nil {
		return err
	}

	client.Connect(cdmongodb.ctx)

	if err = client.Ping(cdmongodb.ctx, readpref.Primary()); err != nil {
		return err
	}

	cdmongodb.connection = client

	fmt.Println("func 'createConnection' ----------- END ----------")

	return nil
}

//GetConnection возвращает дескриптор соединения с БД MongoDB
func (cdmongodb ConnectionDescriptorMongoDB) GetConnection() (*mongo.Client, context.Context) {
	return cdmongodb.connection, cdmongodb.ctx
}

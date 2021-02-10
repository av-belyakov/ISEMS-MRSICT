package interactionmongodb

import (
	"fmt"

	"ISEMS-MRSICT/datamodels"
)

//ChansMongoDBInteraction содержит описание каналов для взаимодействия с БД MongoDB
type ChansMongoDBInteraction struct {
	InputModule, OutputModule chan interface{} //ПОКА interface{} потому что тип не определил
}

//InteractionMongoDB модуль взаимодействия с БД MongoDB
func InteractionMongoDB(sdb *datamodels.MongoDBSettings) (ChansMongoDBInteraction, error) {
	fmt.Println("func 'InteractionMongoDB', START...")
	fmt.Printf("func 'InteractionMongoDB', settings db: '%v'\n", sdb)

	chansInteraction := ChansMongoDBInteraction{
		InputModule:  make(chan interface{}),
		OutputModule: make(chan interface{}),
	}

	return chansInteraction, nil
}

/*
	//connectToDB устанавливает соединение с БД
func connectToDB(ctx context.Context, appc *configure.AppConfig) (*mongo.Client, error) {
	host := appc.ConnectionDB.Host
	port := appc.ConnectionDB.Port

	opts := options.Client()
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    appc.ConnectionDB.NameDB,
		Username:      appc.ConnectionDB.User,
		Password:      appc.ConnectionDB.Password,
	})

	//client, err := mongo.NewClientWithOptions("mongodb://"+host+":"+strconv.Itoa(port)+"/"+appc.ConnectionDB.NameDB, opts)
	client, err := mongo.NewClient(opts.ApplyURI("mongodb://" + host + ":" + strconv.Itoa(port) + "/" + appc.ConnectionDB.NameDB))
	if err != nil {
		return nil, err
	}

	client.Connect(ctx)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
*/

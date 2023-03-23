package db

import (
	"context"
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"mongo-test/config"
	"time"
)

func InitMongo() {
	dataSource := config.Config.Mongo
	var uri string
	var err error

	if dataSource.Username == "" && dataSource.Password == "" {
		//if dataSource.Username != "" && dataSource.Password != "" {
		//uri = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		//	dataSource.Username, dataSource.Password, dataSource.Host, dataSource.Schema)
		uri = fmt.Sprintf("mongodb://%s:%d/%s?w=majority",
			dataSource.Host, dataSource.Port, dataSource.Schema)

		Mongo, err = mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			log.Printf("mongo connect failed! uri: %s error: %v", uri, err)
			log.Fatal(err)
		}

		// following is for mgm
		opt := options.Client().ApplyURI(uri)
		err = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 30 * time.Second}, config.Config.Mongo.Schema, opt)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		err = Mongo.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		//Ping the primary
		err = Mongo.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}
	}

}

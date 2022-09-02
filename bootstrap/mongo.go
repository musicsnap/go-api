package bootstrap

import (
	"context"
	"go-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

var client *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conf := config.GetConfig()
	mongoMap := conf.GetStringMap("mongodb")
	credential := options.Credential{
		Username: mongoMap["username"].(string),
		Password: mongoMap["password"].(string),
	}
	var err error
	uri := "mongodb://" + mongoMap["host"].(string) + ":" + strconv.Itoa(mongoMap["port"].(int))
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		log.Println(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	//defer func() {
	//	if err := client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
}

func GetMongoClient() *mongo.Client {
	return client
}

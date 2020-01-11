package utils

import (
	"context"
	"fmt"

	"github.com/kamicloud/mahjong-science-server/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

const maxConnectTimes = 10

// GetClient 获取连接
func GetClient() error {
	if client != nil && client.Ping(context.TODO(), readpref.Primary()) == nil {
		return nil
	}

	i := 0
	for {
		i++
		c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(app.Config.Mongo))

		if err == nil {
			client = c
			return nil
		}
		if i > maxConnectTimes {
			fmt.Println(err)
		}
	}
}

// GetDatabase 获取数据库
func GetDatabase(databaseName string) *mongo.Database {
	GetClient()

	return client.Database(databaseName)
}

// GetCollection 获取数据集
func GetCollection(databaseName string, collectionName string) *mongo.Collection {
	database := GetDatabase(databaseName)

	if database == nil {
		return nil
	}

	return database.Collection(collectionName)
}

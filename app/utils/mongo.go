package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// GetClient 获取连接
func GetClient() (*mongo.Client, context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100000*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	return client, ctx, err
}

// GetDatabase 获取数据库
func GetDatabase(databaseName string) (*mongo.Database, context.Context) {
	client, ctx, err := GetClient()

	if err != nil {

		return nil, nil
	}

	return client.Database(databaseName), ctx
}

// GetCollection 获取数据集
func GetCollection(databaseName string, collectionName string) (*mongo.Collection, context.Context) {
	database, ctx := GetDatabase(databaseName)

	if database == nil {
		return nil, nil
	}

	return database.Collection(collectionName), ctx
}

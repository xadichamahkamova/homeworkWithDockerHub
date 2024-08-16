package pkg

import (
	"context"
	"fmt"
	"service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client     mongo.Client
	Collection mongo.Collection
}
var ctx = context.Background()

func NewConnection(cfg config.Config) (*MongoDB, error) {

	uri := fmt.Sprintf("mongodb://%s:%s", cfg.Mongo.Host, cfg.Mongo.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	mycoll := client.Database(cfg.Mongo.Database).Collection(cfg.Mongo.Collection)
	return &MongoDB{
		Client: *client,
		Collection: *mycoll,
	}, nil
}
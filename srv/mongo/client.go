package mongo

import (
	"context"
	"log"
	"time"
)
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"

type Config struct {
	DSN         string
	ConnTimeout time.Duration
	Database    string
}

type Client struct {
	Mongo  *mongo.Client
	Config *Config
}

func NewClient(c *Config) *mongo.Client {
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(c.DSN),
		options.Client().SetConnectTimeout(c.ConnTimeout),
		options.Client().SetServerSelectionTimeout(c.ConnTimeout),
	)
	if err != nil {
		log.Fatal(err)
	}

	/*defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()*/

	return client
}

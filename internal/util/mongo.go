package util

import (
	"context"
	"github.com/gofiber/fiber/v3/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoDB *mongo.Database

func init() {
	uri := os.Getenv("LUMI_MONGO_DSN")
	if uri == "" {
		log.Fatal("未设置LUMI_MONGO_DSN环境变量")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	MongoDB = client.Database(os.Getenv("LUMI_MONGO_DB"))

}

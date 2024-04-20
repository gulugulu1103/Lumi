package database

import (
	"context"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoDB *mongo.Database

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("无法加载.env文件")
	}

	uri := os.Getenv("LUMI_MONGO_DSN")
	if uri == "" {
		log.Fatal("未设置LUMI_MONGO_DSN环境变量")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("无法连接MongoDB, ", err.Error())
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("无法ping MongoDB, ", err.Error())
		return
	}

	MongoDB = client.Database(os.Getenv("LUMI_MONGO_DB"))
	log.Info("成功连接MongoDB")
}

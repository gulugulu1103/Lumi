package logger

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"os"
	"strings"
	"time"
)

const DEFAULT_LOG_DB_NAME = "lumi_log"

var Log *zap.Logger

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("读取.env文件失败")
		return
	}
	Log = newLogger()
}

// newLogger 创建并配置一个 zap.Logger 实例
func newLogger() *zap.Logger {
	// 配置 zap
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.MessageKey = "message"

	// 获取mongoDSN
	if dsn := os.Getenv("LUMI_MONGO_DSN"); dsn == "" {
		log.Fatal("LUMI_MONGO_DSN不能为空")
	}
	if err := zap.RegisterSink("mongodb", newMongoSink); err != nil {
		log.Panicf("zap注册mongo失败: %v", err)
		panic(err)
	}
	config.OutputPaths = []string{"stderr", os.Getenv("LUMI_MONGO_DSN")}

	// 初始化 logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

type MongoSink struct {
	zap.Sink
	Database       *mongo.Database
	DatabaseName   string
	CollectionName string
}

func (m *MongoSink) Close() error {
	// 关闭数据库连接
	err := m.Database.Client().Disconnect(context.TODO())
	panic(err)
	return err
}

func (m *MongoSink) Write(p []byte) (n int, err error) {
	var document map[string]interface{}
	if err = sonic.Unmarshal(p, &document); err != nil {
		log.Fatal("日志json反序列化失败: ", err.Error())
		return 0, err
	}
	one, err := m.Database.Collection(m.CollectionName).InsertOne(nil, document)
	if err != nil || one == nil {
		log.Panicf("日志写入mongo失败", err.Error())
		return 0, err
	}
	return len(p), nil
}

func (m *MongoSink) Sync() error {
	return nil
}

func newMongoSink(url *url.URL) (zap.Sink, error) {
	var sink MongoSink
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url.String()))
	if err != nil {
		log.Fatal("无法连接MongoDB, ", err.Error())
		panic(err)
	}

	// 获取数据库名
	if sink.DatabaseName = os.Getenv("LUMI_MONGO_LOG_DB"); sink.DatabaseName == "" {
		sink.DatabaseName = DEFAULT_LOG_DB_NAME
	}

	// 获取数据库
	sink.Database = client.Database(sink.DatabaseName)

	// ping
	err = sink.Database.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("无法ping MongoDB, ", err.Error())
		return nil, err
	}

	// 创建Collection
	sink.CollectionName = strings.ReplaceAll(time.Now().Local().Format(time.DateTime), " ", "/")

	// 创建索引
	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "time", Value: 1}}},
		{Keys: bson.D{{Key: "level", Value: 1}}},
	}
	_, err = sink.Database.Collection(sink.CollectionName).Indexes().CreateMany(context.TODO(), indexes)
	if err != nil {
		log.Fatal("mongo_logger创建索引失败", err.Error())
	}

	return &sink, err
}

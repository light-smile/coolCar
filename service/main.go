package main

import (
	"context"
	"coolCar/service/auth/api/auth"
	"coolCar/service/auth/api/gen/v1"
	"coolCar/service/auth/token"
	"coolCar/service/auth/wechat"
	"coolCar/service/shared"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	logger, err := newZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := "mongodb://localhost:27017/blogDB"
	opt := new(options.ClientOptions)
	// 设置最大连接数量
	opt = opt.SetMaxPoolSize(uint64(10))
	// 设置连接超时时间 5000 毫秒
	du, _ := time.ParseDuration("5000")
	opt = opt.SetConnectTimeout(du)
	// 设置连接的空闲时间 毫秒
	mt, _ := time.ParseDuration("5000")
	opt = opt.SetMaxConnIdleTime(mt)
	// 开启驱动
	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), opt)
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	pkFile, err := os.Open("service/auth/token/private.key")
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx123123",
			AppSecret: "",
		},
		Mongo:          shared.NewMongo(MongoClient.Database("coolcar")),
		Logger:         logger,
		TokenExpire:    2 * time.Hour,
		TokenGenerator: token.NewJWTTokenGen("coolcar/auth", privKey),
	})

	err = s.Serve(lis)
	logger.Fatal("cannot server", zap.Error(err))
}

// 自定义zap
func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}

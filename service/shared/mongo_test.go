package shared

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestResolveAccountID(t *testing.T) {
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
		fmt.Errorf("connect db failed: %v", err)
	}
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Errorf("ping failed: %v", err)
	}

	m := NewMongo(MongoClient.Database("coolcar"))
	//m.col.InsertMany(ctx, []interface{
	//	bson.M{
	//		mgo.IDField: mustObjID("62b0713ff40782827b054f00"),
	//		openIDField: "openid_1",
	//	}
	//})
	if err != nil {
		t.Fatalf("cannot insert initial values: %v", err)
	}

	id, err := m.ResolveAccoutID(ctx, "123")
	if err != nil {
		t.Errorf("faild resolve account id for 123: %v", err)
	} else {
		want := "62b0713ff40782827b054f00"
		if id != want {
			t.Errorf("resolve account id: want: %q, got: %q", want, id)
		}
	}

}

func mustObjID(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return objID
}

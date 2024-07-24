package service

import (
	"context"
	"github.com/dayueba/crontabd/internal/worker/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Data struct {
	mongoClient  *mongo.Client
	etcdClientv3 *clientv3.Client
}

func NewData(conf *config.Config) (*Data, func(), error) {
	var (
		mongoClient  *mongo.Client
		etcdClientv3 *clientv3.Client
		err          error
	)

	// 建立mongodb连接
	if mongoClient, err = mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(conf.MongodbUri),
		options.Client().SetTimeout(time.Duration(conf.MongodbConnectTimeout)*time.Millisecond)); err != nil {
		return nil, nil, err
	}

	// 初始化配置
	config := clientv3.Config{
		Endpoints:   conf.EtcdEndpoints,                                     // 集群地址
		DialTimeout: time.Duration(conf.EtcdDialTimeout) * time.Millisecond, // 连接超时
	}

	// 建立连接
	if etcdClientv3, err = clientv3.New(config); err != nil {
		return nil, nil, err
	}

	return &Data{mongoClient: mongoClient, etcdClientv3: etcdClientv3}, func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}, nil

}

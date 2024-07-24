package config

// 程序配置
type Config struct {
	EtcdEndpoints         []string `json:"etcdEndpoints"`
	EtcdDialTimeout       int      `json:"etcdDialTimeout"`
	MongodbUri            string   `json:"mongodbUri"`
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"`
	JobLogBatchSize       int      `json:"jobLogBatchSize"`
	JobLogCommitTimeout   int      `json:"jobLogCommitTimeout"`
}

func DefaultConfig() *Config {
	return &Config{
		EtcdEndpoints:         []string{"127.0.0.1:2379"},
		EtcdDialTimeout:       2000,
		MongodbUri:            "mongodb://admin:123456@127.0.0.1:27017",
		MongodbConnectTimeout: 2000,
		JobLogBatchSize:       0,
		JobLogCommitTimeout:   0,
	}
}

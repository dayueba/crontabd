package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiPort               int      `yaml:"apiPort"`
	EtcdEndpoints         []string `yaml:"etcdEndpoints"`
	EtcdDialTimeout       int      `yaml:"etcdDialTimeout"`
	MongodbUri            string   `yaml:"mongodbUri"`
	MongodbConnectTimeout int      `yaml:"mongodbConnectTimeout"`
}

func InitConfig(filename string) (*Config, error) {
	conf := &Config{}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	return conf, err
}

func DefaultConfig() *Config {
	return &Config{
		ApiPort:               7200,
		EtcdEndpoints:         []string{"127.0.0.1:2379"},
		EtcdDialTimeout:       2000,
		MongodbUri:            "mongodb://admin:123456@127.0.0.1:27017",
		MongodbConnectTimeout: 2000,
	}
}

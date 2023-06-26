package main

import (
	"flag"
	"github.com/dayueba/crontabd/internal/master"
	"github.com/dayueba/crontabd/internal/master/config"
	"github.com/dayueba/crontabd/internal/master/service"
	"log"

	_ "go.uber.org/automaxprocs"
)

var (
	confFile string // 配置文件路径
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./example-config.yaml", "config file")
	flag.Parse()
}

func main() {
	initArgs()

	var err error
	var conf *config.Config
	//if conf, err = master.InitConfig(confFile); err != nil {
	//	log.Fatalln(err)
	//}
	conf = config.DefaultConfig()

	data, cleanup, err := service.NewData(conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanup()

	workerService := service.NewWorkerService(data)
	logService := service.NewLogService(data)
	jobService := service.NewJobService(data)

	apiServer := master.NewApiServer(conf, jobService, workerService, logService)

	if err = apiServer.Run(); err != nil {
		log.Fatalln(err)
	}
}

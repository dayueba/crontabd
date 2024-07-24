package main

import (
	"github.com/dayueba/crontabd/internal/worker/config"
	"github.com/dayueba/crontabd/internal/worker/service"
	_ "go.uber.org/automaxprocs"
	"log"
	"time"
)

func main() {

	conf := config.DefaultConfig()

	data, cleanup, err := service.NewData(conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer cleanup()

	err = service.InitRegister(data)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Minute * 60)
}

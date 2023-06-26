package service

import (
	"github.com/dayueba/crontabd/internal/master/config"
	"testing"
)

func TestLogMgr_ListLog(t *testing.T) {
	conf := config.DefaultConfig()
	data, cleanup, err := NewData(conf)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	logService := NewLogService(data)

	//docs := []interface{}{
	//	bson.D{{"jobName", "Alice"}},
	//	bson.D{{"jobName", "Bob"}},
	//}
	//opts := options.InsertMany().SetOrdered(false)
	//res, err := GlobalLogMgr.logCollection.InsertMany(context.TODO(), docs, opts)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//t.Log(res)

	list, err := logService.ListLog("Alice", 0, 20)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range list {
		t.Log(item)
	}
}

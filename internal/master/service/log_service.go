package service

import (
	"context"
	"github.com/dayueba/crontabd/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb日志管理
type LogService struct {
	client        *mongo.Client
	logCollection *mongo.Collection
}

func NewLogService(data *Data) *LogService {
	return &LogService{
		client:        data.mongoClient,
		logCollection: data.mongoClient.Database("cron").Collection("log"),
	}
}

// 查看任务日志
func (logService *LogService) ListLog(name string, skip int64, limit int64) ([]*internal.JobLog, error) {
	// example: https://raw.githubusercontent.com/mongodb/docs-golang/master/source/includes/usage-examples/code-snippets/find.go
	opts := options.Find().SetSort(bson.D{{"startTime", -1}}).SetSkip(skip).SetLimit(limit)
	cursor, err := logService.logCollection.Find(context.Background(), bson.D{{"jobName", name}}, opts)
	if err != nil {
		return nil, err
	}

	var results []*internal.JobLog
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

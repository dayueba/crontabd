package service

import (
	"context"
	"github.com/dayueba/crontabd/internal"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
)

// /cron/workers/
type WorkerService struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

// 获取在线worker列表
func (workerService *WorkerService) ListWorkers() (workerArr []string, err error) {
	var (
		getResp  *clientv3.GetResponse
		kv       *mvccpb.KeyValue
		workerIP string
	)

	// 初始化数组
	workerArr = make([]string, 0)

	// 获取目录下所有Kv
	if getResp, err = workerService.kv.Get(context.Background(), internal.JOB_WORKER_DIR, clientv3.WithPrefix()); err != nil {
		return
	}

	// 解析每个节点的IP
	for _, kv = range getResp.Kvs {
		// kv.Key : /cron/workers/192.168.2.1
		workerIP = ExtractWorkerIP(string(kv.Key))
		workerArr = append(workerArr, workerIP)
	}
	return
}

func NewWorkerService(data *Data) *WorkerService {
	var (
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
	)
	client = data.etcdClientv3

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	return &WorkerService{
		client: client,
		kv:     kv,
		lease:  lease,
	}
}

// 提取worker的IP
func ExtractWorkerIP(regKey string) string {
	return strings.TrimPrefix(regKey, internal.JOB_WORKER_DIR)
}

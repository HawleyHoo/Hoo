/*
@Time : 2019-07-03 14:42
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type ServerMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

const WorkerDir = "/admin/services/"

var (
	G_serverMgr *ServerMgr
)

func main() {
	err := InitServerMgr()
	if err != nil {
		fmt.Println(err)
		return
	}
	sers, err := G_serverMgr.ListServices()
	fmt.Println(sers, err)
}

func InitServerMgr() error {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // 集群地址
		DialTimeout: time.Duration(5 * time.Second),
	}

	client, err := clientv3.New(config)
	if err != nil {
		return err
	}

	// 得到KV和Lease的API子集
	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)

	G_serverMgr = &ServerMgr{
		client: client,
		kv:     kv,
		lease:  lease,
	}
	return nil
}

func (mgr *ServerMgr) ListServices() ([]string, error) {
	servicesArr := make([]string, 0)

	resp, err := G_serverMgr.kv.Get(context.TODO(), WorkerDir, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	for _, kv := range resp.Kvs {
		//workerIP := ExtractWorkerIP(string(kv.Key))
		workerIP := string(kv.Key)
		fmt.Println("key:", string(kv.Key))
		fmt.Println("val:", string(kv.Value))
		if len(workerIP) > 0 {
			servicesArr = append(servicesArr, workerIP)
		} else {
			fmt.Println("err kv:%s", kv.Key)
		}
	}
	return servicesArr, nil
}

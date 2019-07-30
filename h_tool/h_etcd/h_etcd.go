/*
@Time : 2019-07-03 12:24
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

func initclient() {
	//clientv3.Config{}
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)
}

func main() {
	//clientv3.Config{}
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), config.DialTimeout)
	defer cancel()
	_, err = client.Status(timeoutCtx, config.Endpoints[0])
	if err != nil {
		fmt.Println("error checking etcd status:", err)
		return
		//return nil, errors.(err, "error checking etcd status: %v", err)
	}

	// 用于读写etcd的键值对
	kv := clientv3.NewKV(client)

	if putResp, err := kv.Put(context.TODO(), "/cron/jobs/job1", "bye", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil { // 打印hello
			fmt.Println("PrevValue:", string(putResp.PrevKv.Value))
		}
	}

	if getResp, err := kv.Get(context.TODO(), "/cron/jobs/job1" /*clientv3.WithCountOnly()*/); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("kvs:", getResp.Kvs, " count:", getResp.Count)
	}

	// 写入另外一个Job
	kv.Put(context.TODO(), "/cron/jobs/job2", "{...}")

	// 读取/cron/jobs/为前缀的所有key
	if getResp, err := kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else { // 获取成功, 我们遍历所有的kvs
		fmt.Println(getResp.Kvs)
	}

	/*{
		// 删除KV
		if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job1", clientv3.WithFromKey(), clientv3.WithLimit(2)); err != nil {
			fmt.Println(err)
			return
		}

		// 被删除之前的value是什么
		if len(delResp.PrevKvs) != 0 {
			for _, kvpair = range delResp.PrevKvs {
				fmt.Println("删除了:", string(kvpair.Key), string(kvpair.Value))
			}
		}
	}*/

	fmt.Println("sucess.")
}

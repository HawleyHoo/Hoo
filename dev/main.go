/*
@Time : 2019-06-20 12:24
@Author : Hoo
@Project: Hoo
*/

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"os/signal"
	"syscall"

	//"github.com/coreos/etcd/clientv3"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("index server.", time.Now())
	io.WriteString(w, "<h1>This is Index Page!  dockerfile </h1>")
}

func register(w http.ResponseWriter, r *http.Request) {
	log.Println("register server.")
	err := InitRegister()
	if err != nil {
		log.Println(err.Error())
	}
	io.WriteString(w, "<h1>register server.! </h1>")
}

func initSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		removeRegister()
		log.Printf("receive a signal:%v.  removeRegister \n", sig)
		os.Exit(-1)
	}()
}

func init() {
	initSignals()
}

func main() {
	fmt.Println("pid:", os.Getpid(), "ppid:", os.Getppid())
	//clientv3.Config{}

	err := InitRegister()
	if err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)

	port := "8080"
	log.Println("listen server: localhost:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicln(err.Error())
	}

}

// 注册节点到etcd： /cron/workers/IP地址
type Register struct {
	client  *clientv3.Client
	kv      clientv3.KV
	lease   clientv3.Lease
	localIP string // 本机IP
}

var (
	G_register *Register
)

// 获取本机网卡IP
func getLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	err = errors.New("没有找到网卡IP")
	return
}

type NodeServer struct {
	Status  int
	Pid     int
	Name    string
	Message string
	IPAddr  string
}

const WORKER_DIR = "/admin/services/"

// 注册到/cron/workers/IP, 并自动续租
func (register *Register) keepOnline() {
	var (
		nodes          NodeServer
		vl             []byte
		regKey, name   string
		leaseGrantResp *clientv3.LeaseGrantResponse
		err            error
		keepAliveChan  <-chan *clientv3.LeaseKeepAliveResponse
		keepAliveResp  *clientv3.LeaseKeepAliveResponse
		cancelCtx      context.Context
		cancelFunc     context.CancelFunc
	)

	for {
		// 服务名字
		name = "testmain"

		// 注册路径
		regKey = WORKER_DIR + register.localIP + ":" + name

		cancelFunc = nil

		// 创建租约
		if leaseGrantResp, err = register.lease.Grant(context.TODO(), 10); err != nil {
			goto RETRY
		}

		// 自动续租
		if keepAliveChan, err = register.lease.KeepAlive(context.TODO(), leaseGrantResp.ID); err != nil {
			goto RETRY
		}

		cancelCtx, cancelFunc = context.WithCancel(context.TODO())

		// 注册到etcd
		fmt.Println("register:", regKey, leaseGrantResp.ID)

		nodes = NodeServer{
			Status:  1,
			Pid:     os.Getpid(),
			Name:    name,
			Message: "success",
			IPAddr:  register.localIP,
		}
		if vl, err = json.Marshal(nodes); err != nil {
			fmt.Println(err)
			goto RETRY
		}
		if _, err = register.kv.Put(cancelCtx, regKey, string(vl), clientv3.WithLease(leaseGrantResp.ID)); err != nil {
			fmt.Println("err:", err.Error())
			goto RETRY
		}

		// 处理续租应答
		for {
			select {
			case keepAliveResp = <-keepAliveChan:
				if keepAliveResp == nil { // 续租失败
					goto RETRY
				}
			}
		}

	RETRY:
		time.Sleep(1 * time.Second)
		if cancelFunc != nil {
			fmt.Println("cancel func.")
			cancelFunc()
		}
	}
}

func InitRegister() (err error) {

	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		lease   clientv3.Lease
		localIp string
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // 集群地址
		DialTimeout: time.Duration(5 * time.Second),
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}

	// 本机IP
	if localIp, err = getLocalIP(); err != nil {
		return
	}

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	G_register = &Register{
		client:  client,
		kv:      kv,
		lease:   lease,
		localIP: localIp,
	}

	// 服务注册
	go G_register.keepOnline()

	return
}

func removeRegister() {
	// 服务名字
	name := "testmain"

	// 本机IP
	localIp, err := getLocalIP()
	if err != nil {
		fmt.Println("get ip:", err)
		return
	}
	// 注册路径
	regKey := WORKER_DIR + localIp + ":" + name
	// 删除KV   /*, clientv3.WithFromKey(), clientv3.WithLimit(2)*/
	fmt.Println("remove sign:", regKey)

	delResp, err := G_register.kv.Delete(context.TODO(), regKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 被删除之前的value是什么
	if len(delResp.PrevKvs) != 0 {
		for _, kvpair := range delResp.PrevKvs {
			fmt.Println("删除了:", string(kvpair.Key), string(kvpair.Value))
		}
	}
	log.Printf("delResp:%+v \n", delResp)
	log.Println("remove sign.", delResp.OpResponse())

	if getResp, err := G_register.kv.Get(context.TODO(), regKey, clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else { // 获取成功, 我们遍历所有的kvs
		fmt.Println("getResp", getResp.Kvs)
	}
}

package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func connectEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		// etcd服务端地址数组，可以配置一个或者多个
		Endpoints:   []string{"192.168.33.13:2379", "192.168.33.13:22379", "192.168.33.13:32379"},
		// 连接超时时间，5秒
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	go Watch(cli)
	Put(cli)
	Lease(cli)
	Get(cli)
	time.Sleep(5 * time.Second)
	Get(cli)
	select {

	}
	//Delete(cli)
}

func Get(cli *clientv3.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	// 读取key="/connect/url" 的值
	resp, err := cli.Get(ctx, "/connect/url")

	if err != nil {
		log.Fatal(err)
	}

	// 虽然这个例子我们只是查询一个Key的值，
	// 但是Get的查询结果可以表示多个Key的结果例如我们根据Key进行前缀匹配,Get函数可能会返回多个值。
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func Put(cli *clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	_, err := cli.Put(ctx, "/connect/url", "hello")
	if err != nil {
		log.Fatal(err)
	}

}

func Delete(cli *clientv3.Client) {
	// 获取上下文，设置请求超时时间为5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	// 删除key="/tizi365/url" 的值
	_, err := cli.Delete(ctx, "/connect/url")

	if err != nil {
		log.Fatal(err)
	}
}

func Watch(cli *clientv3.Client) {
	// 监控key=/tizi 的值
	rch := cli.Watch(context.Background(), "/connect/url")
	// 通过channel遍历key的值的变化
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func Lease(cli *clientv3.Client) {
	resp, err := cli.Grant(context.TODO(), 4)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, /nazha/ 这个key就会被移除
	_, err = cli.Put(context.TODO(), "/connect/url", "dsb", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka,ok := <-ch
		if !ok {
			log.Println("keepalive chan closed")
			break
		}
		fmt.Println("ttl:", ka.TTL)
	}
}

func main() {
	connectEtcd()
}

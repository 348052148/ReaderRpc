package test

import (
	"testing"
	"fmt"
	"context"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
	"go.etcd.io/etcd/clientv3"
)

func TestEtcdServiceRegister(t *testing.T)  {
	client,err:=etcdv3.NewClient(context.Background(),[]string{"127.0.0.1:2379"},etcdv3.ClientOptions{
		DialTimeout:3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client)
	client.Register(etcdv3.Service{
		Key:"ParserService",
		Value:"127.0.0.1:9001",
		TTL:etcdv3.NewTTLOption(time.Millisecond * 500, time.Second * 10),
	})
}

func TestEtcdServiceGet(t *testing.T)  {
	client,err:=etcdv3.NewClient(context.Background(),[]string{"127.0.0.1:2379"},etcdv3.ClientOptions{
		DialTimeout:3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	ens,err := client.GetEntries("ParserService")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ens)
}

func TestEtcdServiceWatch(t *testing.T)  {
	client,err:=etcdv3.NewClient(context.Background(),[]string{"127.0.0.1:2379"},etcdv3.ClientOptions{
		DialTimeout:3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	ch := make(chan struct{})
	client.WatchPrefix("ParserService", ch)
	for c := range ch {
		fmt.Printf("watch : %v", c)
	}
}

func TestClientGet(t *testing.T)  {
	client, err := clientv3.New(clientv3.Config{
		Context:context.Background(),
		DialTimeout:time.Second*3,
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println(client)
	res,_ := client.KV.Get(client.Ctx(), "foo", clientv3.WithPrefix())
	fmt.Println(res)
}

func TestClientSet(t *testing.T)  {
	client, err := clientv3.New(clientv3.Config{
		Context:context.Background(),
		DialTimeout:time.Second*3,
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	client.Put(client.Ctx(), "foo", "hi me")
}

func TestClientDel(t *testing.T)  {
	client, err := clientv3.New(clientv3.Config{
		Context:context.Background(),
		DialTimeout:time.Second*3,
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	client.Delete(client.Ctx(), "foo")
}

//监听
func TestWatch(t *testing.T)  {
	client, err := clientv3.New(clientv3.Config{
		Context:context.Background(),
		DialTimeout:time.Second*3,
		Endpoints:[]string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	ch :=client.Watch(client.Ctx(), "foo")

	for wresp := range ch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
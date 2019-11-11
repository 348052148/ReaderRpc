package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"ReadRpc/service"
	"ReadRpc/srv/protoc"
	"go.etcd.io/etcd/clientv3"
	"context"
	"time"
	"fmt"
	"flag"
	"ReadRpc/logs"
)

func main() {
	host := flag.String("host", "127.0.0.1", "服务主机")
	port := flag.String("port", "9001", "服务端口")
	etcdAddress := flag.String("etcd", "127.0.0.1:2379", "etcd地址")
	logPath := flag.String("logpath", ".", "日志目录")
	flag.Parse()
	fmt.Printf("Host:%s, Port: %s, Etcd: %s, logPath: %s \n", *host, *port, *etcdAddress, *logPath)
	//设置日志目录
	logs.SetLogPath(*logPath)
	logs.Init()
	//地址
	address := (*host) +":" + (*port)
	lis, err := net.Listen("tcp", address)  //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	rpcServer := grpc.NewServer() //创建gRPC服务
	//注册书籍服务
	srv.RegisterBookServiceServer(rpcServer, &service.BookService{})
	srv.RegisterParserServiceServer(rpcServer, &service.ParserService{})
	//注册反射服务
	reflection.Register(rpcServer)
	//服务注册
	RegisterService(address, *etcdAddress)
	// 将监听交给gRPC服务处理
	err = rpcServer.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
//注册服务
func RegisterService(address string, etcdAddress string)  {
	client,err :=clientv3.New(clientv3.Config{
		Context:context.Background(),
		Endpoints:[]string{etcdAddress},
		DialTimeout:time.Second * 5,
	})
	if err != nil {
		panic(err)
	}
	res, err := client.Put(client.Ctx(), "go-service/"+address, address)
	if err != nil {
		panic(err)
	}
	fmt.Println("register Success:",res)
}
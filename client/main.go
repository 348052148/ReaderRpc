package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"test/ReadRpc/book/protoc"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	bookClient := book.NewBookServiceClient(conn)

	// 模拟请求数据
	res, err := bookClient.SearchBookList(context.Background(),&book.SearchBookRequest{Keyword:"逆天"})
	//res, err := bookClient.GetBookChapterList(context.Background(), &book.ChapterRequest{Link: "http://www.quanshuwang.com/book/177/177605"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %v", res.Books)
}

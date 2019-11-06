package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"ReadRpc/srv/protoc"
	"fmt"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	//parserClient := srv.NewParserServiceClient(conn)
	// 模拟请求数据
	//link := "https://www.zaduw.com/1/1349/"
	//link := "https://www.zaduw.com/0/583/354985.html"
	//res, err := parserClient.ParserChapterContents(context.Background(),&srv.ChapterContentRequest{Link:link,Source:"zadu"})
	//res, err := bookClient.GetBookChapterList(context.Background(), &book.ChapterRequest{Link: "http://www.quanshuwang.com/book/177/177605"})
	bookClient := srv.NewBookServiceClient(conn)
	res,_:= bookClient.GetBookSourceChapterInfo(context.Background(), &srv.SourceChapterRequest{
		ChapterSource: []*srv.SourceChapterRequest_ChapterSource{{ChapterLink: "http://www.xbiquge.la/32/32949/", Source: "xbiquge"},
			{ChapterLink: "http://www.quanshuwang.com/book/171/171131", Source: "quanwen"}},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(res.ChapterInfo)
}

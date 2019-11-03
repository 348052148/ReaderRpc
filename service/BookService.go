package service

import (
	"context"
	"ReadRpc/parser"
	"fmt"
	url2 "net/url"
	"ReadRpc/srv/protoc"
)

type BookService struct {
}

func (bookService *BookService) GetBookInfo(cxt context.Context, req *srv.BookRequest) (*srv.BookResponse, error) {

	return nil, nil
}

func (bookService *BookService) SearchBookList(cxt context.Context, req *srv.SearchBookRequest) (*srv.SearchBookResponse, error) {
	//http://www.quanshuwang.com/modules/article/search.php?searchkey=%CF%C9%C4%E6&searchtype=articlename&searchbuttom.x=40&searchbuttom.y=13
	quanWenParser := &parser.QuanwenParser{}
	url := "http://www.quanshuwang.com/modules/article/search.php?"
	v := url2.Values{}
	v.Add("searchkey", req.Keyword)
	v.Add("searchtype", "articlename");
	v.Add("searchbuttom.x", "40");
	v.Add("searchbuttom.y", "13");
	url += v.Encode()
	fmt.Println(url)
	books, _ := quanWenParser.ParserSearchBooks(url)
	var bookList []*srv.SearchBookResponse_Book;
	for _, bookInfo := range books {
		bookList = append(bookList, &srv.SearchBookResponse_Book{
			Title:  bookInfo.Title,
			Author: bookInfo.Author,
			Cover:  bookInfo.Cover,
			Link:   bookInfo.ChapterLink,
			Detail: bookInfo.Detail,
		})
	}

	return &srv.SearchBookResponse{Books: bookList}, nil
}

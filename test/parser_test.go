package test

import (
	"testing"
	"ReadRpc/parser"
	"fmt"
	url2 "net/url"
)

func TestParser(t *testing.T)  {
	quanWenParser := &parser.QuanwenParser{}
	url := "http://www.quanshuwang.com/modules/article/search.php?"
	v := url2.Values{}
	v.Add("searchkey", "逆天")
	v.Add("searchtype", "articlename");
	v.Add("searchbuttom.x", "40");
	v.Add("searchbuttom.y", "13");
	url += v.Encode()
	fmt.Println(url2.ParseQuery(url))
	books,_:=quanWenParser.ParserSearchBooks(url)//"http://www.quanshuwang.com/modules/article/search.php?searchkey=%CF%C9%C4%E6&searchtype=articlename&searchbuttom.x=40&searchbuttom.y=13")
		fmt.Println(books)
}

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

//测试解析分类书籍
func TestZaduParser(t *testing.T)  {
	zaduParser := &parser.ZaduParser{}
	classifys, _ := zaduParser.ParserClassflysBooks("https://www.zaduw.com/sort1/1/")

	fmt.Println(classifys)
}

//
func TestZaduParserBookInfo(t *testing.T)  {
	zaduParser := &parser.ZaduParser{}
	bookinfo,_ := zaduParser.ParserBookInfo("https://www.zaduw.com/3/3965/", 1)

	//http://www.quanshuwang.com/book_123637.html
	quanWenParser := &parser.QuanwenParser{}
	bookinfo1,_ := quanWenParser.ParserBookInfo("http://www.quanshuwang.com/book_123637.html", 1);

	fmt.Println(bookinfo)

	fmt.Println(bookinfo1)
}

//测试解析章节
func TestParserChapters(t *testing.T)  {
	zaduParser := &parser.ZaduParser{}
	chapters,_ := zaduParser.ParserChapters("https://www.zaduw.com/3/3965/", "")

	fmt.Println(chapters)
}
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

func TestParserChapterContents(t *testing.T)  {
	zaduParser := &parser.ZaduParser{}
	c,_:=zaduParser.ParserChapterContents("https://www.zaduw.com/0/583/354985.html")
	fmt.Println(c)
}

func TestParserQuanwenContents(t *testing.T)  {
	quanWenParser := &parser.QuanwenParser{}
	//c,_:=quanWenParser.TransformQuanWenMobileUrl("http://www.quanshuwang.com/book/20/20425/8568168.html")
	c,_:=quanWenParser.ParserChapterContents("http://www.quanshuwang.com/book/20/20425/8568168.html")
	fmt.Println(c)
}

func TestXBiqugeParser(t *testing.T)  {
	xbiqugeParser := parser.NewXbiqugeParser();
	classifys,_ := xbiqugeParser.ParserClassflysBooks("http://www.xbiquge.la/fenlei/1_1.html")
	fmt.Println(classifys)
}

func TestXBiqugeParserBook(t *testing.T)  {
	//http://www.xbiquge.la/32/32391/
	xbiqugeParser := parser.NewXbiqugeParser();
	bookinfo,_ := xbiqugeParser.ParserBookInfo("http://www.xbiquge.la/32/32391/", 1)
	fmt.Println(bookinfo)
}

func TestXBiqugeParserChapters(t *testing.T)  {
	xbiqugeParser := parser.NewXbiqugeParser();
	chapters,_ := xbiqugeParser.ParserChapters("http://www.xbiquge.la/32/32391/", "")
	fmt.Println(chapters)
}

func TestXBiqugeParserChapterContents(t *testing.T)  {
	xbiqugeParser := parser.NewXbiqugeParser();
	contents,_ := xbiqugeParser.ParserChapterContents("http://www.xbiquge.la/32/32391/15246708.html")
	fmt.Println(contents)
}
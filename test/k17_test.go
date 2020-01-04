package test

import (
	"testing"
	parser2 "ReadRpc/parser"
)

func TestK17ParserClassify(t *testing.T)  {
	parser := parser2.NewK17Parser()
	classifys,_ := parser.ParserClassflysBooks("https://www.17k.com/all/book/2_21_0_0_0_0_1_0_1.html")
	t.Logf("%v", classifys)
}

func TestK17ParserBookInfo(t *testing.T)  {
	//
	parser := parser2.NewK17Parser()
	book,_ := parser.ParserBookInfo("https://www.17k.com/book/3051176.html", 0)
	t.Logf("%v", book)
}

func TestK17ParserBookChapters(t *testing.T)  {
	parser := parser2.NewK17Parser()
	chapters,_ :=parser.ParserChapters("https://www.17k.com/list/2681541.html", "1234")
	t.Logf("%v", chapters)
}

func TestK17ParserChapterContents(t *testing.T)  {
	parser := parser2.NewK17Parser()
	contents,_ := parser.ParserChapterContents("https://www.17k.com/chapter/3051176/38819246.html")
	t.Logf("%v", contents)
}

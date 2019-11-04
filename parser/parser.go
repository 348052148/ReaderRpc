package parser

import "ReadRpc/entitys"

//接口
type Parser interface {
	ParserClassflysBooks(url string) (entitys.Classfly,error)

	ParserBookInfo(url string, classifyId int) (entitys.BookInfo, error)

	ParserChapters(url string, bookId string) ([]entitys.Chapter, error)
}

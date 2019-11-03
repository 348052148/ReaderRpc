package service

import (
	"ReadRpc/srv/protoc"
	"context"
	"fmt"
	"ReadRpc/parser"
)

type ParserService struct {
}

// 获取书籍章节列表
func (parserService *ParserService) ParserChapters(cxt context.Context, req *srv.ChapterRequest) (*srv.ChapterResponse, error) {
	quanWenParser := &parser.QuanwenParser{}
	fmt.Println(req.Link)
	chapters, _ := quanWenParser.ParserChapters(req.Link, "1")
	var chapterList []*srv.ChapterResponse_Chapter
	for _, chapter := range chapters {
		chapterList = append(chapterList, &srv.ChapterResponse_Chapter{
			Title: chapter.Title,
			Index: int32(chapter.Index),
			ContentsLink:  chapter.ContentLink,
		})
	}
	return &srv.ChapterResponse{Chapters: chapterList}, nil
}

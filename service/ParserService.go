package service

import (
	"ReadRpc/srv/protoc"
	"context"
	"ReadRpc/parser"
	"ReadRpc/logs"
)

type ParserService struct {
}

// 获取书籍章节列表
func (parserService *ParserService) ParserChapters(cxt context.Context, req *srv.ChapterRequest) (*srv.ChapterResponse, error) {
	logs.Info("Chapter", []interface{}{req.Link, req.Source})
	parserEngin := parserService.BuilderParser(req.Source)
	chapters, err := parserEngin.ParserChapters(req.Link, "1")
	if err != nil {
		logs.Error("Chapter", err)
		return &srv.ChapterResponse{}, nil
	}
	var chapterList []*srv.ChapterResponse_Chapter
	for _, chapter := range chapters {
		chapterList = append(chapterList, &srv.ChapterResponse_Chapter{
			Title: chapter.Title,
			Index: int32(chapter.Index),
			ContentsLink:  chapter.ContentLink,
			Source:req.Source,
		})
	}
	return &srv.ChapterResponse{Chapters: chapterList}, nil
}

func (parserService *ParserService) ParserChapterContents(cxt context.Context, req *srv.ChapterContentRequest)(*srv.ChapterContentResponse, error)  {
	logs.Info("Contents", []interface{}{req.Link, req.Source})
	parserEngin := parserService.BuilderParser(req.Source)
	contents, _ := parserEngin.ParserChapterContents(req.Link)

	return &srv.ChapterContentResponse{Contents:contents}, nil
}

func (parserService *ParserService)BuilderParser(flag string) parser.Parser  {
	var parserEngin parser.Parser
	if flag == "quanwen" {
		parserEngin = parser.NewQuanwenParser()
	}else if flag == "zadu" {
		parserEngin = parser.NewZaduParser()
	}else if flag == "xbiquge" {
		parserEngin = parser.NewXbiqugeParser()
	}else if flag == "17k" {
		parserEngin = parser.NewK17Parser()
	}
	return parserEngin
}

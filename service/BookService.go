package service

import (
	"context"
	"ReadRpc/parser"
	"fmt"
	url2 "net/url"
	"ReadRpc/srv/protoc"
	"sync"
	"ReadRpc/logs"
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

//获取源章节信息
func (bookService *BookService) GetBookSourceChapterInfo(ctx context.Context, req *srv.SourceChapterRequest) (*srv.SourceChapterResponse, error) {
	logs.Info("GetBookSourceChapter :", req.ChapterSource)
	var chapterInfos []*srv.SourceChapterResponse_ChapterInfo;
	wg := &sync.WaitGroup{}
	for _, chapterSource := range req.ChapterSource {
		wg.Add(1)
		go func(source string, chapterLink string) {
			defer wg.Done()
			parserEngine := bookService.BuilderParser(source)
			chapters, err := parserEngine.ParserChapters(chapterLink, "1")
			if err == nil {
				chapterCount := len(chapters)
				chapterInfos = append(chapterInfos, &srv.SourceChapterResponse_ChapterInfo{
					ChapterLink:  chapterLink,
					ChapterCount: int32(chapterCount),
					Source:       source,
				})
			}
		}(chapterSource.Source, chapterSource.ChapterLink)
	}
	wg.Wait()
	fmt.Println(chapterInfos)
	return &srv.SourceChapterResponse{ChapterInfo: chapterInfos}, nil
}

/**
var chapterInfos []*srv.SourceChapterResponse_ChapterInfo;
	chapterInfoChan := make(chan *srv.SourceChapterResponse_ChapterInfo)
	fmt.Println("Last-Source")
	wg := &sync.WaitGroup{}
	for _, chapterSource := range req.ChapterSource {
		wg.Add(1)
		go func(source string, chapterLink string) {
			defer wg.Done()
			parserEngine := bookService.BuilderParser(source)
			chapters, _ := parserEngine.ParserChapters(chapterLink, "1")
			chapterCount := len(chapters)
			chapterInfoChan <- &srv.SourceChapterResponse_ChapterInfo{
				ChapterLink:  chapterLink,
				ChapterCount: int32(chapterCount),
				Source:       source,
			}
		}(chapterSource.Source, chapterSource.ChapterLink)
	}
	go func() {
		for chapterInfo := range chapterInfoChan {
			fmt.Println(chapterInfo)
			chapterInfos = append(chapterInfos, chapterInfo)
		}
	}()
	wg.Wait()
	fmt.Println(chapterInfos)
	return &srv.SourceChapterResponse{ChapterInfo: chapterInfos}, nil
 */

func (bookService *BookService) BuilderParser(flag string) parser.Parser {
	var parserEngin parser.Parser
	if flag == "quanwen" {
		parserEngin = parser.NewQuanwenParser()
	} else if flag == "zadu" {
		parserEngin = parser.NewZaduParser()
	} else if flag == "xbiquge" {
		parserEngin = parser.NewXbiqugeParser()
	}
	return parserEngin
}

package parser

import (
	"github.com/PuerkitoBio/goquery"
	"ReadRpc/entitys"
	"io"
	"ReadRpc/msg"
	"net/http"
	"time"
	"fmt"
	"strings"
)

type XbiqugeParser struct {
	linkSet *msg.LinkSet
}

func NewXbiqugeParser() *XbiqugeParser {
	return &XbiqugeParser{}
}

func (parser *XbiqugeParser) SetLinkSet(linkSet *msg.LinkSet) {
	parser.linkSet = linkSet
}

func (parser *XbiqugeParser) Request(url string) (io.ReadCloser, error) {
	client := http.Client{Timeout: time.Second * 30}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, err
	}
	return res.Body, nil
}

func (parser *XbiqugeParser) ParserClassflysBooks(url string) (entitys.Classfly, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("ClassFly TIME OUT" + url)
		return entitys.Classfly{}, reqErr
	}
	//defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("ClassFly BAN OUT" + url)
		return entitys.Classfly{}, err
	}
	var bookDetailList []entitys.BookDetail
	fmt.Println(url)
	// Find the review items
	doc.Find("#newscontent .l ul li").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i)
		// For each item found, get the band and title
		title := s.Find(".s2>a").Text()
		href, _ := s.Find(".s2>a").Attr("href")
		cover, _ := s.Find(".s2>a").Attr("href")
		//fmt.Printf("书名： %s , 封面: %s , 连接 %s\n",  title, cover, href)
		bookDetailList = append(bookDetailList, entitys.BookDetail{
			Title: title,
			Link:  href,
			Cover: cover,
		})
	})
	return entitys.Classfly{
		Title: "",
		Books: bookDetailList,
		Cover: "",
	}, nil;
}

func (parser *XbiqugeParser) ParserBookInfo(url string, classifyId int) (entitys.BookInfo, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("BOOK TIME OUT" + url)
		return entitys.BookInfo{}, reqErr
	}
	//defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("BOOK BAN TIME OUT" + url)
		return entitys.BookInfo{}, err
	}
	title := doc.Find("#info>h1").Eq(0).Text()
	author :=  strings.TrimSpace(strings.Split(doc.Find("#info>p").Eq(0).Text(), "：")[1])
	detail := doc.Find("#intro>p").Eq(1).Text()
	//author := doc.Find("#info>p").Eq(0).Text()
	cover, _ := doc.Find("#fmimg>img").Eq(0).Attr("src")
	href := url
	status := "连载"
	//
	if title == "" || href == "" {
		fmt.Println("BOOK Parser Faill : " + url)
	}
	return entitys.BookInfo{
		BookId:      entitys.Md5(title + author),
		Title:       title,
		Author:      author,
		Detail:      detail,
		Cover:       cover,
		Status:      status,
		ChapterLink: href,
		Classify_id: classifyId,
	}, nil
}

func (parser *XbiqugeParser) ParserChapters(url string, bookId string) ([]entitys.Chapter, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return []entitys.Chapter{}, reqErr
	}
	//defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return []entitys.Chapter{}, err
	}
	var ChapterList []entitys.Chapter
	doc.Find("#list dl dd>a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		//fmt.Printf("title %s link : %s \n",s.Text(), link)
		ChapterList = append(ChapterList, entitys.Chapter{
			BookId:      bookId,
			Title:       s.Text(),
			Index:       i,
			ContentLink: "http://www.xbiquge.la" + link,
		})
	})
	return ChapterList, nil
}

func  (parser *XbiqugeParser) ParserChapterContents(url string) (string, error)  {
	//yd_text2
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return "", reqErr
	}
	//defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return "", err
	}
	contents := doc.Find("#content").Text()
	return contents, nil
}

package parser

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"github.com/PuerkitoBio/goquery"
	"ReadRpc/entitys"
	"io"
	"ReadRpc/msg"
	"net/http"
	"time"
	"fmt"
	"strings"
)

type ZaduParser struct {
	linkSet *msg.LinkSet
}

func NewZaduParser() *ZaduParser {
	return &ZaduParser{}
}

func (parser *ZaduParser) SetLinkSet(linkSet *msg.LinkSet) {
	parser.linkSet = linkSet
}

func (parser *ZaduParser) Request(url string) (io.ReadCloser, error) {
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


func (parser *ZaduParser) ParserClassflysBooks(url string) (entitys.Classfly, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("ClassFly TIME OUT" + url)
		return entitys.Classfly{}, reqErr
	}
	fmt.Println(url)
	//defer body.Close()
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("ClassFly BAN OUT" + url)
		return entitys.Classfly{}, err
	}
	var bookDetailList []entitys.BookDetail
	// Find the review items
	doc.Find(".booklist ul li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title, _ := s.Find(".sm>a").Eq(0).Attr("title")
		href, _ := s.Find(".sm>a").Attr("href")
		cover, _ := s.Find(".sm>a").Attr("href")
		if title == "" || href == "" {
			return
		}
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

func (parser *ZaduParser) ParserBookInfo(url string, classifyId int) (entitys.BookInfo, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("BOOK TIME OUT" + url)
		return entitys.BookInfo{}, reqErr
	}
	//defer body.Close()
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("BOOK BAN TIME OUT" + url)
		return entitys.BookInfo{}, err
	}
	title := doc.Find(".jieshao .rt>h1").Eq(0).Text()
	detail := doc.Find(".jieshao .intro").Eq(0).Text()
	cover, _ := doc.Find(".jieshao .lf>img").Eq(0).Attr("src")
	href := url
	var author, status string
	doc.Find(".jieshao .msg em").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			author = strings.TrimSpace(strings.Split(s.Text(), "：")[1])
		}
		if i == 1 {
			status = strings.TrimSpace(strings.Split(s.Text(), "：")[1])
		}
	})
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

func (parser *ZaduParser) ParserChapters(url string, bookId string) ([]entitys.Chapter, error) {
	body, reqErr := parser.Request(url)
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return []entitys.Chapter{}, reqErr
	}
	defer body.Close()
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return []entitys.Chapter{}, err
	}
	ChapterList := make([]entitys.Chapter, 0)
	doc.Find(".mulu ul li>a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		//fmt.Printf("title %s link : %s \n",s.Text(), link)
		index := i - 9
		if index >= 0 {
			ChapterList = append(ChapterList, entitys.Chapter{
				BookId:      bookId,
				Title:       s.Text(),
				Index:       index,
				ContentLink: url + link,
			})
		}
	})
	return ChapterList, nil
}

func  (parser *ZaduParser) ParserChapterContents(url string) (string, error)  {
	//yd_text2
	body, reqErr := parser.Request(url)
	defer body.Close()
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return "", reqErr
	}
	//defer body.Close()
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return "", err
	}
	contents := doc.Find(".novel .yd_text2").Text()
	return contents, nil
}

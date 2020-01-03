package parser

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"time"
	//"test/gopanc/engine"
	"ReadRpc/msg"
	"fmt"
	"ReadRpc/entitys"
	"strings"
	urls "net/url"
)

type QuanwenParser struct {
	linkSet *msg.LinkSet
}

func NewQuanwenParser() *QuanwenParser {
	return &QuanwenParser{}
}

func (parser *QuanwenParser) SetLinkSet(linkSet *msg.LinkSet) {
	parser.linkSet = linkSet
}

func (parser *QuanwenParser) Request(url string) (io.ReadCloser, error) {
	//transport := &http.Transport{
	//	Proxy: func(request *http.Request) (*urls.URL, error) {
	//		return urls.Parse("http://58.218.214.165:3487")
	//	}}
	client := http.Client{Timeout: time.Second * 10}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Add("Referer", url)

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, err
	}
	return res.Body, nil
}

func (parser *QuanwenParser) ParserClassflysBooks(url string) (entitys.Classfly, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("ClassFly TIME OUT" + url)
		return entitys.Classfly{}, reqErr
	}
	//defer body.Close()
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("ClassFly BAN OUT" + url)
		return entitys.Classfly{}, err
	}
	var bookDetailList []entitys.BookDetail
	// Find the review items
	doc.Find(".seeWell li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("span>a").Eq(0).Text()
		href, _ := s.Find(".mr10").Attr("href")
		cover, _ := s.Find("a>img").Attr("src")
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

func (parser *QuanwenParser) ParserBookInfo(url string, classifyId int) (entitys.BookInfo, error) {
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
	title := doc.Find(".b-info h1").Eq(0).Text()
	detail := doc.Find(".b-info .infoDetail #waa").Eq(0).Text()
	cover, _ := doc.Find(".detail .mr11>img").Eq(0).Attr("src")
	href, _ := doc.Find(".detail .mr11").Eq(0).Attr("href")
	var author, status string
	doc.Find(".author .bookDetail dl").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			status = strings.TrimSpace(s.Find("dd").Text())
		}
		if i == 1 {
			author = strings.TrimSpace(s.Find("dd").Text())
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

func (parser *QuanwenParser) ParserChapters(url string, bookId string) ([]entitys.Chapter, error) {
	body, reqErr := parser.Request(url)
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return []entitys.Chapter{}, reqErr
	}
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return []entitys.Chapter{}, err
	}
	defer body.Close();
	ChapterList := make([]entitys.Chapter, 0)
	doc.Find(".chapterSo .chapterNum ul li>a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		//fmt.Printf("title %s link : %s \n",s.Text(), link)
		ChapterList = append(ChapterList, entitys.Chapter{
			BookId:      bookId,
			Title:       s.Text(),
			Index:       i,
			ContentLink: link,
		})
	})
	return ChapterList, nil
}

func (parser *QuanwenParser) ParserSearchBooks(url string) ([]entitys.BookInfo, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return []entitys.BookInfo{}, reqErr
	}
	//defer body.Close()
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return []entitys.BookInfo{}, err
	}

	var BookList []entitys.BookInfo
	doc.Find(".seeWell li").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
		cover, _ := s.Find("a>img").Eq(0).Attr("src")
		title := s.Find("span>a").Eq(0).Text()
		author := s.Find("span>a").Eq(1).Text()
		link, _ := s.Find("span>a").Eq(2).Attr("href")
		detail := s.Find("span>em").Eq(0).Text()
		//fmt.Printf("title %s link : %s %s %s \n",title, author, cover, detail)
		BookList = append(BookList, entitys.BookInfo{
			BookId:      "",
			Title:       title,
			Author:      author,
			Cover:       cover,
			ChapterLink: link,
			Detail:      detail,
		})
	})
	return BookList, nil
}

func (parser *QuanwenParser) ParserChapterContents(url string) (string, error) {
	body, reqErr := parser.Request(url)
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return "", reqErr
	}
	bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(bytes)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return "", err
	}
	defer body.Close()
	contents := doc.Find(".bookInfo .mainContenr").Text()
	return contents, nil
}

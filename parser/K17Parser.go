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

type K17Parser struct {
	linkSet *msg.LinkSet
}

func NewK17Parser() *K17Parser {
	return &K17Parser{}
}

func (parser *K17Parser) SetLinkSet(linkSet *msg.LinkSet) {
	parser.linkSet = linkSet
}

func (parser *K17Parser) Request(url string) (io.ReadCloser, error) {
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

func (parser *K17Parser) ParserClassflysBooks(url string) (entitys.Classfly, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("ClassFly TIME OUT" + url)
		return entitys.Classfly{}, reqErr
	}
	fmt.Println(url)
	//defer body.Close()
	//bytes := transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("ClassFly BAN OUT" + url)
		return entitys.Classfly{}, err
	}
	var bookDetailList []entitys.BookDetail
	// Find the review items
	doc.Find(".search-list .alltable table tbody tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find(".td3 a").Eq(0).Text();
		href, _ := s.Find(".td3 a").Attr("href")
		cover := ""
		if title == "" || href == "" {
			return
		}
		if title == "xxxx" {
			return
		}
		//fmt.Printf("书名： %s , 封面: %s , 连接 %s\n",  title, cover, href)
		bookDetailList = append(bookDetailList, entitys.BookDetail{
			Title: title,
			Link:  "https:" + href,
			Cover: cover,
		})
	})
	return entitys.Classfly{
		Title: "",
		Books: bookDetailList,
		Cover: "",
	}, nil;
}

func (parser *K17Parser) ParserBookInfo(url string, classifyId int) (entitys.BookInfo, error) {
	body, reqErr := parser.Request(url)
	//defer body.Close()
	if reqErr != nil {
		fmt.Println("BOOK TIME OUT" + url)
		return entitys.BookInfo{}, reqErr
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("BOOK BAN TIME OUT" + url)
		return entitys.BookInfo{}, err
	}
	defer body.Close()
	title := strings.TrimSpace(doc.Find(".BookInfo .Info>h1>a").Eq(0).Text())
	detail := doc.Find(".BookInfo .Info .Tab .intro").Eq(0).Text()
	cover, _ := doc.Find(".BookInfo .Props .cover .book").Eq(0).Attr("src")
	href, _ := doc.Find(".BookInfo .Props .cover>a").Attr("href")
	author := strings.TrimSpace(doc.Find(".AuthorInfo .author .name").Text())
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
		Cover:       "https:" + cover,
		Status:      status,
		ChapterLink: "https://www.17k.com" + href,
		Classify_id: classifyId,
	}, nil
}

func (parser *K17Parser) ParserChapters(url string, bookId string) ([]entitys.Chapter, error) {
	body, reqErr := parser.Request(url)
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return []entitys.Chapter{}, reqErr
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return []entitys.Chapter{}, err
	}
	defer body.Close()
	ChapterList := make([]entitys.Chapter, 0)
	doc.Find(".List .Volume>dd>a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		//fmt.Printf("title %s link : %s \n",s.Text(), link)
		index := i - 9
		if index >= 0 {
			ChapterList = append(ChapterList, entitys.Chapter{
				BookId:      bookId,
				Title:       strings.TrimSpace(s.Text()),
				Index:       index,
				ContentLink: "https://www.17k.com" + link,
			})
		}
	})
	return ChapterList, nil
}

func (parser *K17Parser) ParserChapterContents(url string) (string, error) {
	//yd_text2
	body, reqErr := parser.Request(url)
	if reqErr != nil {
		fmt.Println("Chapter TIME OUT" + url)
		return "", reqErr
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println("Chapter BAN TIME OUT" + url)
		return "", err
	}
	defer body.Close()
	contents := doc.Find(".readArea .content .p").Text()
	return contents, nil
}

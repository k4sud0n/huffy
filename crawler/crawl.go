package crawler

import (
	"database/sql"
	"github.com/gocolly/colly/v2"
	"huffy/database"
	"strings"
	"sync"
	"time"
)

type SiteCralwer struct {
	URL      string
	Selector string
	Handler  func(e *colly.HTMLElement)
}

func CrawlData(db *sql.DB) error {
	weekAgo := time.Now().AddDate(0, 0, -7)

	sites := []SiteCralwer{
		// 한국외대 공지
		{
			URL:      "https://www.hufs.ac.kr/hufs/11281/subview.do",
			Selector: "tr",
			Handler: func(e *colly.HTMLElement) {
				articleDateStr := strings.TrimSpace(e.ChildText("td:nth-child(4)"))
				articleDate, _ := time.Parse("2006.01.02", articleDateStr)
				if articleDate.After(weekAgo) {
					articleTitle := strings.TrimSpace(e.ChildText("td:nth-child(2) strong"))
					articleLink := "https://hufs.ac.kr/" + e.ChildAttr("td:nth-child(2) a", "href")
					_ = database.SaveData(db, articleTitle, articleLink, articleDate)
				}
			},
		},
		// 한국외대 학사
		{
			URL:      "https://www.hufs.ac.kr/hufs/11282/subview.do",
			Selector: "tr",
			Handler: func(e *colly.HTMLElement) {
				articleDateStr := strings.TrimSpace(e.ChildText("td:nth-child(4)"))
				articleDate, _ := time.Parse("2006.01.02", articleDateStr)
				if articleDate.After(weekAgo) {
					articleTitle := strings.TrimSpace(e.ChildText("td:nth-child(2) strong"))
					articleLink := "https://hufs.ac.kr/" + e.ChildAttr("td:nth-child(2) a", "href")
					_ = database.SaveData(db, articleTitle, articleLink, articleDate)
				}
			},
		},
		// 한국외대 장학
		{
			URL:      "https://www.hufs.ac.kr/hufs/11283/subview.do",
			Selector: "tr",
			Handler: func(e *colly.HTMLElement) {
				articleDateStr := strings.TrimSpace(e.ChildText("td:nth-child(4)"))
				articleDate, _ := time.Parse("2006.01.02", articleDateStr)
				if articleDate.After(weekAgo) {
					articleTitle := strings.TrimSpace(e.ChildText("td:nth-child(2) strong"))
					articleLink := "https://hufs.ac.kr/" + e.ChildAttr("td:nth-child(2) a", "href")
					_ = database.SaveData(db, articleTitle, articleLink, articleDate)
				}
			},
		},
		// 한국외대 채용
		{
			URL:      "https://www.hufs.ac.kr/hufs/11284/subview.do",
			Selector: "tr",
			Handler: func(e *colly.HTMLElement) {
				articleDateStr := strings.TrimSpace(e.ChildText("td:nth-child(4)"))
				articleDate, _ := time.Parse("2006.01.02", articleDateStr)
				if articleDate.After(weekAgo) {
					articleTitle := strings.TrimSpace(e.ChildText("td:nth-child(2) strong"))
					articleLink := "https://hufs.ac.kr/" + e.ChildAttr("td:nth-child(2) a", "href")
					_ = database.SaveData(db, articleTitle, articleLink, articleDate)
				}
			},
		},
		// 한국외대 AI교육원
		{
			URL:      "https://builder.hufs.ac.kr/user/indexSub.action?codyMenuSeq=129898191&siteId=soft&page=1",
			Selector: "tr",
			Handler: func(e *colly.HTMLElement) {
				articleDateStr := strings.TrimSpace(e.ChildText("td:nth-child(2)"))
				articleDate, _ := time.Parse("2006-01-02", articleDateStr)
				if articleDate.After(weekAgo) {
					articleTitle := strings.TrimSpace(e.ChildText("td:nth-child(1) a"))
					articleLink := "http://builder.hufs.ac.kr/user/" + e.ChildAttr("td:nth-child(1) a", "href")
					_ = database.SaveData(db, articleTitle, articleLink, articleDate)
				}
			},
		},
		// 한국외대 컴퓨터공학부 공지시항
		{
			URL:      "https://computer.hufs.ac.kr/computer/10058/subview.do",
			Selector: "tr",
			Handler: func(e *colly.HTMLElement) {
				articleDateStr := strings.TrimSpace(e.ChildText("td:nth-child(4)"))
				articleDate, _ := time.Parse("2006.01.02", articleDateStr)
				if articleDate.After(weekAgo) {
					articleTitle := strings.TrimSpace(e.ChildText("td:nth-child(2) strong"))
					articleLink := "https://hufs.ac.kr/" + e.ChildAttr("td:nth-child(2) a", "href")
					_ = database.SaveData(db, articleTitle, articleLink, articleDate)
				}
			},
		},
	}

	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)

		go func(site SiteCralwer) {
			defer wg.Done()

			siteCollector := colly.NewCollector()
			siteCollector.OnHTML(site.Selector, site.Handler)
			siteCollector.Visit(site.URL)
		}(site)
	}

	wg.Wait()
	return nil
}

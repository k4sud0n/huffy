package crawler

import (
	"database/sql"
	"fmt"
	"github.com/gocolly/colly/v2"
	"huffy/database"
	"strings"
	"sync"
	"time"
)

func GetMenu(db *sql.DB) error {
	currentTime := time.Now()

	sites := []SiteCrawler{
		// 후생관 학생식당
		{
			URL:      fmt.Sprintf("https://wis.hufs.ac.kr/jsp/HUFS/cafeteria/viewWeek.jsp?startDt=%s&endDt=%s&caf_id=h203", currentTime.Format("20060102"), currentTime.Format("20060102")),
			Selector: "table",
			Handler: func(e *colly.HTMLElement) {
				var content string

				e.ForEach("tr[height='35']", func(rowIndex int, row *colly.HTMLElement) {
					text := strings.TrimSpace(row.Text)

					switch rowIndex {
					case 1:
						content += text + "\n\n==========\n\n"
					case 2:
						content += text + "\n\n==========\n\n"
					case 3:
						content += text
					}
				})

				if strings.TrimSpace(content) == "" {
					fmt.Println("Skipping empty menu entry")
					return
				}

				database.SaveMenu(db, currentTime.Format("2006/01/02"), "husaeng_student", content)
			},
		},
		// 후생관 교직원식당
		{
			URL:      fmt.Sprintf("https://wis.hufs.ac.kr/jsp/HUFS/cafeteria/viewWeek.jsp?startDt=%s&endDt=%s&caf_id=h202", currentTime.Format("20060102"), currentTime.Format("20060102")),
			Selector: "table",
			Handler: func(e *colly.HTMLElement) {
				var content string

				e.ForEach("tr[height='35']", func(rowIndex int, row *colly.HTMLElement) {
					text := strings.TrimSpace(row.Text)

					switch rowIndex {
					case 1:
						content += text
					}
				})

				if strings.TrimSpace(content) == "" {
					fmt.Println("Skipping empty menu entry")
					return
				}

				database.SaveMenu(db, currentTime.Format("2006/01/02"), "husaeng_professor", content)
			},
		},
	}

	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)

		go func(site SiteCrawler) {
			defer wg.Done()

			siteCollector := colly.NewCollector()
			siteCollector.OnHTML(site.Selector, site.Handler)
			siteCollector.Visit(site.URL)
		}(site)
	}

	wg.Wait()
	return nil
}

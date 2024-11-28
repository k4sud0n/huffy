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
	sites := []SiteCrawler{
		// 후생관 학생식당
		{
			URL:      "https://wis.hufs.ac.kr/jsp/HUFS/cafeteria/viewWeek.jsp?startDt=20241125&endDt=20241201&caf_id=h203",
			Selector: "table",
			Handler: func(e *colly.HTMLElement) {
				currentTime := time.Now()
				weekday := int(currentTime.Weekday()) + 1

				e.ForEach("tr[height='35']", func(i int, row *colly.HTMLElement) {
					row.ForEach(fmt.Sprintf("td:nth-child(%d)", weekday), func(j int, cell *colly.HTMLElement) {
						var menuDate string
						var menu string
						text := strings.TrimSpace(cell.Text)

						if i == 0 {
							menuDate = text
						} else {
							menu += text
						}

						database.SaveMenu(db, menuDate, menu)
					})
				})
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

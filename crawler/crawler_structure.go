package crawler

import "github.com/gocolly/colly/v2"

type SiteCrawler struct {
	URL      string
	Selector string
	Handler  func(e *colly.HTMLElement)
}

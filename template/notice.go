package template

import (
	"github.com/gofiber/fiber/v2"
	"huffy/database"
)

type Template struct {
	Outputs []Output `json:"outputs"`
}

type Output struct {
	ListCard ListCard `json:"listCard"`
}

type ListCard struct {
	Header  Header   `json:"header"`
	Items   []Item   `json:"items"`
	Buttons []Button `json:"buttons"`
}

type Header struct {
	Title string `json:"title"`
}

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Link        Link   `json:"link"`
}

type Link struct {
	Web string `json:"web"`
}

type Button struct {
	Action     string `json:"action"`
	Label      string `json:"label"`
	WebLinkUrl string `json:"webLinkUrl"`
}

func CreateItems(notices []database.Notice) []Item {
	items := make([]Item, 0, len(notices))
	for _, notice := range notices {
		items = append(items, Item{
			Title:       notice.Title,
			Description: notice.Date,
			ImageUrl:    "https://t1.kakaocdn.net/openbuilder/sample/img_002.jpg",
			Link: Link{
				Web: notice.Link,
			},
		})
	}
	return items
}

func CreateResponse(item []Item) fiber.Map {
	response := fiber.Map{
		"version": "2.0",
		"template": Template{
			Outputs: []Output{
				{
					ListCard: ListCard{
						Header: Header{
							Title: "공지사항",
						},
						Items: item,
						Buttons: []Button{
							{
								Action:     "webLink",
								Label:      "더 확인하기",
								WebLinkUrl: "https://hufs.ac.kr",
							},
						},
					},
				},
			},
		},
	}
	return response
}

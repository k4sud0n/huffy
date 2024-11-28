package template

import (
	"github.com/gofiber/fiber/v2"
	"huffy/database"
)

func CreateNoticeItems(notices []database.Notice) []Item {
	items := make([]Item, 0, len(notices))
	for _, notice := range notices {
		items = append(items, Item{
			Title:       notice.Title,
			Description: notice.Date,
			ImageUrl:    "https://t1.kakaocdn.net/openbuilder/sample/img_002.jpg",
			Link: &Link{
				Web: notice.Link,
			},
		})
	}
	return items
}

func CreateNoticeResponse(items []Item) fiber.Map {
	response := fiber.Map{
		"version": "2.0",
		"template": TemplateData{
			Outputs: []Output{
				{
					ListCard: &ListCard{
						Header: Header{
							Title: "공지사항",
						},
						Items: items,
						Buttons: []Button{
							{
								Action:     "webLink",
								Label:      "더 보기",
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

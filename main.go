package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"huffy/crawler"
	"huffy/database"
	"time"
)

type Template struct {
	Outputs []Output `json:"outputs"`
}

type Output struct {
	ListCard ListCard `json:"listCard"`
}

type ListCard struct {
	Header Header `json:"header"`
	Items  []Item `json:"items"`
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

func main() {
	db, err := database.InitDB("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := crawler.CrawlData(db)
				if err != nil {
					log.Error("Crawling Failed: ", err)
				} else {
					log.Info("Crawling Successful")
				}
			}
		}
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		articles, err := database.ReadData(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch articles",
			})
		}

		items := make([]Item, 0, len(articles))
		for _, article := range articles {
			items = append(items, Item{
				Title:       article.Title,
				Description: "설명",
				ImageUrl:    "https://t1.kakaocdn.net/openbuilder/sample/img_002.jpg",
				Link: Link{
					Web: "https://naver.com",
				},
			})
		}

		response := fiber.Map{
			"version": "2.0",
			"template": Template{
				Outputs: []Output{
					{
						ListCard: ListCard{
							Header: Header{
								Title: "챗봇 관리자센터를 소개합니다.",
							},
							Items: items,
						},
					},
				},
			},
		}

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}

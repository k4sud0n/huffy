package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"huffy/crawler"
	"huffy/database"
	"huffy/template"
	"os"
	"time"
)

func dbExist(fileName string, db *sql.DB) {
	_, error := os.Stat(fileName)
	if os.IsNotExist(error) {
		crawler.CrawlData(db)
	} else {
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
	}
}

func main() {
	db, err := database.InitDB("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	api := app.Group("/api")

	//dbExist("data.db", db)
	crawler.CrawlData(db)

	api.Get("/menu/today", func(c *fiber.Ctx) error {
		articles, err := database.ReadData(db)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch articles",
			})
		}

		items := template.CreateItems(articles)
		response := template.CreateResponse(items)

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"huffy/database"
)

func main() {
	db, err := database.InitDB("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//crawler.CrawlData(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		articles, err := database.ReadData(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch articles",
			})
		}

		return c.JSON(articles)
	})

	log.Fatal(app.Listen(":3000"))
}

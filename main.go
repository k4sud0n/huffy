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
		crawler.GetMenu(db)
		crawler.GetNotice(db)
	} else {
		crawler.GetMenu(db)
		crawler.GetNotice(db)
		go func() {
			ticker := time.NewTicker(24 * time.Hour)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					err := crawler.GetNotice(db)
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

	dbExist("data.db", db)

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/notice", func(c *fiber.Ctx) error {
		notices, err := database.ReadNotice(db)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch notices",
			})
		}

		items := template.CreateNoticeItems(notices)
		response := template.CreateNoticeResponse(items)

		return c.JSON(response)
	})

	api.Get("/menu/today", func(c *fiber.Ctx) error {
		parameter := c.Query("name")
		if parameter == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name parameter is required",
			})
		}

		// 어문관
		if parameter == "eomungwan" {
			response := template.CreateEomungwanResponse()
			return c.JSON(response)
		}

		// 바비든든
		if parameter == "babidundun" {
			response := template.CreateBabidundunResponse()
			return c.JSON(response)
		}

		menus, err := database.ReadMenu(db, parameter)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch menus",
			})
		}

		responseText := template.CreateMenuItem(menus)
		response := template.CreateMenuResponse(responseText)

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}

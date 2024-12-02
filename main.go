package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"huffy/crawler"
	"huffy/database"
	"huffy/template"
	"os"
	"time"
)

func checkDBExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func main() {
	var db *sql.DB
	var err error

	// 데이터베이스 파일 확인
	if checkDBExist("data.db") {
		fmt.Println("Database file found. Initializing DB connection...")
		db, err = database.InitDB("data.db") // 파일이 존재하면 DB 초기화
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// 예약된 크롤링 작업 실행
		scheduleCrawlingTask(db)
	} else {
		fmt.Println("Database file not found. Running initial crawling tasks...")
		// 임시 DB 생성 후 초기 크롤링 실행
		db, err = database.InitDB("data.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// 초기 크롤링 실행
		runCrawlingTask(db)
	}

	// Fiber 서버 설정
	app := fiber.New()
	api := app.Group("/api")

	api.Post("/notice", func(c *fiber.Ctx) error {
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

	api.Post("/menu/today", func(c *fiber.Ctx) error {
		parameter := c.Query("name")
		if parameter == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name parameter is required",
			})
		}

		menus, err := database.ReadMenu(db, parameter)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch menus",
			})
		}

		item := template.CreateMenuItem(menus)
		response := template.CreateMenuResponse(item)

		return c.JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}

func scheduleCrawlingTask(db *sql.DB) {
	go func() {
		// 현재 시간과 다음 실행 시간(매일 00:01) 계산
		now := time.Now()
		nextRun := time.Date(
			now.Year(), now.Month(), now.Day(),
			0, 1, 0, 0, now.Location(), // 매일 00:01
		)

		if now.After(nextRun) { // 이미 00:01을 지나쳤다면 다음 날로 설정
			nextRun = nextRun.Add(24 * time.Hour)
		}

		timeUntilNextRun := time.Until(nextRun)

		// 첫 번째 대기 후 작업 실행
		time.Sleep(timeUntilNextRun)
		runCrawlingTask(db)

		// 이후 매일 24시간 간격으로 실행
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			<-ticker.C
			runCrawlingTask(db)
		}
	}()
}

func runCrawlingTask(db *sql.DB) {
	fmt.Println("Running crawling task at:", time.Now().Format("2006-01-02 15:04:05"))
	// 크롤링 로직 호출
	if err := crawler.GetMenu(db); err != nil {
		log.Error("Failed to crawl menu:", err)
	}
	if err := crawler.GetNotice(db); err != nil {
		log.Error("Failed to crawl notice:", err)
	}
}

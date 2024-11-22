package database

import (
	"database/sql"
	"fmt"
	"time"
)

type Article struct {
	ID    int
	Title string
	Date  time.Time
}

func ReadData(db *sql.DB) ([]Article, error) {
	query := `SELECT ID, TITLE, "DATE" FROM ARTICLE`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Date)
		if err != nil {
			return nil, fmt.Errorf("error scanning article row: %w", err)
		}
		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return articles, nil
}

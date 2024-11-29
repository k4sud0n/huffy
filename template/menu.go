package template

import (
	"github.com/gofiber/fiber/v2"
	"huffy/database"
)

func CreateMenuItem(menus []database.Menu) string {
	var content string

	for _, menu := range menus {
		content = menu.Content
	}

	return content
}

func CreateMenuResponse(responseText string) fiber.Map {
	response := fiber.Map{
		"version": "2.0",
		"template": TemplateData{
			Outputs: []Output{
				{
					SimpleText: &SimpleText{
						Text: responseText,
					},
				},
			},
		},
	}

	return response
}

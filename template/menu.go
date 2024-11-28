package template

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"huffy/database"
)

func CreateMenuItem(menus []database.Menu) string {
	var menuDate string
	var menuContent string
	var responseText string

	for _, menu := range menus {
		menuDate = menu.Date
		menuContent = menu.Content

		responseText += fmt.Sprintf("%s\n%s", menuDate, menuContent)
	}

	return responseText
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

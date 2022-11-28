package responses

import "github.com/gofiber/fiber/v2"

type Responsefilter struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Total   int        `json:"total"`
	Data    *fiber.Map `json:"data"`
}

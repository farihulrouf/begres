package middleware

import (
	helper "begres/helpers"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	clientToken := c.Get("token")
	if clientToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unathenticated header",
		})
	}

	claims, err := helper.ValidateToken(clientToken)
	if err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "ineternal server error",
		})
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("uid", claims.Uid)
	c.Set("user_type", claims.User_type)

	return c.Next()
}

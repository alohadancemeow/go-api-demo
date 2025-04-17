package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func checkMiddleware(c *fiber.Ctx) error {
	// Get the user from the context written by the middleware [jwtware]
	user := c.Locals("user").(*jwt.Token)

	// Get the user from the claims
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}

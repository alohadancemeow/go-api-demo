package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != member.Email || user.Password != member.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	// Set claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token":   t,
		"message": "Login successful",
	})

}

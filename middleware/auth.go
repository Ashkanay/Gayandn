package middleware

import (
	"fmt"

	c "gayandn/configration"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	fmt.Println("Start Protected middleware")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(c.Config("SECRET")),
		//ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	fmt.Println("Start jwtError middleware")
	if err.Error() == "Missing or malformed JWT" {
		fmt.Println("request is problem")
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	fmt.Println("request is ok")
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

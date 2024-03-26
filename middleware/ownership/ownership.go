package ownership

import (
	"fmt"
	"go_jwt/middleware"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func IsCurrentUser(c *fiber.Ctx) error {
	tokenString := strings.Join(c.GetReqHeaders()["Token"], "")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}
	token, err := middleware.VerifyToken(tokenString)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}
	// Access claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token claims")
	}
	fmt.Println("this is claims", claims)

	// Extract user ID from claims
	var userID string
	if id, ok := claims["id"]; ok {
		switch v := id.(type) {
		case string:
			userID = v
		default:
			return c.Status(fiber.StatusUnauthorized).SendString("User ID is not a string")
		}
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("User ID not found in token")
	}

	id := c.Params("id")
	if userID != id {
		return c.Status(fiber.StatusForbidden).SendString("You don't have permission")
	} else {
		return c.Next()
	}
}

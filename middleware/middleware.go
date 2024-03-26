package middleware

import (
	"fmt"
	"go_jwt/src/database"
	models "go_jwt/src/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func IsAdmin(c *fiber.Ctx) error {
	// Extract token from request headers
	tokenString := strings.Join(c.GetReqHeaders()["Token"], "")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}

	// Verify token
	token, err := VerifyToken(tokenString)
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
	// collection create
	dbId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}
	collection := database.GetCollection("admins")
	var admin models.Admins
	error := collection.FindOne(c.Context(), bson.M{"_id": dbId}).Decode(&admin)
	if error != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not Admin")
	}
	// Token is valid, proceed to next middleware or handler
	return c.Next()
}

func IsUser(c *fiber.Ctx) error {
	tokenString := strings.Join(c.GetReqHeaders()["Token"], "")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}
	// Verify token
	token, err := VerifyToken(tokenString)
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
	// collection create
	dbId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}
	collection := database.GetCollection("users")
	var user models.Users
	error := collection.FindOne(c.Context(), bson.M{"_id": dbId}).Decode(&user)
	if error != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("You are not Authorized user")
	}
	// Token is valid, proceed to next middleware or handler
	return c.Next()

}

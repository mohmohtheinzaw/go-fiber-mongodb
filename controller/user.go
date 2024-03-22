package controller

import (
	"go_jwt/src/database"
	models "go_jwt/src/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUsers(c *fiber.Ctx) error {
	collection := database.GetCollection("users")
	// find all admin
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}
	users := make([]models.Users, 0)
	if err = cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}
	return c.Status(200).JSON(users)
}

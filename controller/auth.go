package controller

import (
	"fmt"
	"go_jwt/src/database"
	models "go_jwt/src/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateToken(id string) string {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	// }

	return t
}

// func LoginAdmin(c *fiber.Ctx) error {
// 	collection := database.GetCollection("admins")
// 	var admin models.Admins
// 	email := c.FormValue("email")
// 	fmt.Println(email)
// 	if err := c.FormValue("email"); err != nil {
// 		fmt.Println(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
// 	}
// 	fmt.Println("this", c.FormValue("email"))
// 	err := collection.FindOne(c.Context(), bson.M{"email": "test@gmail.com"}).Decode(&admin)

// 	if err != nil {
// 		fmt.Print(err)
// 		//return c.Status(404).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	fmt.Println("admin is", admin)
// 	fmt.Println(admin.Id)
// 	token := GenerateToken(admin.Id.Hex())
// 	return c.Status(200).JSON(token)

// }

func LoginAdmin(c *fiber.Ctx) error {
	collection := database.GetCollection("admins")
	var admin models.Admins

	// Get email from the request body
	email := c.FormValue("email")
	fmt.Println(email)
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email is required"})
	}
	filter := bson.D{{"email", email}}

	// Find admin by email
	err := collection.FindOne(c.Context(), filter).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Admin not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	// Generate token
	token := GenerateToken(admin.Id.Hex())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

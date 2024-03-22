package controller

import (
	"context"
	"fmt"
	"go_jwt/src/database"
	models "go_jwt/src/model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateToken(_id string, name string) string {
	claims := jwt.MapClaims{
		"id":   _id,
		"name": name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}

func LoginAdmin(c *fiber.Ctx) error {
	collection := database.GetCollection("admins")
	admin := new(models.Admins)

	//Parse req.body and check validation

	if err := c.BodyParser(admin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	// Get email from the request body
	filter := bson.M{"email": admin.Email}
	var foundAdmin models.Admins

	err := collection.FindOne(context.Background(), filter).Decode(&foundAdmin)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Generate token
	fmt.Println(foundAdmin.Id, foundAdmin.Name, foundAdmin.Id.Hex())

	token := GenerateToken(foundAdmin.Id.Hex(), foundAdmin.Name)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func RegisterCustomer(c *fiber.Ctx) error {
	data := new(models.Users)
	data.CreatedAt = time.Now().UTC()

	// validate the request body
	if err := c.BodyParser(data); err != nil {
		fmt.Print(err, "this is error")
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}
	fmt.Print(data)

	collection := database.GetCollection("users")
	fmt.Print(collection)
	result, err := collection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Internal server error": err.Error()})
	}

	// return the inserted todo
	fmt.Print(result.InsertedID)
	return c.Status(200).JSON(fiber.Map{"inserted_id": result.InsertedID})
}

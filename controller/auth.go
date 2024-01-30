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

type Admin struct {
	Email string `json:"email" bson:"email"`
	//Password string `json:"password" bson:"password"`
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
	token := GenerateToken(admin.Id.Hex())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

package AdminControllers

import (
	"fmt"
	"go-test/src/database"
	models "go-test/src/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	type createAdminDto struct {
//		// Id       primitive.ObjectID `json:"omitempty"`
//		Name     string `json:"name,omitempty" validate:"required"`
//		Location string `json:"location,omitempty" validate:"required"`
//		Title    string `json:"title,omitempty" validate:"required"`
//	}

func Create(c *fiber.Ctx) error {
	data := new(models.Admins)
	data.CreatedAt = time.Now().UTC()

	// validate the request body
	if err := c.BodyParser(data); err != nil {
		fmt.Print(err, "this is error")
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}
	fmt.Print(data)

	collection := database.GetCollection("admins")
	fmt.Print(collection)
	result, err := collection.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Internal server error": err.Error()})
	}

	// return the inserted todo
	fmt.Print(result.InsertedID)
	return c.Status(200).JSON(fiber.Map{"inserted_id": result.InsertedID})
}

func GetAll(c *fiber.Ctx) error {
	collection := database.GetCollection("admins")
	// filter := json.M{}
	// opts := options.Find().SetSkip(0).SetLimit(100)

	// find all todos
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// convert cursor to slice
	todos := make([]models.Admins, 0)
	// fmt.Print(cursor.All(c.Context(), &todos))
	if err = cursor.All(c.Context(), &todos); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return todos
	return c.Status(200).JSON(todos)

}

func GetOneAdmin(c *fiber.Ctx) error {
	coll := database.GetCollection("admins")

	id := c.Params("id")

	//chang to object id
	dbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}
	filter := bson.M{"_id": dbId}
	fmt.Println(filter)
	var admin models.Admins

	err = coll.FindOne(c.Context(), filter).Decode(&admin)
	fmt.Println(admin)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(admin)

}

type UpdateAdmin struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required"`
}

func Update(c *fiber.Ctx) error {
	collection := database.GetCollection("admins")
	updateData := new(UpdateAdmin)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}
	//change to object id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	dbId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	fmt.Println(id)
	fmt.Println(updateData)
	result, err := collection.UpdateOne(c.Context(), bson.M{"_id": dbId}, bson.M{"$set": updateData})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update book",
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	dbId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}
	coll := database.GetCollection("admins")
	filter := bson.M{"_id": dbId}
	res, err := coll.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"deleted_count": res.DeletedCount})

}

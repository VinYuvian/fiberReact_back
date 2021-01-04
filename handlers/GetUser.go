package handlers

import (
	"github.com/VinYuvian/Fiber/database"
	"github.com/VinYuvian/Fiber/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(ctx *fiber.Ctx) error {
	var res models.User
	email := ctx.Params("email")
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	client := database.DbConnection()
	userCol := client.Database("testapp").Collection("users")
	err := userCol.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("Bad Request")
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

	/*	paramsId := ctx.Params("id")
	id,err :=strconv.Atoi(paramsId)
	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Can't parse the request",
		})
	}
	for _,handlers := range database.Users{
		if handlers.Id == id{
			ctx.Status(fiber.StatusOK).JSON(handlers)
			return nil
		}
	}
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message" : "data not found",
	})

 */


package handlers

import (
	"github.com/VinYuvian/Fiber/database"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(ctx *fiber.Ctx) error {
	client:=database.DbConnection()
	mail := ctx.Params("email")
	//id,err := strconv.Atoi(paramId)
	filter:=bson.D{primitive.E{Key:"email",Value:mail}}
	db:=client.Database("testapp").Collection("users")
	res,err:=db.DeleteOne(context.Background(),filter)
	if err!=nil{
		return ctx.Status(fiber.StatusBadRequest).JSON("Can't parse the request")
	}
	/*for i,handlers := range database.Users{
		if handlers.Id == id{
			database.Users = append(database.Users[:i],database.Users[i+1:]...)
		}
	}*/
	return ctx.Status(fiber.StatusOK).JSON(res)
}

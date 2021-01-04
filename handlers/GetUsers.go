package handlers

import (
	"github.com/VinYuvian/Fiber/database"
	"github.com/VinYuvian/Fiber/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(ctx *fiber.Ctx) error{
	client:=database.DbConnection()
	userCol:=client.Database("testapp").Collection("users")
	cur,err:=userCol.Find(context.TODO(),bson.D{})
	if err!=nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	var res []*models.User
	for cur.Next(context.TODO()){
		var elem models.User
		err:=cur.Decode(&elem)
		res=append(res,&elem)
		if err!=nil{
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
	/*return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Users":database.Users,
		//"jwt" : ctx.Locals("handlers").(*jwt.Token),*/
}

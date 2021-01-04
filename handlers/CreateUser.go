package handlers

import (
	"github.com/VinYuvian/Fiber/database"
	"github.com/VinYuvian/Fiber/models"
	"context"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error{
	/*file,err := ctx.FormFile("document")
	if err!=nil{
		return err
	}*/
	var body models.User
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "can't parse json",
		})
	}
	user := models.User{
		Id:        len(database.Users)+1,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Gender:    body.Gender,
		Email:     body.Email,
		Password:  body.Password,
	}
	client:=database.DbConnection()
	//defer database.DbConnection()
	db:=client.Database("testapp")
	userColl:=db.Collection("users")
	userRes,err:=userColl.InsertOne(context.Background(),user)
	if err!=nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}
	//database.Users = append(database.Users,handlers)
	//ctx.SaveFile(file,fmt.Sprintf("./uploads/%s",strconv.Itoa(handlers.Id)+file.Filename))
	return ctx.Status(fiber.StatusCreated).JSON(userRes)
}

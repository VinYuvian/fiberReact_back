package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VinYuvian/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}
	app := fiber.New()
	file, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file : %v", err)
	}
	defer file.Close()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "[${time}] ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HELLO")
	})
	api := app.Group("/api", logger.New())
	api.Post("/User/upload", handlers.UserUpload)
	//app.Get("/Users",handlers.AuthRequired(),handlers.GetUsers)
	api.Get("/Users", handlers.GetUsers)
	api.Post("/Signup", handlers.CreateUser)
	api.Get("/Users/:email", handlers.GetUser)
	api.Delete("/Users/:email", handlers.DeleteUser)
	api.Post("/Login", handlers.Login)
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

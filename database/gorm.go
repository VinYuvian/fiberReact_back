package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnection() *mongo.Client {
	//username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	db_name := os.Getenv("DATABASE")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel == nil {
		fmt.Printf("An error occured cancel")
	}
	connString := fmt.Sprintf("mongodb+srv://mongo:%s@cluster0.uocwv.mongodb.net/%s?retryWrites=true&w=majority", password, db_name)
	//connString := fmt.Sprintf("mongodb://%s:%s@app_mongo-db_1:27017/testapp", username, password)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Printf("An error occured creating client connection")
	}
	return client
}

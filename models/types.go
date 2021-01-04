package models

type User struct {
	Id        int `primitive.ObjectID`
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
	Gender    string `bson:"gender"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
}

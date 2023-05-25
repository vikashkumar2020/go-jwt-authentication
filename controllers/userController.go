package controllers

import (
	"go-jwt-authentication/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")	
var validate = validator.New()

func HashPassword(password string) string {
	return password
}

func VerifyPassword(hashedPassword string, password string) bool {
	return hashedPassword == password
}

func Signup() {

}

func Login() {
}

func GetUsers() {
}

func GetUser() {
}

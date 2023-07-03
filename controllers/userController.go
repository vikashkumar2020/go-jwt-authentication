package controllers

import (
	"context"
	"fmt"
	"go-jwt-authentication/database"
	"go-jwt-authentication/models"
	"log"
	"net/http"

	helper "go-jwt-authentication/helpers"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Signup() gin.HandlerFunc {
	return func (c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		validatonErr := validate.Struct(user)
		if validatonErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":validatonErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
		}
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer  cancel()
		if err!=nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the phone"})
		}

		if count>0{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number is invalid"})
		}

		user.Created_at, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));
		user.Updated_at, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339));
		user.ID =  primitive.NewObjectID();
		str := user.ID.Hex()
		user.User_id = &str;
		token, refreshtoken, _ := helper.GenerateAllTokens(*user.Email,*user.First_name,*user.Last_name, *user.User_type, &user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshtoken

		resultInsertionNumber, insertErr :=userCollection.InsertOne(ctx, user)

		if insertErr !=nil{
			msg:=fmt.Sprintf("USER item was not created")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,resultInsertionNumber)
	}
}

func Login() {
}

func GetUsers() {
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")
		if err:= helper.MatchUserTypeToUid(c, userId); err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx,bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}

package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/oriastanjung/01-simple-restAPI/api/users/model"
	"github.com/oriastanjung/01-simple-restAPI/config"
	"github.com/oriastanjung/01-simple-restAPI/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection = db.GetCollection(db.DB, "users")
var validate = validator.New()

func CreateOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user model.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			responses := config.Response{
				Message: "Bad Request 1",
				Data:    err.Error(),
			}
			c.IndentedJSON(http.StatusBadRequest, responses)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			responses := config.Response{
				Message: "Bad Request 2",
				Data:    validationErr.Error(),
			}
			c.IndentedJSON(http.StatusBadRequest, responses)
			return
		}

		isEmailUnique, err := model.IsEmailUnique(user.Email)
		if err != nil {
			responses := config.Response{
				Message: "Bad Request ",
				Data:    err.Error(),
			}
			c.IndentedJSON(http.StatusBadRequest, responses)
			return
		}

		if !isEmailUnique {
			responses := config.Response{
				Message: "Bad Request ",
				Data:    "Email is not unique",
			}
			c.IndentedJSON(http.StatusBadRequest, responses)
			return
		}

		newUser := model.User{
			// Id:    primitive.NewObjectID(), // Set a new ObjectID
			Name:  user.Name,
			Email: user.Email,
		}

		result, err := usersCollection.InsertOne(ctx, newUser)

		if err != nil {
			responses := config.Response{
				Message: "Bad Request 3",
				Data:    err.Error(),
			}
			c.IndentedJSON(http.StatusInternalServerError, responses)
			return
		}

		responses := config.Response{
			Message: "success",
			Data:    result,
		}
		c.IndentedJSON(http.StatusCreated, responses)

	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []model.User
		defer cancel()

		results, err := usersCollection.Find(ctx, bson.M{})
		if err != nil {
			responses := config.Response{
				Message: "Bad Request ",
				Data:    err.Error(),
			}
			c.IndentedJSON(http.StatusBadRequest, responses)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var oneUser model.User

			if err = results.Decode(&oneUser); err != nil {
				responses := config.Response{
					Message: "Error ",
					Data:    err.Error(),
				}
				c.JSON(http.StatusInternalServerError, responses)
				return
			}

			users = append(users, oneUser)
		}

		responses := config.Response{
			Message: "success",
			Data:    users,
		}
		c.IndentedJSON(http.StatusOK, responses)

	}
}

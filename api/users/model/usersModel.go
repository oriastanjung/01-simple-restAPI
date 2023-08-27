package model

import (
	"context"

	"github.com/oriastanjung/01-simple-restAPI/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" validate:"required"`
}

// Check if email is unique
func IsEmailUnique(email string) (bool, error) {
	collection := db.GetCollection(db.DB, "users")

	filter := bson.M{
		"email": email,
	}
	ctx := context.TODO()

	var existingUser User

	err := collection.FindOne(ctx, filter).Decode(&existingUser)
	// fmt.Println(err)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Email is unique
			return true, nil
		}
		return false, err
	}

	// Email is not unique
	return false, nil
}

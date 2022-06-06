package repositories

import (
	"context"
	"time"

	"github.com/Cracker-TG/crboard/database"
	"github.com/Cracker-TG/crboard/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	FindOneUser() (*models.UserModel, error)
	// Create(user *models.UserModel)
}

type UserRepository struct{}

// var collection *mongo.Collection = .OpenCollection("users")

var umodel = new(models.UserModel)

func (ur UserRepository) FindOneUser() (*models.UserModel, error) {
	var dbI *database.DBinstance = &database.DBinstance{}
	var collection = dbI.OpenCollection("users")
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"username": "admin"}).Decode(umodel)
	defer cancel()
	return umodel, err
}

// func (ur UserRepository) Create(user *models.UserModel) {
// var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// result, err := ucollection.InsertOne(ctx, user)
// if err != nil {
// log.Fatalln(err)
// }

// log.Fatalln(result)
// defer cancel()
// }

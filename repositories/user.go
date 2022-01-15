package repositories

import (
	"context"
	"log"
	"time"

	"github.com/Cracker-TG/portfolio-service/database"
	"github.com/Cracker-TG/portfolio-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInteface interface {
	FindOneUser(filter *bson.D) (*models.UserModel, error)
	Create(user *models.UserModel)
}

type UserRepository struct{}

var ucollection *mongo.Collection = database.OpenCollection(database.Client, "users")
var umodel = new(models.UserModel)

func (ur UserRepository) FindOneUser(filter *bson.D) (*models.UserModel, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	err := ucollection.FindOne(ctx, filter).Decode(umodel)
	defer cancel()
	return umodel, err
}

func (ur UserRepository) Create(user *models.UserModel) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	result, err := ucollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(result)
	defer cancel()
}

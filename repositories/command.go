package repositories

import (
	"context"
	"time"

	"github.com/Cracker-TG/portfolio-service/database"
	"github.com/Cracker-TG/portfolio-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommandRepositoryInteface interface {
	FindOneCommand(filter *bson.D) (*models.CommandModel, error)
	Create(command *models.CommandModel) (*mongo.InsertOneResult, error)
}

type CommandRepository struct{}

var ccollection *mongo.Collection = database.OpenCollection(database.Client, "commands")
var cmodel = new(models.CommandModel)

func (cr CommandRepository) FindOneCommand(filter *bson.D) (*models.CommandModel, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	err := ccollection.FindOne(ctx, filter).Decode(cmodel)
	defer cancel()
	return cmodel, err
}

func (cr CommandRepository) Create(command *models.CommandModel) (*mongo.InsertOneResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	result, err := ccollection.InsertOne(ctx, command)
	defer cancel()

	if err != nil {
		return result, err
	}

	return result, nil
}

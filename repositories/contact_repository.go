package repositories

import (
	"context"
	"time"

	"github.com/Cracker-TG/portfolio-service/database"
	"github.com/Cracker-TG/portfolio-service/forms"
	"github.com/Cracker-TG/portfolio-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IContactRepository interface {
	Create(data *forms.Contact) (*mongo.InsertOneResult, error)
}

type ContactRepository struct{}

func (cr ContactRepository) Create(data *forms.Contact) (*mongo.InsertOneResult, error) {
	var dbI *database.DBinstance = &database.DBinstance{}
	var collection = dbI.OpenCollection("contacts")

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	contactData := models.ContactModel{
		ID:        primitive.NewObjectID(),
		Subject:   data.Subject,
		Email:     data.Email,
		Message:   data.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := collection.InsertOne(ctx, contactData)
	defer cancel()

	if err != nil {
		return result, err
	}

	return result, nil
}

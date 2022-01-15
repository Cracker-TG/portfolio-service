package command

import (
	"net/http"
	"time"

	"github.com/Cracker-TG/portfolio-service/controllers"
	"github.com/Cracker-TG/portfolio-service/forms"
	"github.com/Cracker-TG/portfolio-service/models"
	"github.com/Cracker-TG/portfolio-service/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommandController struct{}

var commandRepository repositories.CommandRepositoryInteface = new(repositories.CommandRepository)
var commandForm = new(forms.Command)

func (cstr CommandController) Create(c *gin.Context) {
	err_bind := c.ShouldBindJSON(&commandForm)
	if err_bind != nil {
		controllers.ErrResponseWithCode(c, http.StatusBadRequest, err_bind.Error())
		return
	}
	data := models.CommandModel{ID: primitive.NewObjectID(), Name: commandForm.Name, Detail: commandForm.Detail, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result, err := commandRepository.Create(&data)

	if err != nil {
		controllers.ErrResponseWithCode(c, http.StatusBadRequest, err.Error())
		return
	}

	payload := map[string]interface{}{
		"id": &result.InsertedID,
	}

	controllers.SuccessResponse(c, payload)
}

func (cstr CommandController) Show(c *gin.Context) {
	err_bind := c.ShouldBindJSON(&commandForm)
	if err_bind != nil {
		controllers.ErrResponseWithCode(c, http.StatusBadRequest, err_bind.Error())
		return
	}

	query := bson.D{{Key: "username", Value: commandForm.Name}}

	find, err := commandRepository.FindOneCommand(&query)

	if err != nil {
		controllers.ErrResponseWithCode(c, 404, err.Error())
		return
	}

	payload := map[string]interface{}{
		"id": find.ID,
	}

	controllers.SuccessResponse(c, payload)
}

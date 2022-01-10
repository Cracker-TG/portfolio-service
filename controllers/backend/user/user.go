package user

import (
	"net/http"
	"time"

	"github.com/Cracker-TG/portfolio-service/controllers"
	"github.com/Cracker-TG/portfolio-service/forms"
	"github.com/Cracker-TG/portfolio-service/helpers"
	"github.com/Cracker-TG/portfolio-service/repositories"
	"github.com/Cracker-TG/portfolio-service/securityTokens"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct{}

type response struct {
	UID  primitive.ObjectID
	Code string
}

var userRepository repositories.UserRepositoryInteface = new(repositories.UserRepository)
var loginForm = new(forms.Login)
var h helpers.HelpersInteface = new(helpers.Helpers)
var pasto securityTokens.PasetoInteface = new(securityTokens.PasetoMaker)

func (u UserController) Login(c *gin.Context) {
	err_bind := c.ShouldBindJSON(&loginForm)
	if err_bind != nil {
		controllers.ErrResponseWithCode(c, http.StatusBadRequest, err_bind.Error())
		return
	}

	query := bson.D{{Key: "username", Value: loginForm.Username}}
	findUser, err := userRepository.FindOneUser(&query)

	if err == mongo.ErrNoDocuments {
		controllers.ErrResponseWithCode(c, 401, "unauthorized")
		return
	}

	if err != nil {
		controllers.ErrResponseWithCode(c, 500, err.Error())
		return
	}

	decryp := h.Decryption(findUser.Password, loginForm.Password)

	if !decryp {
		controllers.ErrResponseWithCode(c, 401, "unauthorized")
		return
	}

	expTime := 2 * time.Minute
	token, errtoken := pasto.CreateToken(&findUser.ID, &findUser.UserName, &expTime)

	if errtoken != nil {
		controllers.ErrResponseWithCode(c, 500, errtoken.Error())
		return
	}

	payload := map[string]interface{}{
		"id":    findUser.ID,
		"token": token,
		"expri": time.Now().Add(expTime),
	}

	controllers.SuccessResponse(c, payload)
}

func (u UserController) Info(c *gin.Context) {

	query := bson.D{{Key: "username", Value: "superAdmin"}}
	findUser, err := userRepository.FindOneUser(&query)

	if err == mongo.ErrNoDocuments {
		controllers.ErrResponseWithCode(c, 401, "unauthorized")
		return
	}

	payload := map[string]interface{}{
		"id":       findUser.ID,
		"username": findUser.UserName,
	}

	if err != nil {
		controllers.ErrResponseWithCode(c, 500, err.Error())
		return
	}

	controllers.SuccessResponse(c, payload)
}

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

	query := bson.D{{"username", loginForm.Username}}
	findUser, err := userRepository.FindOneUser(&query)

	if err == mongo.ErrNoDocuments {
		controllers.ErrResponseWithCode(c, 401, "ununauthorized")
		return
	}

	if err != nil {
		controllers.ErrResponseWithCode(c, 500, err.Error())
		return
	}

	cp := h.Decryption(findUser.Password, loginForm.Password)

	if !cp {
		controllers.ErrResponseWithCode(c, 401, "ununauthorized")
		return
	}

	expTime := 2 * time.Second
	token, errtoken := pasto.CreateToken(&findUser.UserName, &expTime)

	if errtoken != nil {
		controllers.ErrResponseWithCode(c, 500, errtoken.Error())
		return
	}

	payload := map[string]interface{}{
		"id":    cp,
		"token": token,
	}

	controllers.SuccessResponse(c, payload)
}

package users

import (
	"fmt"
	"net/http"

	"github.com/Cracker-TG/portfolio-service/controllers"
	"github.com/Cracker-TG/portfolio-service/forms"
	"github.com/gin-gonic/gin"
)

type ContactController struct{}

// var contactRepository repositories.IContactRepository = new(repositories.ContactRepository)

func (controller ContactController) Store(c *gin.Context) {

	var contactForm = new(forms.Contact)
	err_bind := c.ShouldBindJSON(&contactForm)
	if err_bind != nil {
		controllers.ErrResponseWithCode(c, http.StatusBadRequest, err_bind.Error())
		return
	}

	fmt.Print(contactForm)
	// controllers.SuccessResponse(c, {})
}

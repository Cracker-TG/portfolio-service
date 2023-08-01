package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Create a new validator instance
var customValidator = validator.New()

func init() {
	// Register the custom email validation function
	customValidator.RegisterValidation("email", validateEmail)
}

func validateEmail(fl validator.FieldLevel) bool {
	// Get the field value as a string
	email, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	// Your custom email validation logic goes here
	// You can use a regular expression, the "checkmail" library, or any other validation method.

	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

// GetCustomValidator returns the custom validator instance
func GetCustomValidator() *validator.Validate {
	return customValidator
}

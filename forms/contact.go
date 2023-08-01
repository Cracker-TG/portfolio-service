package forms

type Contact struct {
	Subject string `json:"subject"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message"`
}

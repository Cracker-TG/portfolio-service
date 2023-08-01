package forms

type Contact struct {
	Subject string `json:"subject"`
	Email   string `json:"email" binding:"required"`
	Message string `json:"detail"`
}

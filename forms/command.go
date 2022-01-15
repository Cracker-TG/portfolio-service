package forms

type Command struct {
	Name   string `json:"name" binding:"required"`
	Detail string `json:"detail" binding:"required"`
}

package schema

type CreateBook struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" validate:"required" binding:"required"`
	Author      string `json:"author" validate:"required" binding:"required"`
	Publication string `json:"publication," validate:"required" binding:"required"`
}

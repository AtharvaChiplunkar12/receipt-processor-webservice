package dtos

type ItemDTO struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,floatString"`
}

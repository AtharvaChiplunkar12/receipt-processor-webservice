package dtos

type ItemDTO struct {
    ShortDescription string `json:"shortDescription" validate:"required"`
    Price           float64 `json:"price" validate:"required,floatString"`
}
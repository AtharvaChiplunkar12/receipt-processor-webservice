package models

// Reqired during database interaction

type Item struct {
	ShortDescription string `json:"short_description"`
	Price            float64    `json:"price"`
}

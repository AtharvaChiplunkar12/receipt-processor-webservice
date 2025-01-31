package models

type Receipt struct {
	Retailer      string `json:"retailer"`
	PurschaseDate string `json:"purchse_date"`
	PurschaseTime string `json:"purchse_time"`
	Total         int    `json:"total"`
	Items         []Item `json:"items"`
}

package unit_tests


import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"src/models"	
	"github.com/stretchr/testify/assert"
)

func TestGetPoints_Success(t *testing.T) {
	router := RouterSetup()
	receiptJSON := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
			},{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
			},{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
			},{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
			},{
			"shortDescription": "Klarbrunn 12-PK 12 FL OZ",
			"price": "12.00"
			}
		],
		"total": "35.35"
		}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var processResponse models.ReceiptResponse
	_ = json.Unmarshal(res.Body.Bytes(), &processResponse)

	getReq, _ := http.NewRequest(http.MethodGet, "/receipts/"+processResponse.ID+"/points", nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusOK, getRes.Code)
	var pointsResponse models.PointsResponse
	err := json.Unmarshal(getRes.Body.Bytes(), &pointsResponse)
	assert.Nil(t, err)
	assert.NotNil(t, pointsResponse.Points)
	assert.Equal(t, pointsResponse.Points, 28)
}

func TestGetPointsMorning_Success(t *testing.T) {
	router := RouterSetup()
	receiptJSON := `{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var processResponse models.ReceiptResponse
	_ = json.Unmarshal(res.Body.Bytes(), &processResponse)

	getReq, _ := http.NewRequest(http.MethodGet, "/receipts/"+processResponse.ID+"/points", nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusOK, getRes.Code)
	var pointsResponse models.PointsResponse
	err := json.Unmarshal(getRes.Body.Bytes(), &pointsResponse)
	assert.Nil(t, err)
	assert.NotNil(t, pointsResponse.Points)
	assert.Equal(t, pointsResponse.Points, 15)
}
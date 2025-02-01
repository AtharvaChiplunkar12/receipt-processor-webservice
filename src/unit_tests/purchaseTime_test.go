package unit_tests


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPurchaceTime_EmptyField(t *testing.T) {
	router := RouterSetup()

	receiptJSON := `{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	
}

func TestPurchaceTime_InvalidFields(t *testing.T) {
	router := RouterSetup()

	// invalid date
	receiptJSON := `{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-23",
    "purchaseTime": "25:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	
}
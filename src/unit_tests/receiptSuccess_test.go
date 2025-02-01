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


func TestReceipt_Success(t *testing.T) {
	router := RouterSetup()

	receiptJSON := `{
        "retailer": "Target",
        "purchaseDate": "2022-01-01",
        "purchaseTime": "13:01",
        "items": [
            {"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"}
        ],
        "total": "35.35"
    }`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var response models.ReceiptResponse
	err := json.Unmarshal(res.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.ID)
}





func TestProcessReceipt_MissingFields(t *testing.T) {
	router := RouterSetup()

	receiptJSON := `{
		
	}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	
}

func TestGetPoints_NonExistentID(t *testing.T) {
	router := RouterSetup()

	getReq, _ := http.NewRequest(http.MethodGet, "/receipts/invalidID/points", nil)
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	assert.Equal(t, http.StatusNotFound, getRes.Code)

	var errorResponse map[string]string
	err := json.Unmarshal(getRes.Body.Bytes(), &errorResponse)
	assert.Nil(t, err)
	assert.Equal(t, "No receipt found for that ID", errorResponse["error"])
}

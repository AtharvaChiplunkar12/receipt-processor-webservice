package unit_tests


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestItems_EmptyField(t *testing.T) {
	router := RouterSetup()

	receiptJSON := `{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "01:01",
    "total": "2.65",
    "items": [
       
    ]
}`

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)

	
}

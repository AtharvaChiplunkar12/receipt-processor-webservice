package processes

import (
	"math"
	"regexp"
	"src/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

// In-memory store for receipts
var receiptStore = make(map[string]int)

func ReceiptProcessing(receipt models.Receipt) string {
	receiptID := uuid.New().String()
	points := CalculatePoints(receipt)
	receiptStore[receiptID] = points

	return receiptID
}

func GetPointsFromStore(id string) (int, bool) {
	points, exists := receiptStore[id]
	return points, exists
}

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	points += countAlphanumeric(receipt.Retailer)
	total := receipt.Total
	if total == float64(int(total)) {
		points += 50
	}
	if int(total*100)%25 == 0 {
		points += 25
	}

	points += len(receipt.Items) / 2 * 5
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price := item.Price
			points += int(math.Ceil(price * 0.2))
		}
	}
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 == 1 {
		points += 6
	}

	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}
	return points
}

func countAlphanumeric(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1))
}

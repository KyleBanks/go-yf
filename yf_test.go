package yf

import (
	"testing"
	"time"
)

func init() {
	DebugLogging = true
}

func TestGetStock(t *testing.T) {
	symbol := "GOOG"
	start := time.Now().AddDate(0, -3, 0)
	end := time.Now()

	err := GetStock(symbol, start, end)
	if err != nil {
		t.Fatal(err)
	}
}

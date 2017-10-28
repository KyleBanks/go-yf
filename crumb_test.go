package yf

import "testing"

func TestGetCrumb(t *testing.T) {
	crumb, err := getCrumb("AAPL")
	if err != nil {
		t.Fatal(err)
	}

	if len(crumb) == 0 {
		t.Fatal("Expected non-empty crumb")
	}
}

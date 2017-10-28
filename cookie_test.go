package yf

import (
	"testing"
)

func TestGetCookie(t *testing.T) {
	cookie, err := getCookie("GOOG")
	if err != nil {
		t.Fatal(err)
	}

	if cookie == nil {
		t.Fatal("Expected non-nil cookie")
	} else if cookie.Name != cookieName {
		t.Fatalf("Unexpected cookie Name, expected=%v, got=%v", cookieName, cookie.Name)
	} else if len(cookie.Value) == 0 {
		t.Fatal("Expected non-empty cookie value")
	}
}

package yf

import (
	"net/http"
	"strings"
)

const (
	cookieUrl       = "https://finance.yahoo.com/quote/"
	cookieHeader    = "set-cookie"
	cookieName      = "cookie"
	cookieDelimiter = ";"
)

// addCookie applies the Yahoo Finance cookie to the provided request.
//
// This cookie is required for all API requests, and must be retrieved from a public Yahoo Finance page
// using the `getCookie` function.
func addCookie(symbol string, req *http.Request) error {
	cookie, err := getCookie(symbol)
	if err != nil {
		return err
	}

	req.AddCookie(cookie)
	req.Header.Add(cookieName, cookie.Value)

	return nil
}

// getCookie retrieves the cookie value from Yahoo Finance.
//
// This cookie is required for all API requests, and must be retrieved from a public Yahoo Finance page.
func getCookie(symbol string) (*http.Cookie, error) {
	url := cookieUrl + symbol
	debug("Fetching Cookie: " + url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	value := resp.Header.Get(cookieHeader)
	value = strings.Split(value, cookieDelimiter)[0]
	debug("Cookie Value: " + value)
	return &http.Cookie{Name: cookieName, Value: value}, nil
}

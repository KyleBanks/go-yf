package yf

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const crumbUrl = "https://finance.yahoo.com/quote/"

var (
	crumbRegexp = regexp.MustCompile(`CrumbStore":{"crumb":".[^"]+"}`)
)

// getCrumb returns the 'crumb' param required for all Yahoo Finance API requests by parsing it
// from a public Yahoo Finance page.
func getCrumb(symbol string) (string, error) {
	url := crumbUrl + symbol
	debug("Fetching Crumb: " + url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// This is extremely brittle...
	crumbJson := crumbRegexp.FindString(string(body))
	components := strings.Split(crumbJson, `"`)
	crumb := components[len(components)-2]
	debug("Crumb Value: " + crumb)
	return crumb, nil
}

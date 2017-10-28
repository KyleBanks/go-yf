package yf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

const (
	yahooFinanceApiTemplate = "https://query1.finance.yahoo.com/v7/finance/download/{{.Symbol}}?period1={{.Start}}&period2={{.End}}&interval={{.Interval}}&events={{.Events}}&crumb={{.Crumb}}"
	defaultInterval         = "1d"
	defaultEvents           = "history"
)

var (
	DebugLogging = false

	yahooFinanceApi = template.Must(template.New("YF-API").Parse(yahooFinanceApiTemplate))
)

type params struct {
	Symbol string

	Start    int64
	End      int64
	Interval string

	Events string

	Crumb string
}

func GetStock(symbol string, start, end time.Time) error {
	req, err := getRequest(symbol, start, end)
	if err != nil {
		return err
	}

	var c http.Client
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Unexpected Response: %d: %v", resp.StatusCode, string(body))
	}

	return nil
}

func getRequest(symbol string, start, end time.Time) (*http.Request, error) {
	p := params{
		Symbol:   symbol,
		Start:    start.Unix(),
		End:      end.Unix(),
		Interval: defaultInterval,
		Events:   defaultEvents,
	}

	crumb, err := getCrumb(symbol)
	if err != nil {
		return nil, err
	}
	p.Crumb = crumb

	var url bytes.Buffer
	if err := yahooFinanceApi.Execute(&url, p); err != nil {
		return nil, err
	}
	debug(url.String())

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	if err := addCookie(symbol, req); err != nil {
		return nil, err
	}
	debug(fmt.Sprintf("Headers: %v", req.Header))

	return req, nil
}

func debug(str string) {
	if !DebugLogging {
		return
	}

	fmt.Println(str)
}

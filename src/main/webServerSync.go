package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	for _, symbol := range stockSymbols {
		resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		quote := new(QuoteResponse)
		xml.Unmarshal(body, &quote)

		fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

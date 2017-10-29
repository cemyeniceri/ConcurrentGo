package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
	"os"
	"encoding/csv"
	"strconv"
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

type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerNumber int
	PartNumber string
	Quantity int

	UnitCost   float64
	UnitPrice  float64
}

func extract()[]*Order {
	result :=[]*Order{}

	f,_ := os.Open("./orders.txt")
	defer f.Close()
	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		result = append(result, order)
	}
	return result;
}

func transform([]*Order[]*Order)
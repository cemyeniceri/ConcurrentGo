package main

import (
	"fmt"
	"time"
	"os"
	"encoding/csv"
	"strconv"
)

func main() {

	start := time.Now()

	extractChannel := make(chan *OrderAsync)
	transformChannel := make(chan *OrderAsync)
	doneChannel := make(chan bool)

	go extractAsync(extractChannel)
	go transformAsync(extractChannel, transformChannel)
	go loadAsync(transformChannel, doneChannel)

	<-doneChannel

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s", elapsed)
}

type ProductAsync struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type OrderAsync struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int

	UnitCost  float64
	UnitPrice float64
}

func extractAsync(ch chan *OrderAsync) {

	f, _ := os.Open("./orders.txt")
	defer f.Close()
	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(OrderAsync)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		ch <- order
	}

	close(ch)
}

func transformAsync(extractChannel, transformChannel chan *OrderAsync) {
	f, _ := os.Open("./productList.txt")
	defer f.Close()
	r := csv.NewReader(f)

	records, _ := r.ReadAll()
	productList := make(map[string]*ProductAsync)
	for _, record := range records {
		product := new(ProductAsync)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[1], 64)
		productList[product.PartNumber] = product
	}

	for o := range extractChannel {
		time.Sleep(3 * time.Millisecond)
		o.UnitCost = productList[o.PartNumber].UnitCost
		o.UnitPrice = productList[o.PartNumber].UnitPrice
		transformChannel <- o
	}

	close(transformChannel)
}

func loadAsync(transformChannel chan *OrderAsync, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")

	for o := range transformChannel {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
			o.PartNumber, o.Quantity,
			o.UnitCost, o.UnitPrice,
			o.UnitCost*float64(o.Quantity),
			o.UnitPrice*float64(o.Quantity))
	}
	doneChannel <- true
}

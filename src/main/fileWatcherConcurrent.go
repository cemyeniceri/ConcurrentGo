package main

import (
	"os"
	"io/ioutil"
	"time"
	"encoding/csv"
	"strings"
	"strconv"
	"fmt"
)

const watchedPathConcurrent = "./source"

func main() {
	for true {
		d, _ := os.Open(watchedPathConcurrent)
		// By providing negative value to the function, we are asking it to return as many files as it finds
		files, _ := d.Readdir(-1)

		for _, file := range files {
			filePath := watchedPathConcurrent + "/" + file.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.RemoveAll(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoiceConcurrent := new(InvoiceConcurrent)
					invoiceConcurrent.Number = r[0]
					invoiceConcurrent.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoiceConcurrent.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoiceConcurrent.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("Received InvoiceConcurrent '%v' for $%.2f and submitted for processing\n", invoiceConcurrent.Number, invoiceConcurrent.Amount)
				}
			}(string(data))
		}
		d.Close()
	}
}

type InvoiceConcurrent struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}

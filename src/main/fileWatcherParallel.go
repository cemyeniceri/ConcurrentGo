package main

import (
	"os"
	"io/ioutil"
	"time"
	"encoding/csv"
	"strings"
	"strconv"
	"fmt"
	"runtime"
)

const watchedPath = "./source"

func main() {
	runtime.GOMAXPROCS(4)
	for true {
		d, _ := os.Open(watchedPath)
		// By providing negative value to the function, we are asking it to return as many files as it finds
		files, _ := d.Readdir(-1)

		for _, file := range files {
			filePath := watchedPath + "/" + file.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.RemoveAll(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoiceParallel := new(InvoiceParallel)
					invoiceParallel.Number = r[0]
					invoiceParallel.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoiceParallel.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoiceParallel.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("Received InvoiceParallel '%v' for $%.2f and submitted for processing\n", invoiceParallel.Number, invoiceParallel.Amount)
				}
			}(string(data))
		}
		d.Close()
	}
}

type InvoiceParallel struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}

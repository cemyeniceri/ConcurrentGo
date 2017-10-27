package main

import (
	"fmt"
)

func main() {

	po := new(PurchaseOrder)
	po.Value = 42.27
	ch := make(chan *PurchaseOrder)

	go SavePO(po, ch)

	newPo := <-ch
	fmt.Printf("PO: %v", newPo)
}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234
	callback <- po
}

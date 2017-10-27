package main

import (
	"fmt"
	"errors"
	"time"
)

func main() {
	po := new(PurchaseOrderPromise)
	po.Value = 42.27

	SavePOforPromise(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrderPromise)
		fmt.Printf("Purchase Order saved with ID: %d\n", po.Number)
		return nil
	}, func(err error) {
		fmt.Printf("Failed to save Purchase Order: " + err.Error() + "\r")
	}).Then(func(obj interface{}) error {
		fmt.Println("Second promise success")
		return nil
	}, func(err error) {
		fmt.Println("Second promise failed: " + err.Error())
	})

	fmt.Scanln()
}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

type PurchaseOrderPromise struct {
	Number int
	Value  float64
}

func SavePOforPromise(po *PurchaseOrderPromise, shouldFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		time.Sleep(2 * time.Second)
		if shouldFail {
			result.failureChannel <- errors.New("failed to save purchase order")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()
	return result
}

func (promise *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	timeout := time.After(1 - time.Second)

	go func() {
		select {
		case obj := <-promise.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-promise.failureChannel:
			failure(err)
			result.failureChannel <- err
		case <-timeout:
			failure(errors.New("promise timed out"))
		}
	}()

	return result
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	inputPipe := make(chan int)
	/* inputPipe channel act as input pipe and generate odd numbers except(2) async */
	go generate(inputPipe)
	for {
		/* except 2(initial prime) get first number of filter as prime and print it on console */
		prime := <-inputPipe
		fmt.Println(prime)
		/* outputPipe channel act as a filter for prime and test values */
		outputPipe := make(chan int)
		go filter(inputPipe, outputPipe, prime)
		inputPipe = outputPipe
	}
}

//goroutine1
func generate(ch chan int) {
	ch <- 2
	for i := 3; ; i = i+2 {
		ch <- i
	}
}

//goroutine2
func filter(in, out chan int, prime int) {
	for value := range in {
		if value%prime != 0 {
			out <- value
		}
	}
}

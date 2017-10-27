package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
			}()
		}
	}
}

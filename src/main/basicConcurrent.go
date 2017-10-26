package main

import "time"

func main() {

	goDur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(goDur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(goDur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

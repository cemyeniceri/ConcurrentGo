package main

import "time"

func main(){
	go func(){
		println("Hello")
	}()

	go func(){
		println("Go")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
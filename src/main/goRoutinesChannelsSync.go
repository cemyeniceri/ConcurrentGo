package main

import (
	"fmt"
	"runtime"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 1)
	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	ch := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ch <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh<-msg
				fmt.Print(msg)
				<-ch
			}()
		}
	}
}

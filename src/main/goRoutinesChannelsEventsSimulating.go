package main

import (
	"fmt"
)

func main() {

	btn := MakeButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One : " + msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two : " + msg)
		}
	}()

	btn.TriggerEvent("click", "Button clicked!")
	btn.RemoveEventListener("click", handlerTwo)
	btn.TriggerEvent("click", "Button clicked again!")

	fmt.Scanln()
}

type Button struct {
	eventListeners map[string][] chan string
}

func MakeButton() *Button{
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (receiver *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := receiver.eventListeners[event]; present {
		receiver.eventListeners[event] = append(receiver.eventListeners[event], responseChannel)
	} else {
		receiver.eventListeners[event] = []chan string{responseChannel}
	}
}

func (receiver *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := receiver.eventListeners[event]; present {
		for idx := range receiver.eventListeners[event] {
			if receiver.eventListeners[event][idx] == listenerChannel {
				receiver.eventListeners[event] = append(receiver.eventListeners[event][:idx], receiver.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (receiver *Button) TriggerEvent(event string, response string) {
	if _, present := receiver.eventListeners[event]; present {
		for _, handler := range receiver.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

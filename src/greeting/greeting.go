package greeting

import (
	"fmt"
)


type Salutation struct {
	Name string
	Greeting string
}

type Printer func(string)


func Greet(salutations []Salutation, do Printer, isFormal bool) {

	for _, salutation := range salutations {
		message, alternateMessage := CreateMessage(salutation.Name, salutation.Greeting, "yo")
		if prefix:= GetPrefix(salutation.Name); isFormal {
			do(prefix + message + " " + alternateMessage)
		} else {
			do(message)
		}
	}
}

func GetPrefix(name string) (prefix string){

	prefixMap := map[string]string{
		"Bob" : "Mr ",
		"Joe" : "Dr ",
		"Amy" : "Dr ",
		"Mary" : "Mrs ",
	}

	prefixMap["Joe"] = "Jr "
	delete(prefixMap, "Mary")

	return prefixMap[name]
}

func CreateMessage(name string, greeting ...string) (string, string){
	return greeting[0]  + " " + name, "HEY!" + name + " " + greeting[1]
}

func CreateCustMessage(custom string) Printer{
	return func(i string) {
		fmt.Println(i, custom)
	}
}

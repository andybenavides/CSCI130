package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	if prefix := GetPrefix(salutation.name); isFormal {
		do(prefix + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr "
	case name == "Joe", name == "Amy", len(name) == 10:
		prefix = "Dr "
	case name == "Mary":
		prefix = "Mrs "
	default:
		prefix = "Dude "
	}
	return
}

func TypeSwitchTest(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
	}
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {
	var s = Salutation{"Andy", "Hello"}
	Greet(s, CreatePrintFunction("?"), true)
	TypeSwitchTest(123)
}

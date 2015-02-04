package main

import "fmt"

type Fall struct {
	leaves   string
	coldWind string
}

func main() {

	var message1 = Fall{leaves: "fallen leaves.", coldWind: "very cold wind."}

	fmt.Println("Fall brings ")
	fmt.Println(message1.leaves)

	fmt.Println("Fall brings ")
	fmt.Println(message1.coldWind)

	var testMessage string = "Test Pointers"
	var pointer *string = &testMessage
	fmt.Println(testMessage, *pointer)
}

package main

import "fmt"

//loop
func looperfunc(age int) {
	for i := 1; i <= age; i++ {
		if i == 23 {
			fmt.Print("I am ")
			fmt.Println(age)
		} else {
			fmt.Print("I am older than ")
			fmt.Println(i)
		}
	}
}

//slice
func slicesfunc(fname string) {
	name := []string{fname}
	name = append(name, "Benavides")
	fmt.Println(name)
}

//map
func mapfunc(person string) (age int) {
	ageMap := map[string]int{
		"Andy":   23,
		"Cassie": 20,
		"Robert": 17,
	}

	if value, exists := ageMap[person]; exists {
		fmt.Println(ageMap[person])
		return value
	}

	fmt.Println("That person does not exist.")

	return
}

func main() {
	fmt.Println("--------------------------")
	fmt.Println("---------Looping----------")

	age := 23
	looperfunc(age)

	fmt.Println("---------Slicing----------")

	fname := "Andy"
	slicesfunc(fname)

	fmt.Println("---------Mapping----------")
	mapfunc("John")

	fmt.Println("--------------------------")
}

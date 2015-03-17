package main

import (
 "fmt"
 "math"
)

type welcome struct{
  name string
  greeting string
}


type a_welcome string

func (s a_welcome) hello1(){ // variable s is of type a_welcome
  fmt.Println("Welcome!")
}

func (c welcome) hello2(){
  fmt.Println("Welcome, " + c.name)
}

// Method with a pointer
func (c *welcome) diff_name(new_name string){
  c.name=new_name
}

// Methods with return values

func (c welcome) hello3() string {
return "Welcome " + c.name
}

func main() {

 person:= welcome{"Andy", " how are you?"}

 person.hello1()

 message := welcome = "Welcome people."
 fmt.Println(message)
}

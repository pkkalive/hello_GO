package main

import "fmt"

func main() {
	var name string = "Kumar"
	var age int = 30
	var isStudent bool = true
	var pi float64 = 3.14159
	fmt.Println("Hello, World!")
	fmt.Println("Hello, " + name + "!")
	fmt.Println("You are " + fmt.Sprint(age) + " years old.")
	fmt.Println("Are you a student? " + fmt.Sprint(isStudent))
	fmt.Println("The value of π is " + fmt.Sprint(pi))
}

package main

import "fmt"

func main() {
	fmt.Println("hello world")

	//variable declaration
	var name string = "ashik"
	var age = 10
	var job string
	job = "Engineer"
	friend := "hashir"
	var a, b, c, d, e, f int = 1, 2, 3, 4, 5, 6
	x, y := -1, "ashik"
	job = "doctor"
	const aim = "sw engineer" //cannot be reassigned
	fmt.Print("Hello", "World \n")
	fmt.Println("Hello", "World")
	fmt.Printf("a has value %v and type %T \n", a, a)
	fmt.Println(name, age, job, friend, a, b, c, d, e, f, x, y)

}

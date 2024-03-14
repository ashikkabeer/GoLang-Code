package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "ashik"
	person1.age = 12
	person1.job = "swe"
	person1.salary = 3000

	person2.name = "joel"
	person2.age = 55
	person2.job = "ds"
	person2.salary = 5124

	fmt.Println(person2.name)

	//Pass Struct as Function Arguments
	printPerson(person1)
	printPerson(person2)
}
func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

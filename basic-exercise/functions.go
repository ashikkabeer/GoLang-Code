package main

import "fmt"

func main() {
	x, y := myFunction(5, "hello")
	fmt.Println(x, y)
	fmt.Println(factorial_recursion(4))
}

// Here, myFunction() returns one integer (result) and one string (txt1):
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y
	return
}
func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

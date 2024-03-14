package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//average of two numbers
	const height int = 23
	const weight float32 = 53.2
	fmt.Println((float32(height) + weight) / 2)

	//condtional statements
	const num = 5
	if num == 6 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
}

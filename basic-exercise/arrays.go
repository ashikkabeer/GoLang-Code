package main

import "fmt"

func main() {
	var arr = [...]int{9, 5, 7, 6, 2, 3, 5, 1, 4, 5, 6, 7}
	arr2 := [5]int{1: 20, 4: 30}
	fmt.Println(arr2)
	nums := [...]string{"ashik", "kabeer"}
	fmt.Println(nums)
	fmt.Println(arr)
	fmt.Println(len(arr2))
}

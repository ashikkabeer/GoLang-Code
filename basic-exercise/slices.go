package main

import "fmt"

func main() {
	/*
	* unlike arrays, the length of a slice can grow and shrink as you see fit.
	 */
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	mySlice := []int{}
	fmt.Println(len(mySlice))

	//creating slice from array
	arr := [6]int{1, 1, 2, 3, 4, 5}
	slice := arr[2:5]
	fmt.Print(slice)

	//creating slice with make()
	slice2 := make([]int, 5, 10)
	fmt.Print(slice2)
	slice3 := append(slice, slice2...)
	fmt.Print(slice3)

}

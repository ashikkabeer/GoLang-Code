package main

import "fmt"

func main() {
	//loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i := 0; i <= 100; i += 10 {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fruits := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	for i, val := range fruits {
		fmt.Printf("%v\t%v\n", i, val)
	}

	//condtions
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
	time = 5
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")

	}
	day := 4
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

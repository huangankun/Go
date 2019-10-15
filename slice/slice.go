package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func main() {
	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Println("切片是空的")
	}

	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(number)

	fmt.Println("number == ", number)
	fmt.Println("numbrt[1:4]==", number[1:4])
	fmt.Println("number[:3]==", number[:3])
	fmt.Println("number[4:]", number[4:])

	numbers2 := make([]int, 0, 5)
	printSlice(numbers2)
	numbers3 := number[:2]
	printSlice(numbers3)

	numbers4 := number[2:5]
	printSlice(numbers4)
}

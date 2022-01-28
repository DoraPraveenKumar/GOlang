package main

import "fmt"

func main() {
	a := 0
	n := 0
	fmt.Println("Enter the length of array")
	fmt.Scanln(&n)
	arr := make([]int, n)
	fmt.Println("Enter the elements of array")
	for i := 0; i < n; i++ {
		fmt.Scanln(&a)
		arr = append(arr, a)
	}
	fmt.Println(arr)
}

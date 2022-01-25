package main

import "fmt"

func main() {
	var N int
	fmt.Println("Enter a number")
	fmt.Scanln(&N)
	var arr = make([]int, N)
	fmt.Println("Enter ", N, " elements")
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", arr[i])
	}
	var sum int = 0
	f := arr[0]
	for i := 0; i < N; i++ {
		if arr[i] == 0 {
			sum += f
		} else if f > arr[i] {
			sum += f + arr[i]
		} else {
			f = arr[i]
		}
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	var R1, C1, R2, C2 int
	fmt.Println("Enter row and column in matrix 1:")
	fmt.Scanf("%d %d", &R1, &C1)
	fmt.Println("Enter row and column in matrix 2:")
	fmt.Scanf("%d %d", &R2, &C2)
	if C1 != R2 {
		fmt.Println("Matrix multiplication not possible")
	} else {
		arr1 := make([][]int, R1)
		for i := 0; i < R1; i++ {
			arr1[i] = make([]int, C1)
		}
		for i := 0; i < R1; i++ {
			for j := 0; j < C1; j++ {
				fmt.Scanf("%d", &arr1[i][j])
			}
		}
		fmt.Println("Enter the elements of matrix :2")
		arr2 := make([][]int, R1)
		for i := 0; i < R1; i++ {
			arr2[i] = make([]int, C1)
		}
		for i := 0; i < R2; i++ {
			for j := 0; j < C2; j++ {
				fmt.Scan("%d", &arr2[i][j])
			}
		}
		fmt.Println("Matrix: 1")
		for i := 0; i < R1; i++ {
			for j := 0; j < C1; j++ {
				fmt.Print(arr1[i][j])
			}
			fmt.Println(" ")
		}
		fmt.Println("Matrix: 2")
		for i := 0; i < R2; i++ {
			for j := 0; j < C2; j++ {
				fmt.Print(arr2[i][j])
			}
			fmt.Println(" ")
		}
	}

}

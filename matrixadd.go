package main

import "fmt"

func main() {
	var R, C int
	fmt.Println("Enter number of rows and coumns of array")
	fmt.Scanf("%d", &R)
	fmt.Scanf("%d", &C)
	fmt.Println("Enter the elments of matrix :1")
	//var arr1 [R][C]int
	arr1 := make([][]int, R)
	for i := 0; i < R; i++ {
		arr1[i] = make([]int, C)
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Scanf("%d", &arr1[i][j])
		}
	}
	fmt.Println("Enter the elements of matrix :2")
	//var arr2 [R][C]int
	arr2 := make([][]int, R)
	for i := 0; i < R; i++ {
		arr1[i] = make([]int, C)
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Scan("%d", &arr2[i][j])
		}
	}
	fmt.Println("Matrix: 1")
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Print(arr1[i][j])
		}
		fmt.Println(" ")
	}
	fmt.Println("Matrix: 2")
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Print(arr2[i][j])
		}
		fmt.Println(" ")
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			arr1[i][j] += arr2[i][j]
		}
	}
	fmt.Println("The sum is :")
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Print(arr2[i][j])
		}
		fmt.Println(" ")
	}
}

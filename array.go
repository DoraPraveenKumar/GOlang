package main

import "fmt"

func main() {
	var arr [4]int
	//elements one by one
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
	a := [4]int{0, 1, 2, 3}
	fmt.Println(a)
	//unknown length array (ellipses)
	c := [...]int{1, 2, 3, 4, 5, 6, 7} //when dont know length of array
	fmt.Println(c)
	//iterating over array
	for x := 0; x < len(c); x++ {
		fmt.Printf("%d\n", c[x])
	}
	//iterating over array using range keyword
	//range can only be used over a list or array
	for index, value := range c {
		fmt.Println("element at ", index, " ; ", value)
	}
	//2d array
	b := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(b)
	for index1, value1 := range b {
		for index2, value2 := range value1 {
			fmt.Println("element at ", index1, " ", index2, " value1 ", value1, " value2 ", value2)
		}
	}
	//SLICE doesnot have a fixed size
	//better than array

	var slice []int //slice is name of slice
	for i := 0; i < 6; i++ {
		slice = append(slice, 19)
	}
	fmt.Println(slice)
	s := make([]int, 5, 15) //min 5 max15
	fmt.Println(s)
}

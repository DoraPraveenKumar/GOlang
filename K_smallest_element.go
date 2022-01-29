/*
Kth smallest element
Given an array arr[] and an integer K where K is smaller than size of array, the task is to find the Kth smallest element in the given array. It is given that all array elements are distinct.
Example 1:
Input:
N = 6
arr[] = 7 10 4 3 20 15
K = 3
Output : 7
Explanation :
3rd smallest element in the given
array is 7.
*/

package main

import (
	"fmt"
	"sort"
)

var arr []int
var n int
var k int

func kthSmallestElement() int {

	/*for i:=0; i<len(arr); i++{

	    for j:=i+1; j<len(arr); j++{

	        if arr[i]> arr[j]{

	            b:= arr[i]
	            arr[i]=arr[j]
	            arr[j]=b
	        }
	    }
	}*/
	//alternate for above snipet is
	sort.Ints(arr)
	return arr[k-1]
}

func main() {
	var a int
	fmt.Println("Enter the length of array")
	fmt.Scanln(&n)
	fmt.Println("Enter the elements of array")
	for i := 0; i < n; i++ {

		fmt.Scanln(&a)
		arr = append(arr, a)
	}
	fmt.Println("Enter the value of k")
	fmt.Scanln(&k)
	fmt.Println("The kth smallest element is:", kthSmallestElement())
}

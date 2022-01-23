package main

import (
	"errors"
	"fmt"
)

func greetings(name string) (string, error) {

	if name == "" {
		return name, errors.New("name is empty")
	}
	return "Hello," + name, nil
}

//multiple parameters unknown
func add(num ...int) int {
	sum := 0
	//for value := range num it runs from 0-(n-1) so 1-6
	for _, value := range num {
		sum += value
	}
	return sum
}

func main() {
	greetings, err := greetings("CipherText")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greetings)
	}
	fmt.Println("the sum of 1,2,3,4,5,6 is", add(1, 2, 3, 4, 5, 6, 1))
}

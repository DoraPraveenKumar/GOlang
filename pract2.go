package main

import (
	"errors"
	"fmt"
)

func greetings(name string) (string, error) {

	if name == "" {
		return name, errors.New(" name can not be empty")
	} else {
		return "Hello," + name, nil

	}
}

func add(num ...int) float64 {
	sum := 0.0
	for _, value := range num {
		fmt.Println(value)
		sum += float64(value)

	}
	return sum
}

func isPalindrome(number int) bool {
	var remainder, temp int
	var reverse int = 0

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {
		return true
	} else {
		return false
	}

}

func factorial(number int) (fact int, err error) {
	fact = 1
	if number < 0 {
		err = errors.New("Number has to be +ve")
		return
	} else if number == 0 {
		return
	}
	for number > 0 {
		fact *= number
		number--
	}
	return
}
func main() {
	/*greeting, galti := greetings("Cipher")

	if galti != nil {
		fmt.Println(galti)
	} else {
		fmt.Println(greeting)
	}
	*/
	//fmt.Println("the sum of 1,2,3,4,5 is ", add(1, 2, 3, 4, 5))
	//fmt.Println("the sum of 1,2,3,4,5,6,7,8 is ", add(1, 2, 3, 4, 5, 6, 7, 8))
	var number int
	fmt.Println(" Enter any positive number")
	fmt.Scanln(&number)
	fmt.Println(" Result of palindrom", isPalindrome(number))

	fmt.Println(" Enter any positive number for factorial")
	fmt.Scanln(&number)
	result, err := factorial(number)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The factorial is", result)
	}

}

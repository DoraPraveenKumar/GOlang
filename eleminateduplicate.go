/*
Problem #3

Given a string, say S, print it in reverse manner eliminating the repeated characters and spaces.

Example 1:

Input: S = "GEEKS FOR GEEKS"
Output: "SKEGROF"

Example 2:

Input: S = "I AM AWESOME"
Output: "EMOSWAI"

-------------------
Problem#3 has been asked in the interview by samsung.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) (reverse string) {
	for i := len(s) - 1; i >= 0; i++ {
		if string(s[i]) == " " {
			reverse += string(s[i])
		}
	}
	return
}
func main() {
	fmt.Println("Enter a string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	fmt.Println(reverse(line))
}

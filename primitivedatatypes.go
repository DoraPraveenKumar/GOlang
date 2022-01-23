package main

import "fmt"

func main() {
	var n bool
	fmt.Printf("%v, %T", n, n)
	var a int = 3
	b := 5
	var c uint = 4
	fmt.Printf("%v, %T", c, c)
	fmt.Println(a << 3)
	fmt.Println(b % 23)
	var d complex64 = 23 + 2i
	fmt.Printf("%v, %T", d, d)
	fmt.Printf("\n%v, %T", real(d), imag(d))
	s := "String it is"
	fmt.Println("\n", s, s[2], string(s[2]))
}

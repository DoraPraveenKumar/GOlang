package main

import "fmt"

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	c := 1
	for i := 2; ; i++ {
		p := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				p = false
			}
		}
		if p == true {
			if c == n {
				return i, p
			} else {
				c++
			}
		}
	}
}

func main() {
	fmt.Println(Nth(0))
}

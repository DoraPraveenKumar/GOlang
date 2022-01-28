package main

import (
	"fmt"
)

func main() {
	fruits := map[int]string{
		1: "apple",
		2: "orange",
		3: "mango",
	}
	fmt.Println(fruits)
	fruits[4] = "grapes"
	fruits[6] = "avacado"
	fruits[5] = "lichi"
	fmt.Println(fruits)

	for no, name := range fruits {
		fmt.Println("name", name, "number", no)
	}
	a, ok := fruits[10] //only can check for keys
	fmt.Println(a, ok)

}

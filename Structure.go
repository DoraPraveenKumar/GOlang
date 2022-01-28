package main

import "fmt"

type Doc struct {
	number int
	name   string
	old    oldDoc
	//func call()
}
type oldDoc struct {
	name string
}

func (i Doc) call(num string) {
	fmt.Println(" ......" + num)
}
func main() {
	a := Doc{}
	a.name = "SSFF"
	a.number = 132566
	a.call("noo")
	fmt.Println(a)
	b := Doc{
		number: 12,
		name:   "jb",
		old: oldDoc{
			name: "Ramu kaka",
		},
	}
	fmt.Println(b)
}

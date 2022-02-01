package main

import "fmt"

func SumMultiples(limit int, divisors ...int)(sum int) {
    for _,j:= range divisors{
        for i:=0;i<limit;i++{
        	if j%i==0{
            	sum+=1
        	}
    	}
    }
	return
}

func main(){
	var divisors [] int={3,5}
	fmt.println(SumMultiples( 20, divisors))
}
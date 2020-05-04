package main

import (
	"fmt"
)

func factorial(n int) int{
	var total int
	total = n
	for n > 1 {
		n = n - 1
		total = total * n
		
	}
	
	return total
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(factorial(n))
}


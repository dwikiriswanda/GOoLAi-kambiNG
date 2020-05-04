package main

import (
	"fmt"
)

func cetakDeret(n *int) {
	for *n != 1 {
		if *n%2 == 0 {
			*n = *n/2
			fmt.Print(*n, " ")
		} else if *n%2 == 1 {
			*n = 3*(*n) + 1
			fmt.Print(*n, " ")
		}	
	}
}

func main () {

	var nInput int
	
	for nInput != 1{
		fmt.Scan(&nInput)
		fmt.Print(nInput, " ")
		cetakDeret(&nInput)
	}
	
	
}

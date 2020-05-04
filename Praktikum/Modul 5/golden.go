package main

import (
	"fmt"
)

func main() {
	
	var n int
	
	
	fmt.Print("Masukkan nilai n: ")
	fmt.Scanln(&n)
	
	
	for i := 1; i <= n; i++ { 
	i = n
	
	n = n - 1 + n - 2
	
	fmt.Print(n, (" "))
	}
}
package main

import (
	"fmt"
)

func main() {

	var uang int
	fmt.Print("Nilai uang : ")
	fmt.Scanln(&uang)

	for uang%25 != 0 {

		fmt.Printf("%v bukan nilai yang valid\n", uang)
		fmt.Print("Nilai uang : ")
		fmt.Scanln(&uang)

	}
	fmt.Printf("%v nilai yang valid\n", uang)
}

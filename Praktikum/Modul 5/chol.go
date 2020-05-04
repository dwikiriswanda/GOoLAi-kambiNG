package main

import (
	"fmt"
)

func main() {

	var(
	hdl, ldl, trigly, total float64
	sehat bool
	)

	fmt.Print("HDL: ")
	fmt.Scanln(&hdl)
	fmt.Print("LDL: ")
	fmt.Scanln(&ldl)
	fmt.Print("Triglyserida: ")
	fmt.Scanln(&trigly)
	
	total = hdl + ldl + trigly * 0.2
	fmt.Println("Total: ", total)
	
	sehat = trigly / hdl < 2 || hdl / ldl > 0.4 || total / hdl < 3.5
	fmt.Print("Sangat sehat: ", sehat)

}
//DONE
package main

import "fmt"

func main(){
	var c int
	var r int
	var f int
	var k float64
	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&c)
	r = c * 4 / 5
	fmt.Println("Nilai Reamur nya adalah: ", r)
	f = c * 9 / 5 + 32
	fmt.Println("Nilai Fahrenheit nya adalah: ", f)
	k = (float64(f) + 459.67) * float64(5)/ float64(9)
	fmt.Printf("Nilai Kelvin nya adalah: " + "%.0f",k)
}
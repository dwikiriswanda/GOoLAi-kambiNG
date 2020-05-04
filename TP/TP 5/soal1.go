/*
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
*/

package main

import (
	"fmt"
)

func fungsiF(x float64) float64 {
	hasil := x * x
	return hasil
}

func fungsiG(x float64) float64 {
	hasil := x - 2
	return hasil
}

func fungsiH(x float64) float64 {
	hasil := x + 1
	return hasil
}

func fungsiFoGoH(x float64) float64 {
	hasil := fungsiF(fungsiG(fungsiH(x)))
	return hasil
}

func main() {

	fmt.Println("=======================================")
	fmt.Println("	PROGRAM FUNGSI")
	fmt.Println("=======================================")

	var x float64
	fmt.Print("Masukkan nilai x = ")
	fmt.Scanln(&x)
	fmt.Printf("f(%.2f) = %.2f\n", x, fungsiF(x))
	fmt.Printf("g(%.2f) = %.2f\n", x, fungsiG(x))
	fmt.Printf("h(%.2f) = %.2f\n", x, fungsiH(x))
	fmt.Printf("(fogoh) (%.2f) = %.2f\n", x, fungsiFoGoH(x))

	fmt.Println("=======================================")
	fmt.Println("Nama : Muhamad Dwiki Riswanda")
	fmt.Println("NIM  : 1302194015")
	fmt.Println("=======================================")

}
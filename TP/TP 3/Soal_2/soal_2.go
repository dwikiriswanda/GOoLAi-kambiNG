/*
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
*/
package main

import (
	"fmt"
)

func main() {

	var nam float64
	var nmk string

	//Nama Program
	fmt.Println("=======================================")
	fmt.Println("	PROGRAM NILAI KULIAH ")
	fmt.Println("=======================================")
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	//Algoritma
	if nam > 80 && nam <= 100 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nmk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nmk = "C"
	} else if nam > 40 && nam <= 50 {
		nmk = "D"
	} else if nam >= 0 && nam <= 40 {
		nmk = "E"
	} else {
		nmk = "Tidak diketahui"
	}
	fmt.Println("Nilai mata kuliah :", nmk)
	// Nama dan NIM
	fmt.Println("=======================================")
	fmt.Println("Nama	: Muhamad Dwiki Riswanda")
	fmt.Println("NIM	: 1302194015")
	fmt.Println("=======================================")
}

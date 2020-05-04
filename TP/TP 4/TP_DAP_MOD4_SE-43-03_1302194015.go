/*
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
*/
package main

import (
	"fmt"
)	 

func main() {
	var (
		ganjilGenap int
		penyebut, s1, s2, selisih1, selisih2 float64
	)

	fmt.Println("=======================================")
	fmt.Println("	PROGRAM Leibniz  ")
	fmt.Println("=======================================")

	penyebut = 1
	ganjilGenap = 1
	selisih1 = 1
	selisih2 = 1

	for selisih1 * 4 >= 0.00001 && selisih2 * 4 >= 0.00001 {
		s1 = s2

		if ganjilGenap % 2 == 0 {
			s2 -= (1 / penyebut)
		} else {
			s2 += (1 / penyebut)
		}

		selisih1 = s1 - s2
		selisih2 = s2 - s1

		if selisih1 < 0 {
			selisih1 *= -1
		}
		
		if selisih2 < 0 {
			selisih2 *= -1
		}
		penyebut += 2
		ganjilGenap++
	}
	fmt.Printf("Hasil pada PI: %.10f\n", s1 * 4)
	fmt.Printf("Hasil pada PI: %.10f\n", s2 * 4)

	fmt.Println("Pada i ke: ", ganjilGenap)

	fmt.Println("=======================================")
	fmt.Println("Nama	: Muhamad Dwiki Riswanda")
	fmt.Println("NIM	: 1302194015")
	fmt.Println("=======================================")

}
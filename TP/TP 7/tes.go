/*
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
*/

package main

import (
	"fmt"
)

const N = 2019

type RecType struct {
	f1 string
	f2 int
	f3 float64
}
type ArrType [N]RecType

func rmax(data ArrType) float64 {
	var max float64

	max = -99999999

	for i := 0; i < N; i++ {
		if max < data[i].f3 {
			max = data[i].f3
		}
	}
	return max
}

func imin(data ArrType) int {
	var min int

	min = 99999999

	for i := 0; i < N; i++ {
		if min > data[i].f2 {
			min = data[i].f2
		}
	}
	return min
}

func found(data ArrType, key string) bool {
	var temu bool

	temu = false

	for i := 0; i < N; i++ {
		if key == data[i].f1 {
			temu = true

		}
	}
	return temu
}

func pos(data ArrType, key string) int {
	var nilaiTengah int

	batasKiri := 0
	batasKanan := N
	found := false

	for (batasKiri < batasKanan) && !found {
		nilaiTengah = (batasKiri + batasKanan) / 2

		if data[nilaiTengah].f1 < key {
			batasKiri = nilaiTengah + 1
		} else if data[nilaiTengah].f1 > key {
			batasKanan = nilaiTengah
		} else {
			found = true
		}
	}

	if found == false {
		nilaiTengah = -1
	}
	return nilaiTengah
}

func main() {

	var dataArray ArrType
	var cari string

// output seharunya muncul index ke berapa ketika menginputkan huruf 
	dataArray[0].f1 = "A"
	dataArray[1].f1 = "B"
	dataArray[2].f1 = "C"
	dataArray[3].f1 = "D"
	dataArray[4].f1 = "E"
	dataArray[5].f1 = "F"
	dataArray[6].f1 = "G"
	dataArray[7].f1 = "H"
	dataArray[8].f1 = "I"
	dataArray[9].f1 = "J"

// output akan tetap 0 karena dari 2019 array hanya beberapa saja yang diisi, dimana array lainnya akan otomatis di isi dengan 0
	dataArray[10].f2 = 11111
	dataArray[11].f2 = 12222
	dataArray[12].f2 = 2333
	dataArray[13].f2 = 1444
	dataArray[14].f2 = 2555
	dataArray[15].f2 = 1666
	dataArray[16].f2 = 27777
	dataArray[17].f2 = 212
	dataArray[18].f2 = 4223
	dataArray[19].f2 = 11333

// output akan 66.6 karena nilai terbesar dari keseluruhan array (dimana array selain array dibawah nilainya terisi dengan 0)
	dataArray[20].f3 = 6.66
	dataArray[21].f3 = 5.0
	dataArray[22].f3 = 3.14
	dataArray[23].f3 = 8.23
	dataArray[24].f3 = 10.3
	dataArray[25].f3 = 66.6
	dataArray[26].f3 = 44.7
	dataArray[27].f3 = 24.1
	dataArray[28].f3 = 31.1
	dataArray[29].f3 = 51.22

	
	fmt.Println("=======================================")
	fmt.Println("	PROGRAM PENCARIAN	")
	fmt.Println("=======================================")


	fmt.Println("Nilai Real Terbesar: ", rmax(dataArray))
	fmt.Println("Nilai Integer Terkecil: ", imin(dataArray))

	fmt.Print("Masukkan Key: ")
	fmt.Scan(&cari)

	fmt.Println("Apakah Key ada pada array?", found(dataArray, cari))
	fmt.Println("Key ditemukan pada array ke-", pos(dataArray, cari))

	fmt.Println("=======================================")
	fmt.Println("Nama : Muhamad Dwiki Riswanda")
	fmt.Println("NIM  : 1302194015")
	fmt.Println("=======================================")
}

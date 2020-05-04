package main

import "fmt"

func main() {

	var tahun int
	//Nama program
	fmt.Println("=======================================")
	fmt.Println("	PROGRAM CEK TAHUN KABISAT")
	fmt.Println("=======================================")

	fmt.Print("Masukkan Tahun: ")
	fmt.Scanln(&tahun)
	// Jika angka tahun itu tidak habis dibagi 400, tidak habis dibagi 100 akan tetapi habis dibagi 4, maka tahun itu merupakan tahun kabisat
	if tahun%4 == 0 {
		// Jika angka tahun itu tidak habis dibagi 400 tetapi habis dibagi 100, maka tahun itu sudah pasti bukan merupakan tahun kabisat
		if tahun%100 == 0 {
			//Jika angka tahun itu habis dibagi 400, maka tahun itu sudah pasti tahun kabisat
			if tahun%400 == 0 {
				fmt.Println("Tahun", tahun, "yang anda masukkan")
				fmt.Println("adalah Tahun Kabisat")

			} else {
				fmt.Println("Tahun", tahun, "yang anda masukkan")
				fmt.Println("bukan merupakan Tahun Kabisat")
			}

		} else {
			fmt.Println("Tahun", tahun, "yang anda masukkan")
			fmt.Println("adalah Tahun Kabisat")
		}
		//Jika angka tahun tidak habis dibagi 400, tidak habis dibagi 100, dan tidak habis dibagi 4, maka tahun tersebut bukan merupakan tahun kabisat
	} else {
		fmt.Println("Tahun", tahun, "yang anda masukkan")
		fmt.Println("bukan merupakan Tahun Kabisat")
	}
	// Nama dan NIM
	fmt.Println("=======================================")
	fmt.Println("Nama	: Muhamad Dwiki Riswanda")
	fmt.Println("NIM	: 1302194015")
	fmt.Println("=======================================")
}
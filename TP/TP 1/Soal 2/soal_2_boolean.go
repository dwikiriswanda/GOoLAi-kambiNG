/*  
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
Tahun kabisat adalah tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100. Buatlah sebuah program yang menerima input sebuah bilangan bulat dan memeriksa apakah bilangan tersebut merupakan tahun kabisat (true) atau bukan (false) 
*/

package main

import "fmt"

func main(){
// Memberi nama variabel
var (
	kabisat bool
	tahun int
)

//Nama Program
fmt.Println("=======================================")
fmt.Println("	PROGRAM CEK TAHUN KABISAT ")
fmt.Println("=======================================")

//Algoritma
fmt.Print("Tahun: ")
fmt.Scanln(&tahun)
fmt.Print("Kabisat: ")
kabisat = (tahun%400) == (0) || (tahun%4) == (0) && (tahun%100) != (0)
fmt.Println(kabisat)
fmt.Println("")

fmt.Print("Tahun: ")
fmt.Scanln(&tahun)
fmt.Print("Kabisat: ")
kabisat = (tahun%400) == (0) || (tahun%4) == (0) && (tahun%100) != (0)
fmt.Println(kabisat)
fmt.Println("")

fmt.Print("Tahun: ")
fmt.Scanln(&tahun)
fmt.Print("Kabisat: ")
kabisat = (tahun%400) == (0) || (tahun%4) == (0) && (tahun%100) != (0)
fmt.Println(kabisat)

// Nama dan NIM
fmt.Println("=======================================")
fmt.Println("Nama	: Muhamad Dwiki Riswanda")
fmt.Println("NIM	: 1302194015")
fmt.Println("=======================================")

}
/*  
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
Telusuri program berikut dengan cara mengkompilasi dan mengekseksi program. Silakan masukkan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa yang sebenarnya yang dilakukan program tersebut?
*/

package main

import "fmt"

func main() {
// Memberi nama variabel
var(
satu, dua, tiga string
temp string
)

//Nama Program
fmt.Println("=======================================")
fmt.Println("	PROGRAM MENGGANTI URUTAN")
fmt.Println("=======================================")

//Algoritma
fmt.Print("Masukkan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukkan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukkan input string: ")
fmt.Scanln(&tiga)
fmt.Println("")
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = satu
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)

// Nama dan NIM
fmt.Println("=======================================")
fmt.Println("Nama : Muhamad Dwiki Riswanda")
fmt.Println("NIM  : 1302194015")
fmt.Println("=======================================")

}
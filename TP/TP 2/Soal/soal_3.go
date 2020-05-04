/*  
NIM: 1302194015
Nama: Muhamad Dwiki Riswanda
Telusuri program berikut dengan cara mengkompilasi dan mengekseksi program. Silakan masukkan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa yang sebenarnya yang dilakukan program tersebut?
*/
package main

import (
    "fmt"
)

func main(){
    var (
        kiri, kanan float64
        jumlah bool
    )

    //Nama Program
fmt.Println("=======================================")
fmt.Println("	PROGRAM TES BERAT KANTONG")
fmt.Println("=======================================")
 
    for selesai:=false; !selesai; {

        //input nilai kantong
        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        fmt.Scan(&kiri, &kanan)
        jumlah = kiri - kanan >= 9 || kanan - kiri >= 9

        //menampilkan nilai true / false
        fmt.Println("Sepeda motor pak Andi akan oleng: ", jumlah)

        //jika program melebihi 150 kg atau salah satu negatif
        selesai = kiri + kanan>=150 || kiri<0 || kanan<0
    
    }

    fmt.Println("Program selesai.")

    // Nama dan NIM
fmt.Println("=======================================")
fmt.Println("Nama : Muhamad Dwiki Riswanda")
fmt.Println("NIM  : 1302194015")
fmt.Println("=======================================")

}
package main 
import "fmt"
func main (){

var rataahmad, ratabadrun int 
var ahmad1, ahmad2, ahmad3 int
var badrun1, badrun2, badrun3 int

fmt.Scan(&ahmad1,&ahmad2,&ahmad3 )

fmt.Scan(&badrun1,&badrun2,&badrun3)

rataahmad = (ahmad1 + ahmad2 + ahmad3) / 3

ratabadrun = (badrun1 + badrun2 + badrun3) / 3

fmt.Print(rataahmad, " ")

fmt.Println(ratabadrun)

fmt.Println("Rata-rata Ahmad lebih baik dari Badrun? ", rataahmad > ratabadrun)

fmt.Println("Apakah Ahmad pemenang absolut?  ", ahmad1 > badrun2 && ahmad2 > badrun2 && ahmad3 > badrun3)

fmt.Println("Apakah Badrun pemenang absolut?  ", badrun1 > ahmad1 && badrun2 > ahmad2 && badrun3 > ahmad2)

//BONUS TANTANGAN
fmt.Println("BONUS TANTANGAN")
fmt.Println("Ahmad memenangkan kompetisi? ", ahmad1 > badrun2 && ahmad2 > badrun2 || ahmad3 > badrun3)
fmt.Println("Badrun memenangkan kompetisi? ", badrun1 > ahmad1 && badrun2 > ahmad2 || badrun3 > ahmad2)
fmt.Println("Tidak ada pemenang?  ", !(ahmad1 > badrun2 && ahmad2 > badrun2 || ahmad3 > badrun3 && badrun1 > ahmad1 && badrun2 > ahmad2 || badrun3 > ahmad2))


}
package main

import "fmt"

import "math/rand"

func main (){

	var seed int64
	var genapganjil, besarkecil string
	var dadu int

	
	fmt.Print("Angka rahasia: ")
	fmt.Scanln(&seed)
	rand.Seed(seed)

	fmt.Print("Anda: ")
	fmt.Scanln(&genapganjil, &besarkecil)

	dadu = rand.Intn(6)+1
	fmt.Println("Nilai dadu: ", dadu)

	if genapganjil == "genap" && dadu%2 == 0{
		if besarkecil == "besar" && dadu >= 4{
			fmt.Println("Anda menang")
		} else if besarkecil == "kecil" && dadu <4{
			fmt.Println("Anda menang")
		} else {
			fmt.Println("Anda kalah")
		}

	} else {
		if besarkecil == "besar" && dadu >= 4{
			fmt.Println("Anda menang")
		} else if besarkecil == "kecil" && dadu <4{
			fmt.Println("Anda menang")
		} else {
			fmt.Println("Anda kalah")
		}
	}

}
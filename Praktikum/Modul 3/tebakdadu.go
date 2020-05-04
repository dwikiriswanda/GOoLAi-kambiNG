package main

import "fmt"

import "math/rand"

func main (){

	var seed  int64
	var anda, dadang, dadu int
	

	fmt.Print("Angka rahasia: ")
	fmt.Scanln(&seed)
	rand.Seed(seed)
	
	
	fmt.Print("Anda: ")
	fmt.Scanln(&anda)
	
	
	dadang = rand.Intn(6)+1
	fmt.Println("Dadang: ", dadang)
	
	dadu = rand.Intn(6)+1
	fmt.Print("Nilai dadu ", dadu, ", ")
	
	if dadu == anda {
	
	fmt.Print("Pemenang adalah anda")
	
	} else if dadu == dadang {
	
	fmt.Print("Pemenang adalah dadang")
	
	} else if dadu == anda && dadu == dadang {
	
	fmt.Print("Seri")
	
	} else {
	fmt.Print("Tidak ada yang menang")
	}



}



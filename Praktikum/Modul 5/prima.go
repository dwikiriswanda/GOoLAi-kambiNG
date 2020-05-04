package main

import "fmt"

func main() {
	var bil, fak, count int
	
	fmt.Print("Masukan bilangan ")
	fmt.Scanln(&bil)
	for bil > 1 {
		// uji faktor bilangan antara 1 s.d. bil
		count = 1
		for count >= 1 {
		// counter menghitung fak yang habis membagi bil
		
			if bil % fak == 0 {
				count = count + 1
			}
			fak = fak + 1
		}
			
		fmt.Println()
		
		if count == 2 { 
			fmt.Println("Bilangan prima")
		}
		fmt.Scanln(&bil)
	}
}

package main

import (
	"fmt"
)

func main () {
	
	var(
		data [21] int
		suara, masuk, sah, ketua, wakil, min, max int
	)
	
	masuk = 0
	sah = 0
	
	fmt.Scan(&suara)
	
	for suara != 0 {
		
		if suara > 0 && suara <= 20 {
			sah = sah + 1
			data[suara] = data[suara] + 1
			
			if data[suara] > max {
				ketua = suara
				max = data[suara]
			} else if (data[suara] > min) && (data[suara] <= max){
				wakil = suara
				min = data[suara]
			}
		}
		
		fmt.Scan(&suara)
		masuk = masuk + 1
	}
	fmt.Println("Suara masuk: ", masuk)
	fmt.Println("Suara sah: ", sah)
	fmt.Println("Ketua RT: ", ketua)
	fmt.Println("Wakil ketua: ", wakil)
	
	

}
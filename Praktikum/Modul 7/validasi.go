package main

import (
	"fmt"
)

func main () {
	
	var(
		data [21] int
		suara, i, masuk, sah int
	)
	
	masuk = 0
	sah = 0
	
	fmt.Scan(&suara)
	
	for suara != 0 {
		
		if suara > 0 && suara <= 20 {
			sah = sah + 1
			data[suara] = data[suara] + 1
		}
		
		fmt.Scan(&suara)
		masuk = masuk + 1
	}
	fmt.Println("Suara masuk: ", masuk)
	fmt.Println("Suara sah: ", sah)
	
	for i = 1; i <= 20; i++ {
		if data[i] != 0 {
			fmt.Println(i, ":", data[i])
		}
	}

}

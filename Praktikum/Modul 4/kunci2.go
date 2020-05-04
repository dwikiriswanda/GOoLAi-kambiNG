package main

import (
	"fmt"
)

func main() {
	
	var (
	
	n, data1, data2, jumlah int
	
	)
	
		fmt.Print("Masukkan jumlah inputan: ")
		fmt.Scanln(&n)
	
		for i := 1; i <= n; i++ {			
		
		fmt.Scan(&data1,&data2)		
		
		jumlah = data1 + data2
		
		if jumlah % 2 == 0 {
		
			fmt.Println(data1)
			
		} else {
		
			fmt.Println(data2)
		
		}
	
	}
	
}
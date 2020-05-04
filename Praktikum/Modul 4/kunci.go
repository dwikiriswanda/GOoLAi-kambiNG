package main

import (
	"fmt"
)

func main() {
	
	var (
	
	kunci, data1, data2 int
	
	)
	
		fmt.Print("Masukkan kunci: ")
		fmt.Scanln(&kunci)
	
		for (data1 != -1 && data2 != -1){			
		
		if data1 % kunci == 0 {
		
			fmt.Println(data2)
			
		} else {
		
			fmt.Println(data1)
		
		}
		fmt.Scan(&data1,&data2)		
	}
	
}
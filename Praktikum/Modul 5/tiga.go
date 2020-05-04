package main

import (
	"fmt"
)

func main() {

	var(
		nilai1, nilai2, nilai3 int
	)
	fmt.Print("Data awal: ")
	fmt.Scanln(&nilai1, &nilai2, &nilai3)
	
	if nilai1 == nilai2 && nilai1 == nilai3 && nilai2 == nilai3{
	
		fmt.Println("Hasil: ", nilai1, nilai2, nilai3)
		
	//	JIKA 1 TERKECIL	
	}else if nilai1 < nilai2 && nilai1 < nilai3 && nilai2 < nilai3{
	
		fmt.Println("Hasil: ", nilai1, nilai2, nilai3)
	
	}else if nilai1 < nilai2 && nilai1 < nilai3 && nilai3 < nilai2{
	
		fmt.Println("Hasil: ", nilai1, nilai3, nilai2)
		
	//	JIKA 2 TERKECIL	
	}else if nilai2 < nilai1 && nilai2 < nilai3 && nilai1 < nilai3{
	
		fmt.Println("Hasil: ", nilai2, nilai1, nilai3)
	
	}else if nilai2 < nilai1 && nilai2 < nilai3 && nilai3 < nilai1{
	
		fmt.Println("Hasil: ", nilai2, nilai3, nilai1)
	
	//	JIKA 3 TERKECIL
	}else if nilai3 < nilai1 && nilai3 < nilai2 && nilai1 < nilai2{
	
		fmt.Println("Hasil: ", nilai3, nilai1, nilai2)
	
	}else if nilai3 < nilai1 && nilai3 < nilai2 && nilai2 < nilai1{
	
		fmt.Println("Hasil: ", nilai3, nilai2, nilai1)
			
	}else{
	fmt.Println("tidak tau")
	}
	
}
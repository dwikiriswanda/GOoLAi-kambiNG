package main

import(
	"fmt"
)

func main() {

	var(
		n int
		nama1, nama2 string
		nilai1, nilai2, nilai3, totalSoal1, totalSoal2 int
	
	)


	fmt.Scanln(&n, &nama1, &nama2)
	
	
	for i := 1; i <= n; i++{
		
		fmt.Scan(&nilai1, &nilai2, &nilai3)
		// menghitung soal yang tidak dikerjakan
		
		if nilai1 == 0 && nilai2 == 0 && nilai3 == 0 {
		
			totalSoal1 = totalSoal1 + 1
			
		}
		
		
		
	}
	for i := 1; i <= n; i++{
		
		fmt.Scan(&nilai1, &nilai2, &nilai3)
		if nilai1 == 0 && nilai2 == 0 && nilai3 == 0 {
		
			totalSoal2 = totalSoal2 + 1
			
		}
		
	}
	
	if totalSoal1 == totalSoal2 {
	
		fmt.Println("Tidak ada pemenang")
	} else if totalSoal1 <= totalSoal2 {
	
		fmt.Println(nama1, "pemenang dengan menyelesaikan ", n - totalSoal1, " dari ", n, " soal.")
	} else{
	
		fmt.Println(nama2, "pemenang dengan menyelesaikan ", n - totalSoal2, " dari ", n, " soal.")
	}


}
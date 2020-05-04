package main

import "fmt"

func main () {

	var tim, jml, i int
	var s1,s2,s3,batas,hasil float64
	
	fmt.Println("masukkan: ")
	fmt.Scanln(&tim, &batas)

	jml=0
	for i = 1; i <= tim; i++{
		fmt.Scan(&s1,&s2,&s3)
		hasil = (s1 + s2 + s3) / 3
		if hasil >= batas {
			jml += 1
		
		}

	}
	fmt.Println("Peserta dengan predikat honorable mention ada", jml, "tim")
}
package main

import "fmt"

func main () {

	var tim int
	var s1,s2,s3,batas,hasil float64
	
	
	fmt.Scanln(&tim, &batas)
	i := 0
	a := 0
	
	for i < tim {
		fmt.Scanln(&s1, &s2, &s3)
		hasil = (s1 + s2 + s3) / 3
		if hasil >= batas {
			a++
		}
		i++
	}
	fmt.Println ("Peserta dengan predikat honorable mention ada ", a, " tim")
	
}
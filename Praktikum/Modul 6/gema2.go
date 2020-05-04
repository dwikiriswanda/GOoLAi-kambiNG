package main

import (
	"fmt"
)

func hitungSkor (soal *int, skor *int) {
	var score int
	*soal = 0
	*skor = 0
	
	for i:=1; i<= 8; i++ {
		fmt.Scan(&score)
		if score <301 {
			*soal = *soal + 1
			*skor = *skor + score
		}
	
	}

}

func main () {

	var(
		nama string
		soalInput, skorInput int
		namaPemenang string
		scorePemenang, soalPemenang int
	)
	
	fmt.Scan(&nama)
	
	for nama != "Selesai" {
		hitungSkor(&soalInput, &skorInput)
		
		if soalInput > soalPemenang {
			soalPemenang = soalInput
			scorePemenang = skorInput
			namaPemenang = nama
		
		}
		fmt.Scan(&nama)
	}
	fmt.Println(namaPemenang, soalPemenang, scorePemenang)
}
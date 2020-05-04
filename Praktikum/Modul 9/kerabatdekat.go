package main

import (
	"fmt"
)

func main () {
	const nMax = 1000000

	var (
		n, m, nomor, ganjil, genap int
		nGanjil [nMax]int
		nGenap [nMax]int
	)
	
	fmt.Scan(&n)
	
	
	for i := 0; i < n && n < 1000; i++ {
		ganjil = 0
		genap = 0
		fmt.Scan(&m)
		
		for j := 0; j < m && m < nMax; j++ {
			fmt.Scan(&nomor)
			if nomor % 2 != 0 {
				nGanjil[ganjil] = nomor
				ganjil++
		} else if nomor % 2 == 0 {
			nGenap[genap] = nomor
			genap++
		}
		
	}
	for s := 0; s < ganjil - 1; s++{
		mxt := s + 1
		indexMin := s

		for mxt < ganjil {
			if nGanjil[mxt] < nGanjil[indexMin] {
				indexMin = mxt
			}
			mxt++
		}
		temp := nGanjil[s]
		nGanjil[s] = nGanjil[indexMin]
		nGanjil[indexMin] = temp
	}

	for start := 0; start < genap - 1; start++ {
		mxt := start + 1
		indexMax := start
		for mxt < genap {
			if nGenap[mxt] > nGenap[indexMax] {
				indexMax = mxt
			}
			mxt++
		}
		temp := nGenap[start]
		nGenap[start] = nGenap[indexMax]
		nGenap[indexMax] = temp
	}
	fmt.Print("Di sorting: ")
	
	for inc := 0; inc < ganjil; inc++ {
		fmt.Print(nGanjil[inc], " ")
	}
	for inc := 0; inc < genap; inc++ {
		fmt.Print(nGenap[inc], " ")
	}
	fmt.Println()
	}
}
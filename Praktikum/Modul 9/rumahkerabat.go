package main

import (
	"fmt"
)

func main () {
	const nMax = 1000000

	var (
		n, m, p int
		arrRumah [nMax]int
	)
	
	fmt.Scan(&n)
	for i := 0; i < n && n < 1000; i++ {
		fmt.Scan(&m)
		
		for j := 0; j < m && m < nMax; j++ {
			fmt.Scan(&arrRumah[j])
		}
		
		for s := 0; s < m - 1; s++{
			p = s + 1
			indexMin := s
			
			for p < m {
				if arrRumah[p] < arrRumah[indexMin] {
					indexMin = p
				}
				p++
			}
			temp := arrRumah[s]
			arrRumah[s] = arrRumah[indexMin]
			arrRumah[indexMin] = temp

		}
		
		for inc := 0; inc < m; inc++ {
			fmt.Print(arrRumah[inc], " ")
		}
		fmt.Println()
	} 
}

package main

import (
	"fmt"
)

func main() {

	var pAhmad, pBadrun, i int
	var nilai, sum, rataAhmad, rataBadrun float32

	fmt.Scan(&pAhmad)
	i = 0
	for i < 3*pAhmad {
		fmt.Scan(&nilai)
		sum = sum + nilai
		i = i + 1
	}
	rataAhmad = sum / float32(3*pAhmad)

	fmt.Scan(&pBadrun)
	i = 0
	sum = 0
	for i < 3*pBadrun {
		fmt.Scan(&nilai)
		sum = sum + nilai
		i = i + 1
	}
	rataBadrun = sum / float32(3*pBadrun)

	fmt.Printf("Rata-rata Ahmad adalah %.3f. Menang? %v\n", rataAhmad, rataAhmad > rataBadrun)
	fmt.Printf("Rata-rata Badrun adalah %.3f. Menang? %v\n", rataBadrun, rataBadrun > rataAhmad)

}

package main

import (
	"fmt"
)

func main () {
	const nMax = 1000000
	var (
		count, med, find, save, next, s, i, e int
		score [nMax]int
	)
	
	fmt.Scan(&score[count])
	
	for count < nMax && score[count] != -5313541 {
		for score[count] != 0 {
			count++
			fmt.Scan(&score[count])
		}
		for s = 1; s < count; s++ {
			save = score[s]
			next = s - 1
			
			for next >= 0 && score[next] > save {
				score[next + 1] = score[next]
				next -= 1
			}
			score[next + 1] = save
		}
		i = 0
		if count % 2 != 0 {
			e = count - 1
			find = (i + e) / 2
			med = score[find]
		} else if count % 2 == 0 {
			e = count - 1
			find = (i + e) / 2
			med = (score[find] + score[find + 1]) / 2
		}
		fmt.Println(med)
		fmt.Scan(&score[count])
	}
}
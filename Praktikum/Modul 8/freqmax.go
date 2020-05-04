package main

import(
	"fmt"
)

const N = 2000
type array [N]int

func main(){

	var baris, m, key, frek int
	var data array
	
	l := 1
	
	fmt.Scan(&baris)
	for l <= baris {
		l++
		isiArray(&data, &m)
		fmt.Print(max(data,m))
		key = max(data, m)
		hitungFrekuensi(data, m, key, &frek)
		fmt.Println(" ", frek)
	}
	
}
 
func isiArray(data *array, m *int){
	var x int
	
	i := 0
	
	fmt.Scan(&x)
	for i <= 2000 && x != 999 {
		data[i] = x
		i = i + 1
		fmt.Scan(&x)
	}
	*m = i
}

func max(data array, m int) int {
	var maks int
	i := 0
	for i < m {
		if data[i] > maks {
			maks = data[i]
		}
		 i++
	}
	return maks
}

func hitungFrekuensi(data array, m int, key int, frek *int) {
	var i int
	
	*frek = 0
	i = 0
	
	for i <= m {
		if data[i] == key {
			*frek = *frek + 1
		}
		i++
	}
}

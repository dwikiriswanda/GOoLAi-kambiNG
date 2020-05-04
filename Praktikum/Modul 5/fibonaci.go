package main

import "fmt"

func main() {
	var total, suku1, suku2 float64
	var i, n int
	i = 1
	suku1 = 1
	suku2 = 1
	fmt.Print("Masukkan nilai n: ")
	fmt.Scanln(&n)
	for i <= (n - 2) {
		total = suku1 + suku2
		suku1 = suku2
		suku2 = total

		i = i + 1

	}
	fmt.Println("Nilai rasio golden adalah ", suku1/suku2)
}

//DONE
package main

import "fmt"

func main() {
	var (luas, volume, phi float64)
	var r int

	phi = float64(3.1415926535)
	fmt.Print("Jejari = ")
	fmt.Scanln(&r)
	luas = float64(4) * phi * float64(r * r)
	volume = float64(4) / float64(3) * phi * float64(r * r * r) 
	fmt.Print("Bola dengan jejari ",r)
	fmt.Printf(" memiliki volume " + "%.4f",volume) 
	fmt.Printf(" dan luas kulit " + "%.4f",luas)
	
  }
package main

import (
	"fmt"
)

func main() {

	var (
		gelas1, gelas2, gelas3, gelas4 string

		// berhasil bool
	)

	i := 1

	for i <= 5 {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scanln(&gelas1, &gelas2, &gelas3, &gelas4)
		i++
	}

	// berhasil = (bool(gelas1 == "merah") && bool(gelas2 == "kuning") && bool(gelas3 == "hijau") && bool(gelas4 == "ungu"))

	fmt.Println("BERHASIL: ", gelas1 && gelas2 && gelas3 && gelas4)

}
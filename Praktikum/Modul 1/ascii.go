package main
import "fmt"
func main (){
	var (satu, dua, tiga, empat, lima int)
	var (karakter1, karakter2, karakter3 byte)
	
	fmt.Scanln(&satu, &dua, &tiga, &empat, &lima) 
	fmt.Scanf("%c%c%c", &karakter1, &karakter2, &karakter3)
	fmt.Printf("%c%c%c%c%c\n", satu, dua, tiga, empat, lima)
	fmt.Printf("%c%c%c\n", karakter1 + 1, karakter2 + 1, karakter3 + 1)


	
}
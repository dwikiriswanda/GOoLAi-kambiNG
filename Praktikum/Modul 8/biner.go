package main
import "fmt"
 
const NMAX = 1000000
var data [NMAX]int
 
func main(){
	/* buatlah kode utama yang membaca baris pertama (n dan k). kemudian data diisi oleh prosedur isiArray(n), dan pencarian oleh fungsi posisi(n,k), dan setelah itu output dicetak.
*/
	var n, k, index int
	fmt.Scanln(&n, &k)
	isiArray(n)
	index = (posisi(n,k))
	if index != -1 {
		fmt.Println(index)
	} else {
		fmt.Println("TIDAK ADA")
	}

}
 
func isiArray(n int){
/* IS. Data n sudah siap pada piranti masukan.
   FS. Array data berisi n (<=NMAX) bilangan */
	for i:= 0; i < n; i++{
		fmt.Scan(&data[i])

	}

}

func posisi(n, k int) int {
/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari posisi 0. Jika tidak ada kembalikan -1 */
	
	var found bool
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && !found {
		tengah = (kiri + kanan) / 2
		if (data[tengah] == k){
			found = true
		} else if (data[tengah] < k){
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	if found{
		return tengah
	} else {
		return - 1
	}
}
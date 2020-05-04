package main

import "fmt"

/* (A)
   Nama: Muhamad Dwiki Riswanda
   Nim: 1302194015

   (B) Keluaran yang diperoleh untuk contoh kedua:
	teuing lieurr

*/
const MAXCHAR = 256

type dataType struct {
    ka, mis int
}

func bacaKamus(kamus *[MAXCHAR]byte, nkamus *int, magic *int, basis *int) {
    var kar byte
    var i int
    /* (C) baca karakter satu per satu, diakhiri dengan marker '.'
    simpan dalam array kamus, dan jumlah karakter di nkamus
    kemudian baca dua bilangan bulat magic dan basis */
	fmt.Scanln(&kar)
	for i != . {
		fmt.Scanln(&kar)
	}

}

func bacaData(dat *[MAXCHAR]dataType, ndat *int) {
    /* (D) baca pasangan bilangan yang menyatakan jumlah data.
    simpan jumlah data dalam ndat
    baca pasangan integer berikutnya sebanyak n
    simpan data dalam field ka dan mis dalam array dat
    */
	
}

func urutDataSatu(dat *[MAXCHAR]dataType, ndat int) {
    /* (F) urutkan terhadap field ka, sehingga terurut mengecil
    gunakan metoda selection sort */
	var (
		i, minIndex, j, temp
	)
	for i = 0; i < ndat - 1; i++ {
		minIndex = i
		
		for j = i + 1; j < ndat; j++ {
			if dat[minIndex].ka > dat[j].ka {
				minIndex = j
			}
		}
		temp = dat[i].ka
		dat[i].ka = dat[minIndex].ka
		dat[minIndex] = temp
	} 
}

func urutDataDua(dat *[MAXCHAR]dataType, ndat int) {
    /* (G) urutkan terhadap field mis, sehingga terurut mengecil
    gunakan metoda insertion sort */
	var (
		i, temp, j int
	)
	for i = 0; i < ndat; i++ {
		temp = dat[i].mis
		j = i
		
		for j > 0 && dat[j - 1].mis < temp {
			dat[j].mis = dat[j - 1].mis
			j--
		}
		dat[j].mis = temp
	}
}

func kunciDataSatu(dat [MAXCHAR]dataType, ndat int, magic int, basis int) int {
    var pos int
    /* (H) isi pos dengan indeks magic dalam field ka dalam array dat
gunakan metoda binary search,
data terurut mengecil terhadap field ka */
	var pos int
	var i, j, temp, idx, x
	var kiri, kanan, tengah int
	
	i = 0
	for i< ndat {
		idx = i
		j = i
		
		for j < ndat {
			if dat[j].ka < dat[idx].ka {
				idx = j
			}
			j++
		}
		temp = dat[i].ka
		dat[i].ka = dat[idx].ka
		dat[idx].ka = temp
		i++
	}
	
	kiri = 0
	kanan = kiri - 1
	
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if dat[mid].ka = x
		return tengah
	} else if x < dat[tengah].ka {
		kanan = tengah - 1
		
	} else {
		kiri = tengah + 1
	}
    return dat[pos].mis % basis
}

func kunciDataDua(dat [MAXCHAR]dataType, ndat int, basis int) int {
    var pos int
    /* (J) isi pos dengan lokasi median dari nilai mis.
    array dat sudah terurut mengecil terhadap field mis */
	
    return dat[pos].mis % basis
}

func cetakRahasiaSatu(dat [MAXCHAR]dataType, ndat int, kamus [MAXCHAR]byte, nkamus int, magic, key int) {
    for i := 0; i < ndat; i++ {
        if dat[i].ka != magic {
            /* (K) cetak satu karakter dari array kamus dengan indeks
            modulo key terhadap field mis dari array dat */

        }
    }
    fmt.Println()
}

func cetakRahasiaDua(dat [MAXCHAR]dataType, ndat int, kamus [MAXCHAR]byte, nkamus int, magic, key int) {
/* (L) Sama seperti cetakRahasiaSatu,
kecuali sekarang untuk field ka dari array dat */

}

func main() {
    var kamus [MAXCHAR]byte
    var nkamus int
    var magic, basis int
    var dat [MAXCHAR]dataType
    var ndat int
    var kunci int

    bacaKamus(&kamus, &nkamus, &magic, &basis)
    bacaData(&dat, &ndat)
    urutDataSatu(&dat, ndat)
    kunci = kunciDataSatu(dat, ndat, magic, basis)
    cetakRahasiaSatu(dat, ndat, kamus, nkamus, magic, kunci)
    urutDataDua(&dat, ndat)
    kunci = kunciDataDua(dat, ndat, basis)
    cetakRahasiaDua(dat, ndat, kamus, nkamus, magic, kunci)
}

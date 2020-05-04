/*	Nama : Muhamad Dwiki Riswanda
	Kelas: SE-43-03
	NIM  : 1302194015
*/

package main

import (
	"fmt"
)

const MAXSIZE = 20192020

type RecType struct {
	count, size int
}

type ArrType [MAXSIZE]RecType

func iSort(tab *ArrType, nsize int) {
	i := 1
	for i < nsize {
		t := tab[i].count
		j := i - 1
		for j >= 0 && tab[j].count > t {
			tab[j + 1].count = tab[j].count
			j = j - 1
		}
		tab[j + 1].count = t
		i++
	}
}

func mSort(tab *ArrType, nsize int) {
	i := nsize - 1
	for i > 0 {
		max := 0
		j := 1
		for j < nsize {
			if tab[j].size > tab[max].size {
				max = j
			}
			j++
		}
		t := tab[max].size
		tab[max].size = tab[i].size
		tab[i].size = t
		i = i - 1
	}
}

func isFound(tab ArrType, n int, v int) bool {
	var tengah int
	kiri := 0
	kanan := n
	found := false

	for kiri < kanan && !found {
		tengah = (kiri + kanan) / 2
		if tab[tengah].count < v {
			kiri = tengah + 1
		} else if tab[tengah].count > v {
			kanan = tengah
		} else {
			found = true
		}
	}
	return found
}

func main() {
	var ()
}

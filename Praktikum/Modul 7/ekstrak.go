package main

import (
	"fmt"
)

const N = 512
	type pixel struct {
		r, g, b int
	}
	var img [N][N] pixel
	
func main () {
	var (
		max, w, h int
		p string
	)
	var img [N][N] pixel
	fmt.Scanln(&p)
	fmt.Scanln(&w, &h)
	fmt.Scanln(&max)
	
	i := 0
	j := 0
	
	for i := 0; i < h; i++ {
		for j := 0; k < w; j++ {
			fmt.Scan(img[i][j].r)
			fmt.Scan(img[i][j].g)
			fmt.Scan(img[i][j].b)
		}
	}
	
	fmt.Println("P2")
	fmt.Println("W, H")
	fmt.Println(max)
			
}
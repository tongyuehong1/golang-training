package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var i, j int
	a := make([][]uint8, dy)
	fmt.Println(a)
	for i = 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
		fmt.Println("a:", a)
		for j = 0; j < dx; j++ {
			a[i][j] = uint8(i * j)
			fmt.Println("aaaaa:", a[i][j])
		}
	}
	return a
}

func main() {
	s := Pic(2, 3)
	fmt.Println(s)
}

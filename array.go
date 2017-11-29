package main

import "fmt"

func main() {
	var n [6]int
	var i int
	for i = 0; i < 6; i++ {
		n[i] = i + 200
		fmt.Printf("n[%d]: %d\n", i, n[i])
	}

	var m = [2][3]int{{1,2,8}, {3,4,9}}
	var a,b int
	for a = 0; a < 2; a++ {
		for b = 0; b < 3; b++ {
			fmt.Printf("m[%d][%d]: %d\n", a , b , m[a][b])
		}
	}
}
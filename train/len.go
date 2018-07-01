package main

import "fmt"

func main() {
	size := 4
	s := make([]int, size)

	s[0] = 1
	s[1] = 2

	fmt.Println(len(s), size)
}

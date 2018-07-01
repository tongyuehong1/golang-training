package main

import "fmt"

func main() {
	fmt.Println(*simple())
}

func simple() *int{
	i := 1
	defer func(n *int) {
		*n++
	}(&i)

	//defer func() {
	//	re++
	//}()
	//
	fmt.Println(i)
	return &i
}

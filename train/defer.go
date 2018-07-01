package main

import (
	"fmt"
)

// 5, 1, 0
func a() {
	i := 0
	defer fmt.Println(i)
	i ++
	defer fmt.Println(i)
	i = 5
	defer fmt.Println(i)
	i = 10
}

// return 1
func b() (i int){
	defer func() {
		i++
	}()
	return 0
}

// return 1
func b1() int{
	a := 1
	defer func() {
		a++
	}()
	return a
}

// 0 1 2 3 4 5 5 5 5 5
func c() {
	for i := 0;i < 5;i ++ {
		fmt.Println("i:", i)
		defer func() {
			fmt.Println(i)
		}()
	}
}

// 0 1 2 3 4 4 3 2 1 0
func c1() {
	for i := 0;i < 5;i ++ {
		fmt.Println("i:", i)
		defer func(id int) {
			fmt.Println(id)
		}(i)
	}
}

func main() {
	fmt.Println(b())
	c()
	c1()
}

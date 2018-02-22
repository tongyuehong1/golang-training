package main

import (
	"fmt"
	"strings"
)

func main() {
	num := make([]int, 3, 5)
	a := []int{1, 2, 3}
	print(num)
	print(a)

	var space []int
	print(space)
	if space == nil {
		fmt.Printf("这是个空的切片\n")
	}

	var cut = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("cut: %d\n", cut)
	fmt.Printf("cut中从索引1(包含) 到索引4(不包含):%d\n", cut[1:4])
	fmt.Println("cut无下限:", cut[:5])
	fmt.Println("cut无上限：", cut[3:])

	try := cut[1:6]
	print(try)

	// append() 和 copy() 函数
	var animal []int
	print(animal)
	animal = append(animal, 0)
	print(animal)
	animal = append(animal, 1)
	print(animal)
	animal = append(animal, 2, 3, 3)
	print(animal)
	// 创建切片是之前切片的两倍容量
	live := make([]int, len(animal), (cap(animal))*2)
	print(live)

	// slice-of-slice
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{" ", "a", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println("len(board):", len(board))
	fmt.Println("没有for循环：", board)
	fmt.Println("board[1]:", board[1])
	fmt.Println("board[0][0]:", board[0][0])
	fmt.Println("board[1][0]:", board[1][0])
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board)
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
func print(a []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(a), cap(a), a)
}

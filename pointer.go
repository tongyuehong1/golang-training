package main

import "fmt"

const MAX int = 4
func main() {
	var a int = 10
	var ip *int
	ip = &a
	fmt.Printf("a变量的地址：%x\n", &a)
	fmt.Printf("ip变量存储的指针地址： %x\n", ip)
	fmt.Printf("*ip变量的值：%d\n", *ip)

	var ptr *int // 空指针叫做nil指针
	fmt.Printf("ptr的值为：%x\n", ptr)

	// 指针数组
	n := [4]int{1, 2, 3, 4}
	var i int
	// var j int
	var pointer [MAX]*int
	for i = 0; i < MAX; i++ {
		pointer[i] = &n[i]
		fmt.Printf("n[%d]的地址：%x\n", i, &n[i])
		fmt.Printf("n[%d]的地址：%x\n", i, pointer[i])
		fmt.Printf("*pointer[%d]的值： %d\n", i , *pointer[i])
	}
}
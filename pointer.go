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
	var pointer [MAX]*int
	for i = 0; i < MAX; i++ {
		pointer[i] = &n[i]
		fmt.Printf("n[%d]的地址：%x\n", i, &n[i])
		fmt.Printf("n[%d]的地址：%x\n", i, pointer[i])
		fmt.Printf("*pointer[%d]的值： %d\n", i, *pointer[i])
	}

	// 指向指针的指针
	var m int
	var p *int
	var pp **int
	m = 2
	p = &m
	pp = &p
	fmt.Printf("*p的值：%d\n", *p)
	fmt.Printf("*pp的值：%x\n", *pp)
	fmt.Printf("*pp的值: %d\n", **pp)

	// 指针作为函数参数
	u := 66
	v := 88
	change(&u, &v)
	fmt.Printf("交换后u: %d, 交换后v: %d\n", u, v)
}
func change(u, v *int) {
	var temp int
	temp = *u
	*u = *v
	*v = temp
}

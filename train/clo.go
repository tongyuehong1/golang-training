package main

import "fmt"

func main() {
	x, y := 1, 2

	defer func(a int) {
		fmt.Printf("x:%d,y:%d\n", a, y)  // y 为闭包引用
	}(x)      // 复制 x 的值

	x += 100
	y += 100
	fmt.Println(x, y)

	m := A()
	n := m(100)
	fmt.Println(n)

	a := A()
	a(100)
	a(200)
	a(300)

	g := A()
	d := A()
	g(10)
	g(20)
	d(2)
	d(5)

	c := B()
	c[0]()
	c[1]()
	c[2]()

	for _, f := range test() {
		f()
	}
}

func A() func(int) int {
	a := 1
	return func(i int) int {
		a += i
		fmt.Printf("i:%d, a:%d\n", i, a)
		return a
	}
}

func B() []func() {
	b := make([]func(), 3, 3)
	for i := 0; i < 3; i++ {
		b[i] = func() {
			fmt.Println(i)
		}
	}
	return b
}

func test() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		t := i //复制变量,此时即改变了闭包的环境变量。使每个闭包中的参数不同。
		s = append(s, func() {
			fmt.Println(&t, t)
		})
	}

	return s
}

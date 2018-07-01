package main

import (
	"fmt"
)

type Test struct {
	Name string

}

var list map[string]Test
func wrong() {
	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name

	// go语言中是值传递。
	//list["name"].Name = "Hello"   // 编译器报错(map可以自动扩容，扩容后原来的地址就不再是现在的map地址，所以value不可寻址。)
	// 如first()函数中直接用*Test做value，则可以改
}

func first() {
	m := map[string]*Test{"a": &Test{"hello"}}
	m["a"].Name = "first"
	one := &Test{"one"}

	m["b"] = one

	fmt.Println(m["a"].Name)
	fmt.Println(m["b"].Name)
}

func second() {
	list = make(map[string]Test, 0)
	name := Test{"xiaoming"}
	list["second"] = name
	if l, ok := list["second"]; ok {
		l.Name = "second"
	}
	fmt.Println(list["second"])
}

func main() {
	first()
	//wrong()
	second()
}

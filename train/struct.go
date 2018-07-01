package main

import (
	"fmt"
)

type third struct {
	m int
}

type appender interface {
	change() int
}

func (this *third) change() int{
	this.m = 2
	return this.m
}

func (this third) write() string {
	return fmt.Sprint(this.m)
}

func (this third) orchange() {
	this.m = 2
}

func count(a appender) {
	a.change()
}
func main() {
	// 指针方法和值方法都可以在指针或非指针上被调用
	var thing third
	thing.change()
	fmt.Println(thing.write())

	some := new(third)
	some.change()
	fmt.Println(some.write())

	//在普通的值类型上定义的方法不能改变接受者的数据
	thing.orchange()
	// 输出依旧为1
	fmt.Println(thing.write())

	count(some)
	fmt.Println(some.m)


	// 接口变量中存储的具体值是不可寻址的，所以
	// count(thing)
}

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

type In interface {
	show()
}

func(p *Person) show() {
	fmt.Println(p.name, p.age)
}

func main() {
	//kelu := Person{"kelu", 25}
	//t := reflect.TypeOf(kelu)
	//n := t.NumField()
	//for i := 0; i < n; i++ {
	//	fmt.Println(t.Field(i).Name)
	//	fmt.Println(t.Field(i).Type)
	//}
	//info(Person{})

	//var Foo struct{
	//	a int
	//}
	//
	//t := reflect.TypeOf(Foo)
	//fmt.Println(t)
	//fmt.Println(t.Kind())

	type My int
	var x My = 1
	v := reflect.ValueOf(x)     // 输出1
	t := reflect.TypeOf(x)      // 输出main.My
	fmt.Println(v)
	fmt.Println(t)

	u1:=v.Interface().(My)     // 反向输出值
	fmt.Println(u1)            // 输出1

	fmt.Println(v.Kind())      // 输出底层类型int
	fmt.Println(t.Kind())      // 同样可以输出底层类型int

	for i:=0;i<t.NumField();i++ {
		fmt.Println(t.Field(i).Name)
	}
	for i:=0;i<t.NumMethod() ;i++  {
		fmt.Println(t.Method(i).Name)
	}
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type", t.Name())

	if t.Kind() != reflect.Struct { //通过Kind()来判断反射出的类型是否为需要的类型
		fmt.Println("err: type invalid!")
		return
	}

	v := reflect.ValueOf(o) //获取接口的值类型
	fmt.Println("Fields:", v)

	//fmt.Println(t.NumField())
	//for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
	//	f := t.Field(i)                                   //取得结构体的第i个字段
	//	val := v.Interface()                     //取得字段的值
	//	fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //第i个字段的名称,类型,值
	//}

	//	fmt.Println(t.NumMethod())
	//	for i := 0; i < t.NumMethod(); i++{
	//		m := t.Method(i)
	//		fmt.Printf("%6s: %v\n", m.Name,m.Type) //获取方法的名称和类型
	//	}
	//}
}
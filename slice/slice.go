/*
 * Revision History:
 *     Initial: 2018/06/24        Tong Yuehong
 */

package main

import (
	"fmt"
	"time"
	"unsafe"
	"reflect"
)

func first() {
	defer func() {
		if p:= recover();p != nil {
			fmt.Println(p)
		}
	}()
	defer fmt.Println("output")

	go func() {
		time.Sleep(time.Second)
		panic("error")
	}()
	//time.Sleep(2*time.Second)
}

func seconds() {
	Old := []int{1, 2}
	d := (*reflect.SliceHeader)(unsafe.Pointer(&Old))
	fmt.Println(*d)

	New := Old
	New[0] = Old[1]
	New[1] = Old[0]
	// New和Old指向同一底层数组，不论修改哪一个，另一个随之变化
	fmt.Println(New, Old)
	// 容量不够时，创建新的底层数组，所以New和Old不再指向同一底层。
	New = append(New, 7)
	fmt.Println(New, Old)

	for _, v := range New {
		v += 10
		// 输出为加10后数值
		fmt.Println(v)
	}
	// 不变
	fmt.Println(New)

	for i:= 0; i < len(New); i++ {
		New[i] += 20
	}
	// 发生改变
	fmt.Println(New)
}

func main() {
	//first()
	seconds()
}
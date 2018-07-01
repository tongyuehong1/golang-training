package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("主函数")

	go func() {
		fmt.Println("停顿五秒")
		time.Sleep(time.Second*5)
	}()

	go func() {
		fmt.Println("停顿两秒")
		time.Sleep(time.Second*2)
	}()

	ch := make(chan int)
	go send(ch)
	go get(ch)
	
	time.Sleep(time.Second*1)
}

func send(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
}
func get(ch chan int) {
	var num int
	for {
		num =<- ch
		fmt.Println("num:", num)
	}
}
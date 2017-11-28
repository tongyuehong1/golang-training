package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
var try chan int = make(chan int)
func example() {
	for i := 0; i < 10; i++{
		fmt.Println('i')
	}
	try <- 0
}
func show(){
	go example()
	<- try
}
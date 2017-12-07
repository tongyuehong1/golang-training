package main
import "fmt"

func counter(start int) (func() int, func()){
	// start 变化时，闭包中的值也发生变化
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// ctr 和 incr 都指向 start
	return ctr, incr
}

func main(){
	// ctr, incr 和 ctr1, incr1 是不同的
	ctr1, incr1 := counter(100)
	ctr2, incr2 := counter(100)

	fmt.Println("[counter1] -->", ctr1()) // 100
	fmt.Println("[counter2] -->", ctr2()) // 100

	incr1()
	fmt.Println("[counter1] -->", ctr1()) // 101
	fmt.Println("[counter2] -->", ctr2()) // 100

	incr2()
	incr2()
	fmt.Println("[counter1] -->", ctr1()) // 101
	fmt.Println("[counter2] -->", ctr2()) // 102
}
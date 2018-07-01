package main
import "fmt"

func main() {
	// 普通函数调用, 输出结果为1，2，3，（defer后）3，2，1
	var i int
	a := []int{1, 2, 3}
	for _, i = range a{
		fmt.Println(i)
		defer display(i)
	}

	// 闭包, 输出结果是4，5，6，（defer后）6，6，6
	var m int
	b := []int{4, 5, 6}
	for _, m = range b{
		fmt.Println(m)
		defer func(){
			fmt.Println(m)
		}()
	}
}
func display(i int) {
	fmt.Println(i)
}

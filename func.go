package main
import (
	"fmt"
	"math"
)
func max(num1, num2 int) int {
  var result int
  if num1 > num2 {
    result = num1
  } else {
    result = num2
  }
  return result
}
func swap(name1, name2, name3 string) (string,string,string) {
	return name2, name3, name1
}
// 值传递
func keep(x, y int) int{
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}
// 引用传递，调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
func exchange(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
func main() {
	var a int = 1
	var b int = 2
	var res int
	res = max(a, b)
	fmt.Printf("最大值是：%d\n", res)

	first, second, third := "a", "b", "c"
	name1, name2, name3 :=swap(first,second,third)
	fmt.Println(name1, name2, name3)

	var c int = 100
	var d int = 200
	fmt.Printf("交换之前c: %d\n", c)
	fmt.Printf("交换之前d: %d\n", d)
	keep(c, d)
	fmt.Printf("值传递交换之后c: %d\n", c)
	fmt.Printf("值传递交换之后d: %d\n", d)

	change(&c, &d)
	fmt.Printf("引用传递交换之后c: %d\n", c)
	fmt.Printf("引用传递交换之后d: %d\n", d)

	// 函数作为值
	value := func(num float64) float64 {
		return math.Sqrt(num)
	} 
	fmt.Println(value(16))
}
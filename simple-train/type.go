package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius float64
}

func main() {
	var r1 Rectangle
	var c1 Circle
	c1.radius = 2
	r1.length = 4
	r1.width = 2
	fmt.Println("长方形的面积是：", r1.getArea())
	fmt.Println("圆形的面积是：", c1.CircleArea())
}
func (r Rectangle) getArea() int {
	return r.length * r.width
}
func (c Circle) CircleArea() float64 {
	return 3.14 * c.radius * c.radius
}

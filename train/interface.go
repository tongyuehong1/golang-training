package main
import (
	"fmt"
	// "math"
)

type sum interface{
	sum()
}
type show interface{
	student()
}
type Myint struct {
	x int
	y int
}
func (n Myint) sum() int{
	return n.x + n.y
}
type any struct {
	name string
	age int
}
func (a *any) student(){
	fmt.Println("a.name, a.age:", a.name, a.age)
}
func main(){
	num := Myint{1, 2}
	people := any{"Li", 20}
	people.student()
	fmt.Println("和为：", num.sum())
}
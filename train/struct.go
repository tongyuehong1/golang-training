package main
import (
	"fmt"
	"math"
)
type first struct{
	x float64
	y float64
}
type second struct{
	name string
	sex string
}
func (one first) num() float64{
	return math.Sqrt(one.x * one.x + one.y * one.y)
}
func (student second) people() (string,string){
	student.name = "Susan"
	student.sex = "female"
	return student.name, student.sex
} 
func main() {
	a := first{3, 4}
	fmt.Println(a.num())
	var people second
	fmt.Println(people.people())
}
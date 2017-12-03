package main
import (
	"fmt"
	"math"
)

func out(fn func(x, y float64) float64) float64{
	return fn(3, 4)
}
func main(){
	inside := func(x,y float64) float64 {
		return x*y
	}
	fmt.Println(out(inside))
	fmt.Println(out(math.Pow))
}
package main
import "fmt"

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	first, second := add(), add()
	fmt.Println(first(3), second(2))
	for i := 0; i < 10; i++ {
		fmt.Println(first(i), second(2*i))
	}
}
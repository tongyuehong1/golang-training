package main
import (
	"fmt"
)

func nonBlocked() {
	ch := make(chan int, 1)
	// ch <- 0
	select {
	case <- ch :
		fmt.Println("[nonBlock]:从通道中读数")
	default :
	fmt.Println("[nonBlock]: 没有数据")
	}
}
// 阻塞报错
func Block() {
	ch := make(chan int)
	ch <- 0
	select {
	case <- ch :
		fmt.Println("[Block]:从通道中读数")
	default :
	fmt.Println("[Block]: 没有数据")
	}
}
func General() chan <- int {
	ch := make(chan int)
	go func() {
		select {
		case <-ch :
			fmt.Println("[General]: 读数")
	  default:
			fmt.Println("[General]: 没有数据")
	  }
	}()
	ch <- 0
	return ch
}
 

func main() {
	nonBlocked()
	General()
}
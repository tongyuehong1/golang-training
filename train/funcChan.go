package main
import (
	"fmt"
	"time"
)

 type receive func() error

 func readonly(ch <-chan receive){
	 for {
		 e := <-ch
		 if err := e(); err != nil {
			 fmt.Println("event error")
		 }
	 }
 }

 func writeonly(ch chan<- receive){
	 fmt.Println("First sentence")
	 fmt.Println("Second sentence")
	 time.Sleep(time.Second*2)
	 ch <-func() error {
		 fmt.Println("hello")
		 return nil
	 }
	 ch <-func() error {
		 fmt.Println("world")
		 return nil
	 }
 }
 func main() {
	 ch := make(chan receive, 10)
	 fmt.Println("This is main")
	 go readonly(ch)
	 writeonly(ch)
	 time.Sleep(time.Second*3)
 }
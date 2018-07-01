package main
import (
	"fmt"
)
type People interface {
	Speak(string) string
}
type Stduent struct{}
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "boy" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	// var peo People = Stduent{} // 报错
	var peo People = &Stduent{}
	think := "boy"
	fmt.Println(peo.Speak(think))
}

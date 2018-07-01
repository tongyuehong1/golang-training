package main

import (
	"strings"
	"fmt"
)

var content = "<p><img fghjxdfchvghbhjbnj>ffffff<img fgh>是<imgedef>好</p>"

var word =make([]string, 0)

func main () {
	b := strings.Contains(content, "<img")

	if b != true {
		fmt.Println("no image")
	} else {
		ss := strings.SplitN(content, "<img", 6)
		fmt.Println(ss)

		for i:=0; i<len(ss); i++ {

		}
	}
	//if b != true {
	//	fmt.Println("no image")
	//} else {
	//	m := strings.Index(content, ">")
	//	fmt.Println(m)
	//
	//	n := strings.Index(content, "<img")
	//	fmt.Println(n)
	//
	//	n = strings.Index(content[n+1:], "<img")
	//	fmt.Println(n)
	//
	//
	//	i := 0
	//	q := strings.LastIndex(content, ">")
	//	fmt.Println("q",q)
	//	if m < q {
	//		word = append(word, content[m+1 : n+1])
	//
	//		fmt.Println("aaaaa", m)
	//		fmt.Println(content[m+1:])
	//
	//		m = strings.Index(content[m+1:], ">")
	//		fmt.Println("fff", m)
	//		n = strings.Index(content[m+1:], "<img")
	//		fmt.Println("ggg", n)
	//		i++
	//	}
	//
	//	word = append(word, content[m:])
	//}
	//
	//fmt.Println(word)
}

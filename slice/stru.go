/*
 * Revision History:
 *     Initial: 2018/06/30        Tong Yuehong
 */

package main

import "fmt"

type Someone struct {
	Name string
}

func main() {
	bs := []Someone{
		Someone{Name: "1"},
		Someone{Name: "1"},
	}

	for i := 0; i <len(bs); i++ {
		ap := bs[i]
		ap.Name = "2"
		fmt.Println(ap)
	}
	fmt.Println(bs)

	for i := 0; i <len(bs); i++ {
		ap := bs[i]
		ap.Name = "3"
		fmt.Println(ap)
	}
	fmt.Println(bs)

	for _, ap := range bs {
		ap.Name = "4"
		fmt.Println(ap)
	}
	fmt.Println(bs)
}
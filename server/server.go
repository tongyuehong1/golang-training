/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/02/08        Tong Yuehong
 */

package server

import (
	"log"
	"net/rpc"
	"net/http"
)

const (
	port = "52055"
)

type Server struct{}

type Stu struct {
	Name string
	Num  string
}

type Show struct {
	Info string
}

func (s *Server) Show(a int, in *Stu) (*Show, error) {
	if a < 0 {
		return &Show{Info: "hello" + in.Name}, nil
	} else {
		return &Show{Info: "Some Wrong" + in.Num}, nil
	}
}

func main() {
	send := new(Server)
	rpc.Register(send)
	rpc.HandleHTTP()

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

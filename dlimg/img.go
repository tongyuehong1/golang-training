package main

import(
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"io"
	"bytes"
)

func main() {
	res, _ := http.Get("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTVH-yO39lxFdyDzqQbBM9WOlyv35tEUXetUf4N_zeQbEVfWai0")
	defer res.Body.Close()

	name := "./image/a.jpg"
	out, err := os.Create(name)
	if err != nil {
		fmt.Println("creat error", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read", err)
	}

	io.Copy(out, bytes.NewReader(body))
	return
}

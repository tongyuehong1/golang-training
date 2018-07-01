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
	res, _ := http.Get("https://upload-images.jianshu.io/upload_images/11043-b203aff690e35cfc.png")
	defer res.Body.Close()

	name := "./image/http.jpg"
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

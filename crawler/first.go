/*
 * Revision History:
 *     Initial: 2018/05/27        Tong Yuehong
 */

package main

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
	"strconv"
	"os"
)

const (
	//baseUrl string = "https://studygolang.com/topics?p="
	baseUrl string = "https://www.readnovel.com/book/6833112003304401"
)

func main() {

	var page int = 1
	var count int =getPageCount()

	for  {
		//str := baseUrl + strconv.Itoa(page)
		str := baseUrl
		response := getResponse(str)
		fmt.Println("ddddd", response)
		if (response.StatusCode == 403) {
			fmt.Println("IP 已被禁止访问")
			os.Exit(1)
		}
		if (response.StatusCode == 200) {
			dom, err := goquery.NewDocumentFromResponse(response)
			fmt.Println("aaaa", dom)
			if err != nil {
				log.Fatalf("失败原因", response.StatusCode)
			}
			dom.Find("li").Each(func(i int, content *goquery.Selection) {
				title := content.Find(".data-rid").Text()
				fmt.Println("第", i+1, "个帖子的标题：", title)
			})
		}
		page++
		if page >= count{
			fmt.Println("--------------------------------数据拉取完成共"+strconv.Itoa(page)+"条------------------------------------")
			os.Exit(1)
		}
	}

}

/**
* 返回response
*/
func getResponse(url string) *http.Response {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	response, _ := client.Do(request)
	return response
}

/**
* 得到文章总数
*/
func getPageCount() int {
	response := getResponse(baseUrl)
	dom, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatalf("失败原因", response.StatusCode)
	}
	resDom := dom.Find(".text-center .pagination-sm li a")
	//len := resDom.Length()
	count,_ := strconv.Atoi(resDom.Eq(resDom.Length()-2).Text())
	return count
}

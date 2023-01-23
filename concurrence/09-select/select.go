package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func title(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return ch
}

func fastest(url1, url2, url3 string) string {
	c1 := title(url1)
	c2 := title(url2)
	c3 := title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "all responses received"
	default:
		return "no response"
	}

}

func main() {
	first := fastest(
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.cod3r.com.br",
	)

	fmt.Println(first)
}

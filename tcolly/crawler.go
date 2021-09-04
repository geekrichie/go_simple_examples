package main

import (
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)
func main() {
	url := "https://weibo.com/u/5821585980?is_hot=1"
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
	c.OnHTML("video", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		c.Visit(link)
	})
	c.OnResponse(func(response *colly.Response) {
		req := response.Request.URL.Path
		log.Println(req,response.StatusCode)
		lastIndex := strings.LastIndex(req, "/")
		filename := req[lastIndex+1:]
		var f ,_ = os.OpenFile(filename, os.O_CREATE| os.O_RDWR, 0 )
		if response.Headers.Get("Content-Type") == "video/mp4" {
               f.Write(response.Body)
               f.Close()
		}
		response.Save("demo.html")
	})
	c.Visit(url)
}
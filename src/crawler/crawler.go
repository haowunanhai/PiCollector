package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
	
	"common/log"
	"golang.org/x/net/html"

	_ "env"
	"util/timerecorder"
)

func main() {
	tr := timerecorder.New("Main")
	tr.Mark("start")

	defer log.Shutdown()
	defer log.Debug("FRAMEWORK", "elapsedTime", tr)
	defer tr.End()

	breadFirst(crawler, os.Args[1:])
}

const (
	downloadPath = "../../download/"
)

func crawler(url string) []string {
	tr := timerecorder.New("Crawler")
	tr.Mark("start")

	defer log.Debug("Crawler", "elapsedTime", tr)
	defer tr.End()

	resp, err := http.Get(url)
	if err != nil {
		log.Error("Crawler", "get url fail", url)
		return nil
	}
	tr.Mark("Get URL")

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Error("Crawler", os.Stderr, "%v\n", err)
		return nil
	}
	tr.Mark("Parse HTML")

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key != "href" {
						continue
					}
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}

					links = append(links, link.String())
				}
			} else if n.Data == "img" {
				for _, a := range n.Attr {
					if a.Key == "src" {
						imgLink, err := resp.Request.URL.Parse(a.Val)
						if err != nil {
							continue
						}
						
						tokens := strings.Split(a.Val, ".")
						if len(tokens) == 0{
							log.Error("Crawler", "img name has no suffix")
							continue
						}
						suffix := "." + tokens[len(tokens) - 1]
						
						imgName := downloadPath + strconv.FormatInt(time.Now().UnixNano(), 10) + suffix
						//读取img数据
						pic := ReadImgData(imgLink.String())

						err = ioutil.WriteFile(imgName, pic, 0644)
						if err != nil {
							log.Error("Crawler", "write file fail.", imgName)
						}
						fmt.Println("key:", a.Key, "value:", imgLink)
					}
				}
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	tr.Mark("Visit HTML")

	for _, link := range links {
		fmt.Println(link)
	}

	return links
}

func ReadImgData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Error("ReadImgData", "http get fail url: ", url)
		return nil
	}

	defer resp.Body.Close()

	pix, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("ReadImgData", "read resp body fail.")
		return nil
	}

	return pix
}
func breadFirst(f func(item string) []string, workList []string) {
	//排重
	urlMap := make(map[string]struct{})
	loopCnt := 2
	for loopCnt > 0 {
		items := workList
		workList = nil

		for _, item := range items {
			if _, ok := urlMap[item]; ok {
				continue
			}

			workList = append(workList, f(item)...)

			urlMap[item] = struct{}{}
		}
		
		loopCnt--
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

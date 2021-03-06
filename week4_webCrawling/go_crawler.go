package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

type result struct {
	link  string
	title string
	tag   string
	flag  bool
}

func main() {
	c := make(chan []result)

	results := []result{}

	page, err := strconv.Atoi(getPages())
	checkErr(err)

	for i := 1; i < page+1; i++ {
		go getContents(i+1, c)
	}
	for i := 0; i < page; i++ {
		r := <-c
		// ...을 할 경우 값들만 돌려 받게 된다.
		results = append(results, r...)
	}
	fmt.Println(len(results))

	// time.Sleep(time.Second * 10)
	cnt := 0
	var tmpStr []string
	for _, result := range results {
		if strings.Contains(result.title, "2019") {
			cnt++
			if result.flag == true {
				tmpStr = searchUserWant(result.link)
				if len(tmpStr) != 0 {
					fmt.Println(tmpStr)
				}

			}
			result.flag = false
		}
	}

	// i := 0
	// for i < cnt {
	// 	<-resultChan
	// }

}

func searchUserWant(link string) []string {
	fmt.Print("search ...")
	fmt.Println(link)
	res, err := http.Get(link)
	checkErr(err)
	checkStatusCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	result := []string{}

	doc.Find(".cont").Each(func(i int, contents *goquery.Selection) {
		// fmt.Println("OK, Find")
		result = append(result, inputUserWant(contents))
	})

	// fmt.Println("searchFin")
	// for i := 0; i < doc.Find(".cont").Length(); i++ {
	// 	result = append(result, r)
	// 	// fmt.Println(r)
	// }

	return result
}

func inputUserWant(cont *goquery.Selection) string {
	// for i := 0; i < cont.Find("span").Length(); i++ {
	// result, _ := iconv.ConvertString(cont.Find("span").Text(), "euc-kr", "utf-8")
	var results string
	cont.Children().Each(func(i int, contents *goquery.Selection) {
		if contents.Text() != "" {
			result, _ := iconv.ConvertString(contents.Text(), "euc-kr", "utf-8")
			results = results + result
		}

	})

	// cont.Find("span").Each(func(i int, contents *goquery.Selection) {
	// 	if contents.Text() != "" {
	// 		result, _ := iconv.ConvertString(contents.Text(), "euc-kr", "utf-8")
	// 		results = results + result
	// 	}
	// })

	// fmt.Println("fin")

	return results

}

func getPages() string {
	res, err := http.Get("https://security.cau.ac.kr/board.htm?bbsid=notice")
	checkErr(err)
	checkStatusCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// 마지막 값을 불러 오는 것
	page, _ := doc.Find(".paging>a").Last().Attr("href")

	var lastPage []string
	lastPage = strings.Split(page, "page=")

	return lastPage[1]
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getContents(page int, contentsChan chan []result) {
	resp, err := http.Get("https://security.cau.ac.kr/board.htm?bbsid=notice&ctg_cd=&skey=&keyword=&mode=list&page=" + strconv.Itoa(page))
	checkErr(err)
	checkStatusCode(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	c := make(chan result)

	var result []result
	// each : al을 찾고 각각 어떤 일을 할 것이다.
	doc.Find(".al").Each(func(i int, contents *goquery.Selection) {
		go inputValue(contents, c)
	})

	for i := 0; i < doc.Find(".al").Length(); i++ {
		r := <-c
		result = append(result, r)
	}
	// fmt.Println(page, "Done")
	contentsChan <- result
}

func checkStatusCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
}

func inputValue(cont *goquery.Selection, c chan<- result) {
	var result result
	result.link, _ = cont.Find("a").Attr("href")
	result.link = "https://security.cau.ac.kr/board.htm" + result.link
	result.title, _ = iconv.ConvertString(cont.Find("a").Text(), "euc-kr", "utf-8")
	result.flag = true
	fmt.Println(result.title)

	// 태그 분리
	if strings.Contains(result.title, "[ 학부 ]") {
		result.tag = "학부"
		result.title = strings.Replace(result.title, "[ 학부 ]", "", 1)
	} else {
		result.tag = "공통"
		result.title = strings.Replace(result.title, "[ 학부 ]", "", 1)
	}

	c <- result
}

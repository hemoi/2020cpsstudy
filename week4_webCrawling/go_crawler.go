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

	// var inputStr string
	cnt := 0
	resultChan := make(chan []string)

	fmt.Print("search ...")

	for _, result := range results {
		if strings.Contains(result.title, "2019") {
			cnt++
			go searchUserWant(result.link, resultChan)
		}
	}

	i := 0
	for i < cnt {
		<-resultChan
	}

}

func searchUserWant(link string, c chan []string) {
	fmt.Println(link)
	res, err := http.Get(link)
	checkErr(err)
	checkStatusCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	tmpChannel := make(chan []string)

	var result []string

	doc.Find(".cont").Each(func(i int, contents *goquery.Selection) {
		// fmt.Println("OK, Find")
		go inputUserWant(contents, tmpChannel)
	})

	// fmt.Println("searchFin")
	for i := 0; i < doc.Find(".cont").Length(); i++ {
		r := <-tmpChannel
		// result = append(result, r)
		fmt.Println(r)
		fmt.Println("")
	}
	c <- result
}

func inputUserWant(cont *goquery.Selection, c chan<- []string) {
	// for i := 0; i < cont.Find("span").Length(); i++ {
	// result, _ := iconv.ConvertString(cont.Find("span").Text(), "euc-kr", "utf-8")
	var results []string
	cont.Find("span").Each(func(i int, contents *goquery.Selection) {
		if contents.Text() != "" {
			result, _ := iconv.ConvertString(contents.Text(), "euc-kr", "utf-8")
			results = append(results, result)
		}
	})

	// fmt.Println("fin")
	if len(results) != 0 {
		c <- results
	}
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
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
}

func inputValue(cont *goquery.Selection, c chan<- result) {
	var result result
	result.link, _ = cont.Find("a").Attr("href")
	result.link = "https://security.cau.ac.kr/board.htm" + result.link
	result.title, _ = iconv.ConvertString(cont.Find("a").Text(), "euc-kr", "utf-8")

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

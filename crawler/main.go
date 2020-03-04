package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var BaseUrl string = "https://movie.douban.com/top250"

// 每页的条数
const pageSize int = 25

// 总共的页数
const pageTotal int = 10

type movie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

var movies []movie

type Page struct {
	Page int
	Url  string
}

// 获取分页
func getPages() []Page {
	// 每页25个, 10页
	// https://movie.douban.com/top250?start=0&filter=
	var pages []Page
	for i := 0; i <= 9; i++ {
		pages = append(pages, Page{
			Page: i,
			Url:  BaseUrl + fmt.Sprintf("?start=%d&filter=", i*pageSize),
		})
	}

	return pages
}

func parsePage(page Page) []movie {
	doc, err := goquery.NewDocument(page.Url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%#v \n", s)
		title := s.Find(".hd a span").Eq(0).Text()

		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, "  / ")

		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := s.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := s.Find(".quote .inq").Text()

		doubanmovie := movie{
			Title:    title,
			Subtitle: subtitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		log.Printf("i: %d, movie: %v", i, doubanmovie)

		movies = append(movies, doubanmovie)
	})

	return movies
}

var wg sync.WaitGroup

func main() {
	//pages := getPages()
	//for _, page := range pages {
	//	wg.Add(1)
	//	parsePage(page)
	//}
	//
	//fmt.Println(movies)

	resp, err := http.Get("https://www.baidu.com/")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Html())
	doc.Find("body > noscript").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("meta  http-equiv='refresh'")
		fmt.Printf("%#v", s)
	})
}

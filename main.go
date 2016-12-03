package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Qb struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("public/*")
	r.GET("/", Index)

	r.Run()
}

func Index(c *gin.Context) {
	p := c.DefaultQuery("page", "0")
	page, _ := strconv.Atoi(p)
	if page <= 0 {
		page = 1
	}

	result := getPage(page)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"items":    result,
		"title":    "糗百热门",
		"nextPage": page + 1,
		"prevPage": page - 1,
	})
	return
}

func getPage(p int) []Qb {

	qburl := "http://www.qiushibaike.com/hot/page/"
	qburl += strconv.Itoa(p)

	doc, err := goquery.NewDocument(qburl)
	if err != nil {
		log.Fatal(err)
	}

	var qb []Qb
	doc.Find("#content-left .article").Each(func(i int, s *goquery.Selection) {
		content := s.Find(".content span").Text()
		qb = append(qb, Qb{Id: i, Content: content})
	})

	return qb
}

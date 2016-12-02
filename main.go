package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	doc, err := goquery.NewDocument("http://www.qiushibaike.com")
	if err != nil {
		log.Fatal(err)
	}

	var qb []Qb

	doc.Find("#content-left .article").Each(func(i int, s *goquery.Selection) {
		content := s.Find(".content span").Text()
		qb = append(qb, Qb{Id: i, Content: content})
	})

	c.HTML(http.StatusOK, "index.html", gin.H{
		"items": qb,
		"title": "糗百热门",
	})
	return
}

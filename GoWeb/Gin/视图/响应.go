package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应JSON
func responseJSON(c *gin.Context) {
	type Msg struct {
		Name   string `json:"user"`
		Msg    string
		Number int
	}

	msg := Msg{"name", "你好", 2}
	c.JSON(http.StatusOK, msg)

}

// 响应XML
func responseXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "name", "message": "hey", "status": http.StatusOK})
}

// 响应YAML
func responseYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "name", "message": "hey", "status": http.StatusOK})
}

// 响应HTML
func responseHTML(c *gin.Context) {
	type Msg struct {
		Name   string `json:"user"`
		Msg    string
		Number int
	}

	msg := Msg{"Santonlee", "你好", 2}
	c.HTML(http.StatusOK, "index.html", msg)
}

// 重定向
func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/index.html") // 响应HTML文件的前置配置

	//响应文件
	router.StaticFile("/image.png", "static/image.png") // 响应单个文件
	router.StaticFS("/static", http.Dir("static/static"))
	router.GET("/json", responseJSON)
	router.GET("/xml", responseXML)
	router.GET("/yaml", responseYAML)
	router.GET("/html", responseHTML)
	router.GET("/baidu", redirect)

	router.Run(":8080")
}

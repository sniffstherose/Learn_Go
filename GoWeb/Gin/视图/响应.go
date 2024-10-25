package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseJSON(c *gin.Context) {
	type Msg struct {
		Name   string `json:"user"`
		Msg    string
		Number int
	}

	msg := Msg{"name", "你好", 2}
	c.JSON(http.StatusOK, msg)

}

func responseXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "name", "message": "hey", "status": http.StatusOK})
}

func responseYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "name", "message": "hey", "status": http.StatusOK})
}

func responseHTML(c *gin.Context) {
	type Msg struct {
		Name   string `json:"user"`
		Msg    string
		Number int
	}

	msg := Msg{"Santonlee", "你好", 2}
	c.HTML(http.StatusOK, "index.html", msg)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/*")
	router.GET("/json", responseJSON)
	router.GET("/xml", responseXML)
	router.GET("/yaml", responseYAML)
	router.GET("html", responseHTML)

	router.Run(":8080")
}

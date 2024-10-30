package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.GetQueryArray("user"))
	fmt.Println(c.GetQueryMap(("map")))
	fmt.Println(c.DefaultQuery("addr", "重庆市"))
}

func param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

func form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "重庆市"))
	forms, err := c.MultipartForm() // 接收所有form参数，包括文件
	fmt.Println(forms, err)
}

func main() {
	router := gin.Default()
	router.GET("/query", query)
	router.GET("/param/:user_id", param)
	router.GET("/param/:user_id/:book_id", param)
	router.POST("/form", form)
	router.Run(":80")
}

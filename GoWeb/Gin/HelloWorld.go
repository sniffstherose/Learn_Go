package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.String(http.StatusOK, "Hello Santonlee")
}

func main() {
	router := gin.Default()

	router.GET("/index", Index)

	// Run是对原生ListenAndServe的封装，所以也可用原生函数
	// router.Run(":8080")
	http.ListenAndServe(":8080", router)
}

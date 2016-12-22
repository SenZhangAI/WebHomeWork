package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	StartGin()
}

func StartGin() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)

	router.Run(":80")
}

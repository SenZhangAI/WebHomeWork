package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(200, "index.templ.html", gin.H{
		"testHeader": "test",
		"timestamp":  time.Now().Unix(),
	})
}

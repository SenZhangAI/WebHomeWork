package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.templ.html", gin.H{
		"timestamp": time.Now().Unix(),
	})
}

func userJoin(c *gin.Context) {
	c.HTML(http.StatusOK, "user_join.templ.html", gin.H{
		"timestamp": time.Now().Unix(),
	})
}

func usersList(c *gin.Context) {
	c.HTML(http.StatusOK, "user_list.templ.html", gin.H{
		"timestamp": time.Now().Unix(),
	})
}

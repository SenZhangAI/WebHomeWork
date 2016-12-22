package main

// RESTful style api
// See https://www.zhihu.com/question/27785028/answer/48096396
// Github RESTful api V3
// See https://developer.github.com/v3/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	HOST        string = "127.0.0.1"
	API_HOST    string = HOST + "/api"
	API_HOST_V1 string = HOST + "/api/v1"
)

func getAPI(c *gin.Context) {
	api_v1 := gin.H{
		"get_api  ": "GET  " + API_HOST,
		"get_user ": "GET  " + API_HOST_V1 + "/user",
		"post_user": "POST " + API_HOST_V1 + "/user",
	}

	c.IndentedJSON(http.StatusOK, api_v1)
}

func postUser(c *gin.Context) {
	var user User

	if c.Bind(&user) == nil {
		db.Save(&user)
	}

	c.Redirect(http.StatusMovedPermanently, "/user/login")
}

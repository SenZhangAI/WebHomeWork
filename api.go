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
		"get_users": "GET  " + API_HOST_V1 + "/users",
		"post_user": "POST " + API_HOST_V1 + "/user",
	}

	c.IndentedJSON(http.StatusOK, api_v1)
}

func postUser(c *gin.Context) {
	type User struct {
		ID       int32  `form:"user_id" json:"user_id"`
		UserName string `form:"user_name" json:"user_name" binding:"required"`
		UserNick string `form:"user_nick" json:"user_nick" binding:"required"`
		Password string `form:"user_password" json:"user_password" binding:"required"`
	}

	var user User

	if c.Bind(&user) == nil {
		db.Save(&user)
		c.Redirect(http.StatusMovedPermanently, "/user/login")
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
			"usage": "curl -i -d " +
				"\"user_name=XXX&user_nick=XXX&user_password=XXX\"" +
				API_HOST_V1 + "/user",
		})
	}
}

func getUsers(c *gin.Context) {
	type User struct {
		ID       int32
		UserName string
		UserNick string
	}

	var retData struct {
		Users []User
	}

	db.Select("id, user_name, user_nick").Find(&retData.Users)

	c.IndentedJSON(http.StatusOK, retData)
}

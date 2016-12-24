package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db            *gorm.DB
	sqlConnection string
)

func main() {
	ConnectDB()
	StartGin()
}

func StartGin() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")

	router.GET("/", login)
	router.GET("/user/login", login)
	router.GET("/user/join", userJoin)
	router.GET("/users", usersList)

	//RESTful API
	router.GET("/api", getAPI)
	router.POST("/api/v1/user", postUser)
	router.GET("/api/v1/users", getUsers)

	router.Run(":8080")
}

func ConnectDB() {
	var err error
	//username:password@tcp(host)/databasename?additional connection params
	sqlConnection = "root:zhs@tcp(127.0.0.1:3306)/web_server?parseTime=True"

	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
		return
	}

	//defer db.Close()

	if !db.HasTable("users") {
		panic("table [users] not found !")
	}
}

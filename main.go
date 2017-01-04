package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/go-oauth2/gin-server"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	aserver "gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"net/http"
)

var (
	db            *gorm.DB
	sqlConnection string
)

func main() {
	ConnectDB()
	OauthInit()
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

	//Oauth
	auth := router.Group("/oauth2")
	{
		auth.GET("/token", server.HandleTokenRequest)
	}

	//RESTful API
	router.GET("/api", getAPI)
	router.POST("/api/v1/user", postUser)
	authVerity := router.Group("/api/v1")
	{
		authVerity.Use(server.HandleTokenVerify())

		authVerity.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get("AccessToken")
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusOK, "not found")
		})

		authVerity.GET("/users", getUsers)
	}

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

func OauthInit() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	server.InitServer(manager)
	server.SetAllowGetAccessRequest(true)
	server.SetClientInfoHandler(aserver.ClientFormHandler)
	server.SetPasswordAuthorizationHandler(passwordAuthHandler)

}

func passwordAuthHandler(username, password string) (userID string, err error) {
	//TODO 对比用户名密码
	userID = "test"
	return
}

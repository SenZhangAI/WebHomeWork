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

	"golang.org/x/oauth2"
)

var (
	db            *gorm.DB
	sqlConnection string

	config = oauth2.Config{
		ClientID:     "000000",
		ClientSecret: "999999",
		Scopes:       []string{"all"},
		RedirectURL:  "",
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/oauth2/token",
		},
	}
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
	router.POST("/user/login", postLogin)
	router.GET("/user/join", userJoin)
	router.GET("/users", usersList)

	//Oauth
	auth := router.Group("/oauth2")
	{
		auth.GET("/token", tokenHandler)
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

func tokenHandler(c *gin.Context) {
	server.HandleTokenRequest(c)
	return
}

func passwordAuthHandler(username, password string) (userID string, err error) {
	//TODO 对比用户名密码
	userID = "test"
	return
}

func postLogin(c *gin.Context) {
	type User struct {
		UserName string `form:"user_name" json:"user_name" binding:"required"`
		Password string `form:"user_password" json:"user_password" binding:"required"`
	}
	var user User

	if c.Bind(&user) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad Request",
			"usage": "curl -i -d " +
				"\"user_name=XXX&user_password=XXX\"" +
				API_HOST_V1 + "/user/login",
		})
	}

	token, err := config.PasswordCredentialsToken(c, user.UserName, user.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, token)
}

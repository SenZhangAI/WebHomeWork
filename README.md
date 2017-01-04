# 设计要求

功能需求：

1. 用户注册，并填写用户基本信息，用户基本信息包括用户名、密码、昵称，用户名不得重复
2. 用户登录
3. 登陆之后可以查看所有已注册的用户列表，未登录状态不能查看用户列表

需要提供简单的Web展示页面，以及这三个功能对应的REST API，REST API测试可以通过curl命令或者chrome插件postman

在web页面中，记录用户登陆状态采用cookie的方式
在REST API里面，通过token来记录用户登陆状态，具体token的原理参考下oauth规范

Web框架采用gin，orm框架用gorm，数据库用mysql


## TODO
此项目暂时不修订更新，
以完成用户注册，和RESTful api

Oauth实现只进行了一半，
Oauth部分server和client分别试了两套框架，

server采用 gin-server框架，里面有相关参数的验证之类
通过获取username,password,client_id,client_secret来获取Access_token

可以测试：

```sh
$ curl -i "127.0.0.1:8080/oauth2/token?grant_type=password&username=sen&password=123&client_id=000000&client_secret=999999"
```

client端借用golang/x/oauth2中的PasswordCredentialsToken函数作为代理来实现以上提交功能，类似如下代码。
只是还有bug为搞定，所以返回null，bug可能是因为PasswordCredentialsToken采用POST方式提交，而gin-server框架验证中只能采用get方式，以后再考虑修改之

``` go
import (
	"golang.org/x/oauth2"
)

var (
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

func postLogin(c *gin.Context) {
//do some thing
	token, err := config.PasswordCredentialsToken(c, user.UserName, user.Password)
//do some thing
}
```

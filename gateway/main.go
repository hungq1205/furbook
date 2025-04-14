package main

import (
	"gateway/auth"
	"gateway/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var SecretKey = []byte("as you have seen, a very secret key")

type LoginOrSignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CheckUsernameRequest struct {
	Username string `json:"username"`
}

func main() {
	app := gin.New()
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func MakeGatewayHandler(app *gin.Engine) {
	group := app.Group("")
	group.Use(internal.AuthMiddleware())

	group.GET("/api/message/*path", ProxyTo("MESSAGE_SERVICE_URL"))
	group.GET("/api/post/*path", ProxyTo("POST_SERVICE_URL"))
	group.GET("/api/user/*path", ProxyTo("USER_SERVICE_URL"))
}

func MakeAuthHandler(app *gin.Engine, authRepo *auth.AuthRepository) *gin.Engine {
	group := app.Group("/api/auth")
	{
		group.POST("/login", func(c *gin.Context) {
			var body LoginOrSignUpRequest
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			val, err := authRepo.Authenticate(body.Username, body.Password)
			c.JSON(http.StatusOK, gin.H{"token": "token"})
		})

		group.POST("/signup", func(c *gin.Context) {
			var body LoginOrSignUpRequest
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": "token"})
		})
	}
}

func ProxyTo(targetHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tUrl, err := url.Parse(targetHost)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Request.URL.Host = tUrl.Host
		c.Request.URL.Path = tUrl.Path
		c.Request.Host = tUrl.Host
		c.Request.Header.Set("Referer", "http://gateway")

		proxy := httputil.NewSingleHostReverseProxy(tUrl)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

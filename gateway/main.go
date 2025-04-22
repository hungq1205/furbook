package main

import (
	"gateway/auth"
	"gateway/internal"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type LoginOrSignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := gin.New()
	MakeAuthHandler(app, MakeAuthRepository())
	MakeGatewayHandler(app)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func MakeAuthRepository() *auth.Repository {
	authDsn := "host=authdb user=postgres password=root dbname=auth port=5432 sslmode=disable"
	authRepo, err := auth.NewAuthRepository(authDsn)
	if err != nil {
		panic(err)
	}
	return authRepo
}

func MakeGatewayHandler(app *gin.Engine) {
	group := app.Group("")
	group.Use(internal.AuthMiddleware())

	group.GET("/api/message/*path", ProxyTo("MESSAGE_SERVICE_URL"))
	group.GET("/api/post/*path", ProxyTo("POST_SERVICE_URL"))
	group.GET("/api/user/*path", ProxyTo("USER_SERVICE_URL"))
}

func MakeAuthHandler(app *gin.Engine, authRepo *auth.Repository) {
	group := app.Group("/api/auth")
	{
		group.GET("/exists/:username", func(c *gin.Context) {
			user, err := authRepo.GetUser(c.Param("username"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"exists": user != nil})
		})

		group.POST("/login", func(c *gin.Context) {
			var body LoginOrSignUpRequest
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			val, err := authRepo.Authenticate(body.Username, body.Password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if !val {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
				return
			}
			token, err := internal.GenerateJwt(body.Username)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
		})

		group.POST("/signup", func(c *gin.Context) {
			var body LoginOrSignUpRequest
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			user, err := authRepo.GetUser(body.Username)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if user == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
				return
			}
			c.Status(http.StatusCreated)
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

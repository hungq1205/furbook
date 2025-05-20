package main

import (
	"gateway/auth"
	"gateway/client"
	"gateway/internal"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	messageServiceURL string = os.Getenv("MESSAGE_SERVICE_URL")
	postServiceURL    string = os.Getenv("POST_SERVICE_URL")
	userServiceURL    string = os.Getenv("USER_SERVICE_URL")
)

type LoginOrSignUpRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
}

func main() {
	app := gin.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	MakeAuthHandler(app, MakeAuthRepository(), MakeUserClient())
	MakeGatewayHandler(app)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func MakeUserClient() client.UserClient {
	return client.NewUserClient(userServiceURL)
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

	group.Any("/api/message/*path", ProxyTo(messageServiceURL))
	group.Any("/api/group/*path", ProxyTo(messageServiceURL))
	group.Any("/api/post/*path", ProxyTo(postServiceURL))
	group.Any("/api/user/*path", ProxyTo(userServiceURL))
}

func MakeAuthHandler(app *gin.Engine, authRepo *auth.Repository, userClient client.UserClient) {
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
			// val, err := authRepo.Authenticate(body.Username, body.Password)
			// if err != nil {
			// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// 	return
			// }
			// if !val {
			// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			// 	return
			// }
			token, err := internal.GenerateJwt(body.Username)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			user, err := userClient.GetUser(body.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
		})

		group.POST("/signup", func(c *gin.Context) {
			var body LoginOrSignUpRequest
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			euser, err := authRepo.GetUser(body.Username)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if euser == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
				return
			}
			user, err := userClient.CreateUser(body.Username, body.DisplayName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			err = authRepo.CreateUser(body.Username, body.Password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, gin.H{"user": user})
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
		c.Request.Host = tUrl.Host
		c.Request.Header.Set("Referer", "http://gateway")

		proxy := httputil.NewSingleHostReverseProxy(tUrl)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

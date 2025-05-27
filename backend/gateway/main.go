package main

import (
	"gateway/auth"
	"gateway/client"
	"gateway/internal"
	"gateway/websocket"
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

func main() {
	app := gin.New()

	app.RedirectTrailingSlash = false
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	auth.MakeAuthHandler(app, MakeAuthRepository(), MakeUserClient())
	websocket.MakeHandler(app, MakeGroupClient())
	MakeGatewayHandler(app)

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func MakeUserClient() client.UserClient {
	return client.NewUserClient(userServiceURL)
}

func MakeGroupClient() client.GroupClient {
	return client.NewGroupClient(messageServiceURL)
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

	group.Any("/api/message", ProxyTo(messageServiceURL))
	group.Any("/api/message/*path", ProxyTo(messageServiceURL))

	group.Any("/api/group", ProxyTo(messageServiceURL))
	group.Any("/api/group/*path", ProxyTo(messageServiceURL))

	group.Any("/api/post", ProxyTo(postServiceURL))
	group.Any("/api/post/*path", ProxyTo(postServiceURL))

	group.Any("/api/user", ProxyTo(userServiceURL))
	group.Any("/api/user/*path", ProxyTo(userServiceURL))
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

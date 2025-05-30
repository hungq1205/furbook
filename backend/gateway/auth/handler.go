package auth

import (
	"gateway/client"
	"gateway/internal"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginOrSignUpRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
}

func MakeAuthHandler(app *gin.Engine, authRepo *Repository, userClient client.UserClient) {
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
			if euser != nil {
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

		group.GET("/check", func(c *gin.Context) {
			authHeader := c.Request.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
				return
			}
			token := strings.TrimPrefix(authHeader, "Bearer ")
			username, err := internal.ParseJwt(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			user, err := userClient.GetUser(username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
		})
	}
}

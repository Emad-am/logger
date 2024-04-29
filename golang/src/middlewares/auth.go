package middlewares

import (
	usermodel "logger/src/models/user_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var password string
var userName string
var invalidCredential bool = false
var User usermodel.User

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isAuthenticated(c) {
			if invalidCredential {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("authUser", &User)
		c.Next()
	}
}

func GetAuthUser() *usermodel.User {
	return &User
}

func isAuthenticated(c *gin.Context) bool {
	if passwordHeaders, ok := c.Request.Header["Password"]; ok && len(passwordHeaders) > 0 {
		password = passwordHeaders[0]
	} else {
		password = ""
	}

	if userNamedHeaders, ok := c.Request.Header["User-Name"]; ok && len(userNamedHeaders) > 0 {
		userName = userNamedHeaders[0]
	} else {
		userName = ""
	}

	if userName == "" || password == "" {
		return false
	}

	var err error
	User, err = usermodel.FindByUserNameAndPassword(userName, password)

	if err != nil {
		invalidCredential = true
		return false
	}

	return true
}

package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user := c.Request.Header["userName"]
		// password := c.Request.Header["password"]

		// hashedPassword := tools.HashPassword(password)

		// before request

		c.Next()

		// after request

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

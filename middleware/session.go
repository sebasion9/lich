package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func Auth(c *gin.Context) {
	session := sessions.Default(c)
	idStr := session.Get("id")
	if idStr == nil {
		c.JSON(http.StatusUnauthorized, gin.H {
			"msg": "failed to authorize",
		})
		c.Abort()
		return
	}
	if _, ok := idStr.(uint); ok == true {
		c.Next()
		return
	}
	c.JSON(http.StatusBadRequest, gin.H {
		"msg": "invalid id type",
	})
	c.Abort()
	return
}


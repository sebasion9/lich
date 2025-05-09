package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func PathParamUint(param_names...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, name := range param_names {
			param := c.Param(name)
			id, err := strconv.ParseUint(param, 10, 0)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H {
					"msg" : "invalid id",
					"err" : err.Error(),
				})
				c.Abort()
				return
			}
			c.Set(name, uint(id))
		}
		c.Next()
	}
}

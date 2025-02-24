package resource

import (
	"errors"
	"lich/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// define whole CRUD for resource
// extend it with sync functionalities (pending?, fetch)

func New(dbOp func()) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resource model.Resource
		err := c.ShouldBindJSON(&resource)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}
	}
}

func GetAll(dbOp func() ([]model.Resource, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resources, err := dbOp()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "server error",
				"err" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resources)
	}
}



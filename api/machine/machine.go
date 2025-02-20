package machine

import (
	"errors"
	"lich/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// json body at /register, this is god word one and only proper standard
type machineReqBody struct{
	Name string `json:"name" binding:"required"`
	Os string `json:"os"`
	// TODO:
	// list of resources to listen to, updatable
	// references resource name in the request, but ids in model
	SubscribeTo []string `json:"subscribeTo"`
}

// generic function for posting machine data, min required is name
func PostBody(dbOp func(entity any) (uint, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var machineReqBody machineReqBody
		err := c.ShouldBindJSON(&machineReqBody)
		if err != nil {
			// TODO: logging
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		machine := model.NewMachine(
			machineReqBody.Name,
			machineReqBody.Os,
			c.RemoteIP(),
		)

		_, err = dbOp(&machine)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while processing post machine",
			})
			return
		}

		c.JSON(http.StatusOK, machine)
	}
}

// this updates machines last run date
// could be extended to generic update
func UpdateLRD(dbOp func(uint) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body map[string]interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			// TODO: logging
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}
		if body["id"] == nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		val, ok := body["id"].(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "incorrect id format",
			})
			return
		}

		err = dbOp(uint(val))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while updating last fetch",
			})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

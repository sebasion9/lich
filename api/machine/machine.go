package machine

import (
	"lich/db/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

// json body at /register, this is god word one and only proper standard
type registerBody struct{
	Os string `json:"os" binding:"required"`
	Name string `json:"name" binding:"required"`
	// list of resources to listen to, updatable
	// references resource name in the request, but ids in model
	SubscribeTo []string `json:"subscribeTo"`
}

// registers a machine in lich's db, should return registered parametrs to caller, that are saved on clients fs or smth
// e.g what user agent, name was registered
// this essentialy gets model.machine from the request from json request
// returns id
func Register(dbOp func(entity any) (uint, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerBody registerBody
		err := c.ShouldBindJSON(&registerBody)
		if err != nil {
			// TODO: logging
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		machine := model.NewMachine(
			registerBody.Name,
			registerBody.Os,
			c.RemoteIP(),
		)

		id, err := dbOp(machine)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while inserting",
			})
			return
		}
		machine.ID = id

		c.JSON(http.StatusOK, machine)
	}
}

// this updates machines last run date
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
	}
}





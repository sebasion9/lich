package machine

import (
	lich_time "lich/tool/time"
	"lich/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// json body at /register, this is god word one and only proper standard
type registerBody struct{
	Os string `json:"os" binding:"required"`
	Name string `json:"name" binding:"required"`
	// list of resources to listen to, updatable
	SubscribeTo []string `json:"subscribeTo"`
}


// registers a machine in lich's db, should return registered parametrs to caller, that are saved on clients fs or smth
// e.g what user agent, name was registered
// this essentialy gets model.machine from the request from json request
func Register(dbOp func(entity any) error) gin.HandlerFunc {
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
			lich_time.Now(),
			c.RemoteIP(),
			registerBody.Os,
		)

		err = dbOp(machine)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while inserting",
			})
			return
		}

		c.JSON(http.StatusOK, machine)
	}
}






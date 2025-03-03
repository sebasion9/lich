package sync

import (
	"errors"
	"lich/db/model"
	lich_db "lich/db/stmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SyncForce(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := c.MustGet("machine_id").(uint)
		resource_id := c.MustGet("resource_id").(uint)
		version_id := c.Param("version_id")


		var sub model.Subscription
		sub.MachineID = machine_id
		sub.ResourceID = resource_id
		ver, err := dbs.Sync.SyncOneVer(sub, version_id)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "internal server error",
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ver)
	}
}



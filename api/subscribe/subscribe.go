package subscribe

import (
	"errors"
	lich_db "lich/db/stmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Subscribe(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := c.MustGet("machine_id").(uint)
		resource_id := c.MustGet("resource_id").(uint)

		sub, err := dbs.Subscribe.Insert(machine_id, resource_id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg": "interval server error",
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, sub)
	}
}
func GetMult(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("id").(uint)
		byParam := c.Query("by")
		by := "machine"
		if byParam == "resource" {
			by = byParam
		}
		subs, err := dbs.Subscribe.GetById(id, by)
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

		c.JSON(http.StatusOK, subs)
	}
}


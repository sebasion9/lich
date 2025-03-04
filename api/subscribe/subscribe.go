package subscribe

import (
	"errors"
	lich_db "lich/db/stmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func Subscribe(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)
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
func GetOne(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)
		resource_id := c.MustGet("resource_id").(uint)

		sub, err := dbs.Subscribe.GetById(resource_id, machine_id)
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

		c.JSON(http.StatusOK, sub)
	}
}
func GetMult(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)

		sub, err := dbs.Subscribe.GetMult(machine_id)
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
		c.JSON(http.StatusOK, sub)
	}
}
func DeleteById(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)
		resource_id := c.MustGet("resource_id").(uint)

		rows, err := dbs.Subscribe.DeleteById(resource_id, machine_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "internal server error",
				"err" : err.Error(),
			})
			return
		}
		if rows == 0 {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

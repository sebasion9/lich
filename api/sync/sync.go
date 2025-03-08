package sync

import (
	"errors"
	lich_db "lich/db/stmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Sync(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)
		vers, err := dbs.Sync.Sub(machine_id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNoContent, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg": "internal server error",
				"err" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, vers)
		// get subs by mach id
		// for every sub, if sub.res.change_date > mach.last_sync
		// return []ver where ver = res[curr_ver]
	}
}


func SyncRes(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)
		resource_id := c.MustGet("resource_id").(uint)
		ver, err := dbs.Sync.ByResource(machine_id, resource_id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg": "internal server error",
				"err" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, ver)
	}
}

func SyncVer(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := sessions.Default(c).Get("id").(uint)
		ver_num := c.MustGet("version_num").(uint)
		resource_id := c.MustGet("resource_id").(uint)
		ver, err := dbs.Sync.ByVerNum(machine_id, resource_id, ver_num)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg": "internal server error",
				"err" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, ver)
	}
}

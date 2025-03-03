package version

import (
	"errors"
	"net/http"

	"lich/api"
	lich_db "lich/db/stmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this package will be responsible for retrieving data about versions, and for example reverting to some specific version

func GetVersions(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("resource_id").(uint)
		versions, err := dbs.Resource.GetVersionsById(uint(id))
		exit, status, res := api.QueryErr(err)
		if exit {
			c.JSON(status, res)
			return
		}

		c.JSON(http.StatusOK, versions)
	}
}

func New(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		resource_id := c.MustGet("resource_id").(uint)
		machine_id := sessions.Default(c).Get("id").(uint)


		var body map[string]any
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		if blob, ok:= body["blob"].(string); ok == true {
			version, err := dbs.Resource.NewVersion(resource_id, machine_id, blob)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, nil)
				return
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H {
					"msg" : "something went wrong while creating new version",
					"err": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, version)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H {
			"msg" : "bad body format",
		})
		return

	}
}

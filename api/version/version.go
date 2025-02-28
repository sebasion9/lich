package version

import (
	"errors"
	"net/http"
	"strconv"

	"lich/api"
	lich_db "lich/db/stmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this package will be responsible for retrieving data about versions, and for example reverting to some specific version

func GetVersions(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg": "invalid id format",
			})
			return
		}

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
		id, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg": "invalid id format",
			})
			return
		}

		var body map[string]any
		err = c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		if blob, ok:= body["blob"].(string); ok == true {
			version, err := dbs.Resource.NewVersion(uint(id), blob)
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

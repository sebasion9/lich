package version

import (
	"net/http"
	"strconv"

	lich_db "lich/db/stmt"
	"lich/api"
	"github.com/gin-gonic/gin"
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

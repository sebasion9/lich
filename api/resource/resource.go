package resource

import (
	"errors"
	"lich/api"
	"lich/db/model"
	lich_db "lich/db/stmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// define whole CRUD for resource
// extend it with sync functionalities (pending?, fetch)

func New(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resource model.Resource
		err := c.ShouldBindJSON(&resource)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}
		_, err = dbs.Machine.GetById(resource.MachineID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "machine with provided id doesn't exist",
			})
			return
		}

		res, err := dbs.Resource.Insert(resource)
		exit, status, response := api.InsertErr(err)
		if exit {
			c.JSON(status, response)
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func GetAll(dbOp func() ([]model.Resource, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resources, err := dbOp()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "server error",
				"err" : err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resources)
	}
}

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

func Edit(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.JSON(http.StatusOK, gin.H {
				"msg" : "invalid id format",
			})
			return
		}
	}
}


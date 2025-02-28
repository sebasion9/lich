package resource

import (
	"errors"
	"lich/api"
	"lich/db/model"
	lich_db "lich/db/stmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type newResource struct {
	Blob string `json:"blob"`
	Name string `json:"name"`
	Type string `json:"type"`
	MachineId uint `json:"machine_id"`
}

func New(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newResource newResource
		var resource model.Resource
		err := c.ShouldBindJSON(&newResource)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}
		resource.Name = newResource.Name
		resource.MachineID = newResource.MachineId
		resource.Type = newResource.Type

		resource.CurrentVersionID = 0
		resource.LastChangeAt = time.Now()

		blob := newResource.Blob


		_, err = dbs.Machine.GetById(resource.MachineID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "machine with provided id doesn't exist",
			})
			return
		}

		res, err := dbs.Resource.Insert(resource, blob)
		exit, status, response := api.InsertErr(err)
		if exit {
			c.JSON(status, response)
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
func GetById(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg": "invalid id format",
			})
			return
		}
		resource, err := dbs.Resource.GetById(uint(id))
		exit, status, res := api.QueryErr(err)
		if exit {
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusOK, resource)
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
func DeleteById(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// delete versions too
	}
}


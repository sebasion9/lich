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
		id := c.MustGet("id").(uint)
		resource, err := dbs.Resource.GetById(uint(id))
		exit, status, res := api.QueryErr(err)
		if exit {
			c.JSON(status, res)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}

func GetAll(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		resources, err := dbs.Resource.GetAllResource()
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
		id := c.MustGet("id").(uint)
		 
		rows, err := dbs.Resource.DeleteById(uint(id))
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


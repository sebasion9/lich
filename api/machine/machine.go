package machine

import (
	"errors"
	"lich/db/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WhoAmI(dbOp func(entity any) (uint, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var machine model.Machine
		err := c.ShouldBindJSON(&machine)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		machine.Ip = c.RemoteIP()

		var entity any = &machine
		if machine.Name == "" {
			machines := []model.Machine{machine}
			entity = &machines
		}

		_, err = dbOp(entity)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while processing post machine",
			})
			return
		}

		c.JSON(http.StatusOK, entity)
	}
}

func Register(dbOp func(entity any) (uint, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var machine model.Machine
		err := c.ShouldBindJSON(&machine)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad body format",
			})
			return
		}

		if machine.Name == "" || machine.Os == "" {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "fields: \"name\" and \"os\" are required",
			})
			return
		}

		machine.Ip = c.RemoteIP()


		_, err = dbOp(&machine)

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg": "machine with this name already exist, choose another one",
				"err": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while processing post machine",
			})
			return
		}

		c.JSON(http.StatusOK, machine)
	}

}


func UpdateLRD(dbOp func(uint) (int, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg" : "bad id format",
			})
			return
		}

		rowsAff, err := dbOp(uint(id))
		if rowsAff == 0 {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while updating last fetch",
			})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

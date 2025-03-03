package machine

import (
	"errors"
	"lich/db/model"
	lich_db "lich/db/stmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WhoAmI(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var machine model.Machine
		var entity any

		name := c.Param("name")
		machine_id := sessions.Default(c).Get("id")

		if machine_id, ok := machine_id.(uint); ok == true {
			machine.ID = machine_id
			entity = &machine
		} else if name != ":name" {
			machine.Name = name
			entity = &machine
		} else {
			machine.Ip = c.RemoteIP()
			entity = &[]model.Machine{machine}
		}

		_, err := dbs.Machine.GetOneOrMult(entity)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while processing post machine",
				"err" : err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, entity)
	}
}

func Register(dbs *lich_db.DbService) gin.HandlerFunc {
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
		machine.LastSync = time.Now()

		_, err = dbs.Machine.Insert(&machine)

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H {
				"msg": "machine with this name already exist, choose another one",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "something went wrong while processing post machine",
				"err": err.Error(),
			})
			return
		}

		session := sessions.Default(c)
		session.Set("id", machine.ID)
		session.Save()
		c.JSON(http.StatusOK, machine)
	}

}

func ActAs(dbs *lich_db.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		machine_id := c.MustGet("machine_id").(uint)
		_, err := dbs.Machine.GetById(machine_id)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"msg" : "internal server error",
				"err" : err.Error(),
			})
			return
		}

		session := sessions.Default(c)
		session.Set("id", machine_id)
		session.Save()
		c.JSON(http.StatusNoContent, nil)
	}
}


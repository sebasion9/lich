package main

// this is lich sync server
// a client part is not yet implemented, and definitiely not on this repository

import (
	api_machine "lich/api/machine"
	api_resource "lich/api/resource"
	api_version "lich/api/resource_version"
	lich_db "lich/db"
	"lich/db/model"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: config
	db, err := gorm.Open(sqlite.Open("lich.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open db %s\n", err.Error())
	}

	// migrate model to db
	db.AutoMigrate(
		&model.Resource{},
		&model.Machine{},
		&model.ResourceVersion{},
	)
	// database service
	dbs := lich_db.DbService {
		Db: db,
	}

	// http server
	r := gin.New()

	machine := r.Group("machine")
	{
		machine.POST("/register", api_machine.Register(dbs.Insert))
	}

	resource := r.Group("resource")
	{
		resource.GET("/dummy", api_resource.Dummy)

	}
	// version is ok now, because only one type of entity is planned to be versioned (resource)
	version := r.Group("version")
	{
		version.GET("/dummy", api_version.Dummy)
	}

	panic(r.Run())

}











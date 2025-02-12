package main

// this is lich sync server
// a client part is not yet implemented, and definitiely not on this repository

import (
	lich_db "lich/db"
	api_machine "lich/api/machine"
	api_resource "lich/api/resource"
	api_version "lich/api/resource_version"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: config
	addr := "localhost:1111"
	db, err := gorm.Open(sqlite.Open("lich.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open db %s\n", err.Error())
	}

	// migrate model to db
	db.AutoMigrate(
		&lich_db.Resource{},
		&lich_db.Machine{},
		&lich_db.ResourceVersion{},
	)


	r := gin.New()

	machine := r.Group("machine")
	{
		machine.GET("/register", api_machine.Register)

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

	panic(r.Run(addr))

}











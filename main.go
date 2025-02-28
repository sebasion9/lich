package main

// this is lich sync server
// a client part is not yet implemented, and definitiely not on this repository

import (
	api_machine "lich/api/machine"
	api_resource "lich/api/resource"
	api_version "lich/api/resource_version"
	lich_db "lich/db/stmt"
	"lich/db/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: config
	db, err := gorm.Open(sqlite.Open("lich.db"), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("Failed to open db %s\n", err.Error())
	}

	// migrate model to db
	db.AutoMigrate(
		&model.Resource{},
		&model.Machine{},
		&model.Version{},
		&model.Subscription{},
	)
	// database service
	dbs := lich_db.NewDb(db)

	// http server
	r := gin.New()

	machine := r.Group("machine")
	{
		machine.POST("/register", api_machine.Register(dbs.Machine.Insert))
		machine.POST("/whoami", api_machine.WhoAmI(dbs.Machine.GetOneOrMult))
	}

	resource := r.Group("resource")
	{
		resource.GET("/all", api_resource.GetAll(dbs.Resource.GetAllResource))
		resource.GET("/versions/:id", api_resource.GetVersions(&dbs))
		resource.POST("/new", api_resource.New(&dbs))
		resource.POST("/edit/:id", api_resource.Edit(&dbs))
	}
	// version is ok now, because only one type of entity is planned to be versioned (resource)
	version := r.Group("version")
	{
		version.GET("/dummy", api_version.Dummy)
	}

	subscribe := r.Group("subscribe")
	{
		subscribe.GET(":id", func(ctx *gin.Context) {})
	}

	panic(r.Run())

}











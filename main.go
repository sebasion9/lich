package main

import (
	api_machine "lich/api/machine"
	api_resource "lich/api/resource"
	api_subscribe "lich/api/subscribe"
	api_sync "lich/api/sync"
	api_version "lich/api/version"

	lich_db "lich/db/stmt"
	middleware "lich/middleware"

	"lich/db/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	// session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("lichsession", store))


	machine := r.Group("machine")
	{
		machine.PUT("/register", api_machine.Register(&dbs))
		machine.GET("/:name", api_machine.WhoAmI(&dbs))
		machine.GET("/actas/:machine_id", middleware.PathParamUint("machine_id"), api_machine.ActAs(&dbs))
	}

	resource := r.Group("resource")
	{
		resource.PUT("/new", middleware.Auth(), api_resource.New(&dbs))
		resource.GET("/all", api_resource.GetAll(&dbs))

		resource.GET(":id", middleware.PathParamUint("id"), api_resource.GetById(&dbs))
		resource.DELETE(":id", middleware.PathParamUint("id"), api_resource.DeleteById(&dbs))

		// version
		resource.PUT("/version/new-version/:id", middleware.PathParamUint("id"), api_version.New(&dbs))
		resource.GET("/version/:id", middleware.PathParamUint("id"), api_version.GetVersions(&dbs))
	}

	subscribe := r.Group("subscribe")
	{
		subscribe.PUT(":resource_id/:machine_id", middleware.PathParamUint("resource_id", "machine_id"), api_subscribe.Subscribe(&dbs))
		subscribe.GET(":id", middleware.PathParamUint("id"), api_subscribe.GetMult(&dbs))
	}
	sync := r.Group("sync")
	{
		// TODO:
		sync.GET(":machine_id/:resource_id/", middleware.PathParamUint("machine_id", "resource_id"), api_sync.SyncForce(&dbs))
		sync.GET(":machine_id/:resource_id/:version_id", middleware.PathParamUint("machine_id", "resource_id"), func(ctx *gin.Context) {})
	}

	panic(r.Run())

}











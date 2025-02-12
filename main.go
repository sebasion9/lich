package main

import (
	"fmt"
	"log"
	lich_db "lich/db"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// machine - a machine is needed to be identified, to save data of last fetch, timestamps, whos
// fetch changes - should be req every e.g every day, retrieves all changed entities >= last run date on this machine
// update resource - should be req when updating an resource (e.g a file), only last change is viable, but old versions are kept
// post resource - new resource to be posted
// revert resource by version
// get all resource versions

// stack
// sqlite3 with gorm
// some good logging lib
// gin/vanilla go

// sqlite3 -init init.sql lich.db .quit
// drop it, just use gorm

// define model for db
// write db stuff
// requests


func main() {
	// TODO: config
	fmt.Println("hello")
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


}











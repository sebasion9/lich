package db
// this is responsible for model definition

import (
	"time"
	"gorm.io/gorm"
)

type Machine struct {
	gorm.Model
	Name string
	LastFetch time.Time
	// identification
	Ip string
	UserAgent string
}

type Resource struct {
	gorm.Model
	Name string
	Path string
	MachineID uint
	Machine Machine
}


type ResourceVersion struct {
	gorm.Model
	ResourceID uint
	Resource Resource
	Num uint
	Date time.Time
}


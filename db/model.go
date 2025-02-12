package db
// this is responsible for model definition

import (
	"time"
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name string
	Path string
	MachineID uint
	Machine Machine
}

type Machine struct {
	gorm.Model
	Name string
	LastFetch time.Time
}

type ResourceVersion struct {
	gorm.Model
	ResourceID uint
	Resource Resource
	Num uint
	Date time.Time
}


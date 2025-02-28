package model
// this is responsible for model definition

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey;unique" json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Machine struct {
	Model
	// identification
	Name string `gorm:"unique" json:"name"`
	Ip string `json:"ip"`
	Os string `json:"os"`

	LastSync time.Time `json:"last_sync"`
}

type Resource struct {
	Model
	Name string `json:"name" binding:"required" gorm:"unique"`
	Type string `json:"type" binding:"required"`
	LastChangeAt time.Time `json:"last_change_at"`

	CurrentVersionID uint `json:"current_version_id"`

	MachineID uint `json:"machine_id" binding:"required"`
	Machine Machine `json:"machine"`
}


// Url?
// Blob?
type Version struct {
	Model
	Num uint `json:"num"`
	Blob string `json:"blob"`

	ResourceID uint `json:"resource_id"`
	Resource Resource `json:"resource"`
}

type Subscription struct {
	Model
	MachineID uint
	Machine Machine

	ResourceID uint
	Resource Resource
}



package model
// this is responsible for model definition

import (
	"time"
	lich_time "lich/tool/time"
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

	LastFetch time.Time `json:"last_fetch"`

	// TODO:
	// list of resources to listen to, updatable
	// references resource name in the request, but ids in model
	// SubscribeTo []string `json:"subscribeTo"`
}
// default constructor for db.Machine
func NewMachine(name string, os string, ip string) Machine {
	return Machine {
		Name: name,
		Ip: ip,
		Os: os,
		LastFetch: lich_time.Now(),
		Model: Model {
			CreatedAt: lich_time.Now(),
		},
	}
}

type Resource struct {
	Model
	Name string
	Path string
	MachineID uint
	Machine Machine
}


type ResourceVersion struct {
	Model
	ResourceID uint
	Resource Resource
	Num uint
	Date time.Time
}


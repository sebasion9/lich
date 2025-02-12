package model
// this is responsible for model definition

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Machine struct {
	Model
	Name string
	LastFetch time.Time
	// identification
	Ip string
	UserAgent string
}
// default constructor for db.Machine
func NewMachine(name string, lastFetch time.Time, ip string, userAgent string) Machine {
	return Machine {
		Name: name,
		LastFetch: lastFetch,
		Ip: ip,
		UserAgent: userAgent,
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


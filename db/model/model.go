package model
// this is responsible for model definition

import (
	"time"
	lich_time "lich/tool/time"
)

type Model struct {
	ID        uint `gorm:"primarykey;unique"`
	CreatedAt time.Time
}

type Machine struct {
	Model
	Name string
	LastFetch time.Time
	// identification
	Ip string
	UserAgent string
	// D:
	// SubscribeTo []uint
}
// default constructor for db.Machine
func NewMachine(name string, userAgent string, ip string) Machine {
	return Machine {
		Name: name,
		Ip: ip,
		UserAgent: userAgent,
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


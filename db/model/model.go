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

}

type Resource struct {
	Model
	Name string `json:"name" gorm:"unique"`
	Type string `json:"type"`
	LastChangeAt time.Time `json:"last_change_at"`

	CurrentVersionID uint `json:"current_version_id"`

	AuthorMachineID uint `json:"-"` 
	AuthorMachine Machine `json:"author_machine"`
}


type Version struct {
	Model
	Num uint `json:"num"`
	Blob string `json:"blob"`

	ResourceID uint `json:"-"`
	Resource Resource `json:"resource"`

	VersionAuthorID uint `json:"-"`
	VersionAuthor Machine `json:"version_author"`
}

type Subscription struct {
	Model
	MachineID uint `json:"-"`
	Machine Machine `json:"machine"`

	ResourceID uint `json:"-"`
	Resource Resource `json:"resource"`

	LastSync time.Time `json:"last_sync"`
}



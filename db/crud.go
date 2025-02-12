package db

// this is responsible for interfacing with sqlite

import (
	"errors"
	"lich/db/model"
	"gorm.io/gorm"
)


type DbService struct {
	Db *gorm.DB
}

func (dbs *DbService) Insert(entity any) error {
	err := errors.New("Invalid entity type")
	query := dbs.Db.Debug()
	switch val := entity.(type) {
	case model.Machine:
		err = query.Create(&val).Error
	case model.Resource:
		err = query.Create(&val).Error
	case model.ResourceVersion:
		err = query.Create(&val).Error
	}
	
	return err
}



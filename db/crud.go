package db

// this is responsible for interfacing with sqlite

import (
	"errors"
	"lich/db/model"
	lich_time "lich/tool/time"

	"gorm.io/gorm"
)


type DbService struct {
	Db *gorm.DB
}

func (dbs *DbService) Insert(entity any) (uint, error) {
	err := errors.New("Invalid entity type")
	var id uint
	query := dbs.Db.Debug()
	switch val := entity.(type) {
	case *model.Machine:
		err = query.Create(&val).Error
		id = val.ID
		entity = val
	case *model.Resource:
		err = query.Create(&val).Error
		id = val.ID
		entity = val
	case *model.ResourceVersion:
		err = query.Create(&val).Error
		id = val.ID
		entity = val
	}
	return id, err
}

func (dbs *DbService) Get(entity any) (uint, error) {
	err := errors.New("Invalid entity type")
	var id uint
	query := dbs.Db.Debug()
	switch val := entity.(type) {
		// machine is either []Machine by(ip) or Machine by(name),
	case *model.Machine:
		err = query.Where("name = ?", val.Name).First(&val).Error
		id = val.ID
		entity = val
	case *[]model.Machine:
		if len(*val) < 1 {
			err = gorm.ErrRecordNotFound
			break
		}
		first := (*val)[0]
		err = query.Where("ip = ?", first.Ip).Find(&val).Error
		id = first.ID
		entity = val
	case *model.Resource:
		err = query.First(&val).Error
		id = val.ID
		entity = val
	case *model.ResourceVersion:
		err = query.First(&val).Error
		id = val.ID
		entity = val
	}
	return id, err
}


func (dbs *DbService) UpdateLRD(id uint) (int, error) {
	now := lich_time.Now()
	result := dbs.Db.Model(&model.Machine{}).Where("id = ?", id).Update("last_fetch", now)
	return int(result.RowsAffected), result.Error
}


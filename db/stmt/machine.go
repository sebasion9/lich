package stmt
// this is responsible for interfacing with sqlite
// TODO: REFACTOR

import (
	"errors"
	"lich/db/model"
	"gorm.io/gorm"
)

type machineService struct { 
	*gorm.DB 
}

func (ms *machineService) Insert(entity any) (uint, error) {
	err := errors.New("Invalid entity type")
	var id uint
	query := ms.Debug()
	switch val := entity.(type) {
	case *model.Machine:
		err = query.Create(&val).Error
		id = val.ID
		entity = val
	}
	return id, err
}

func (ms *machineService) GetOneOrMult(entity any) (uint, error) {
	err := errors.New("Invalid entity type")
	var id uint
	query := ms.Debug()
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
	}
	return id, err
}

func (ms *machineService) GetById(id uint) (model.Machine, error) {
	var machine model.Machine
	machine.ID = id
	err := ms.First(&machine).Error
	return machine, err
}


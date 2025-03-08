package stmt

import (
	"lich/db/model"
	"time"

	"gorm.io/gorm"
)


type subService struct { 
	*gorm.DB
}


func (ss *subService) Insert(machine_id uint, resource_id uint) (model.Subscription, error) {
	var sub model.Subscription
	sub.MachineID = machine_id
	sub.ResourceID = resource_id
	sub.LastSync = time.UnixMicro(0)

	err := ss.Transaction(func(tx *gorm.DB) error {
		var machine model.Machine
		var resource model.Resource

		machine.ID = machine_id
		resource.ID = resource_id
		err := tx.First(&machine).Error
		if err != nil {
			return err
		}
		err = tx.First(&resource).Error

		if err != nil {
			return err
		}


		err = tx.Create(&sub).Error
		if err != nil {
			return err
		}

		err = tx.Preload("Resource.AuthorMachine").Preload("Machine").First(&sub).Error
		if err != nil {
			return err 
		}
		return nil
	})
	return sub, err
}


func (ss *subService) GetById(resource_id uint, machine_id uint) (model.Subscription, error) {
	var sub model.Subscription
	sub.MachineID = machine_id
	err := ss.
		Where("resource_id = ?", resource_id).
		Preload("Resource.AuthorMachine").
		Preload("Machine").
		First(&sub).Error
	return sub, err
}

func (ss *subService) GetMult(machine_id uint) ([]model.Subscription, error) {
	var sub []model.Subscription
	err := ss.
		Where("machine_id = ?", machine_id).
		Preload("Resource.AuthorMachine").
		Preload("Machine").
		Find(&sub).Error
	return sub, err
}

func (ss *subService) DeleteById(resource_id uint, machine_id uint) (int, error) {
	var rows int
	err := ss.Transaction(func(tx *gorm.DB) error {
		var sub model.Subscription
		res := tx.Where("resource_id = ?", resource_id).Where("machine_id = ?", machine_id).Delete(&sub)
		rows = int(res.RowsAffected)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	return rows, err
}


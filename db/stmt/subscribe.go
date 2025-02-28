package stmt

import (
	"lich/db/model"

	"gorm.io/gorm"
)


type subService struct { 
	*gorm.DB
}


func (ss *subService) Insert(machine_id uint, resource_id uint) (model.Subscription, error) {
	var sub model.Subscription
	sub.MachineID = machine_id
	sub.ResourceID = resource_id

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

		err = tx.Preload("Resource.Machine").Preload("Machine").First(&sub).Error
		if err != nil {
			return err 
		}
		return nil
	})
	return sub, err
}


func (ss *subService) GetById(id uint, by string) ([]model.Subscription, error) {
	var sub []model.Subscription
	var whereStmt string
	switch by {
	case "resource":
		whereStmt = "resource_id = ?"
	case "machine":
		whereStmt = "machine_id = ?"
	default:
		whereStmt = "resource_id = ?"
	}
	err := ss.Where(whereStmt, id).Preload("Resource.Machine").Preload("Machine").Find(&sub).Error
	return sub, err
}



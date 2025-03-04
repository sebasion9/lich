package stmt

import (
	"lich/db/model"
	"time"

	"gorm.io/gorm"
)



type syncService struct { 
	*gorm.DB
}

func (sy *syncService) ByResource(machine_id uint, resource_id uint) (model.Version, error){
	var ver model.Version
	var res model.Resource
	err := sy.Transaction(func(tx *gorm.DB) error {
		res.ID = resource_id
		err := tx.First(&res).Error
		if err != nil {
			return err
		}

		ver.ID = res.CurrentVersionID
		err = tx.Preload("Resource.AuthorMachine").Preload("VersionAuthor").First(&ver).Error
		if err != nil {
			return err
		}

		res := tx.Model(&model.Subscription{}).
		Where("resource_id = ?", resource_id).
		Where("machine_id = ?", machine_id).
		Update("LastSync", time.Now())
		if res.Error!= nil {
			return res.Error
		}

		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return err
	})
	return ver, err
}

func (sy *syncService) ByVerNum(machine_id uint, resource_id uint, ver_num uint) (model.Version, error) {
	var ver model.Version

	err := sy.Where("resource_id = ?", resource_id).
	Where("num = ?", ver_num).
	Preload("Resource.AuthorMachine").
	Preload("VersionAuthor").
	First(&ver).Error

	return ver, err
}

func (sy *syncService) Sub(machine_id uint) ([]model.Version, error) {
	var vers []model.Version
	var subs []model.Subscription
	err := sy.Transaction(func(tx *gorm.DB) error {

		return nil
	})
	return vers, err
}






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
	err := sy.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("resource_id = ?", resource_id).
		Where("num = ?", ver_num).
		Preload("Resource.AuthorMachine").
		Preload("VersionAuthor").
		First(&ver).Error

		if err != nil {
			return err
		}

		if ver.Resource.CurrentVersionID == ver.ID {
			res := tx.Model(&model.Subscription{}).
			Where("machine_id = ?", machine_id).
			Where("resource_id = ?", resource_id).
			Update("LastSync", time.Now())
			if res.Error!= nil {
				return res.Error
			}

			if res.RowsAffected == 0 {
				return gorm.ErrRecordNotFound
			}
		}

		return nil
	})


	return ver, err
}

func (sy *syncService) Sub(machine_id uint) ([]model.Version, error) {
	var vers []model.Version
	var subs []model.Subscription
	var ver_ids []int
	var sub_ids []int
	err := sy.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("machine_id = ?", machine_id).
		Preload("Machine").
		Preload("Resource.AuthorMachine").
		Find(&subs).Error
		if err != nil {
			return err
		}
		for _, sub := range subs {
			if sub.Resource.LastChangeAt.Unix() > sub.LastSync.Unix() {
				ver_ids = append(ver_ids, int(sub.Resource.CurrentVersionID))
				sub_ids = append(sub_ids, int(sub.ID))
			}
		}
		if len(ver_ids) < 1 {
			return gorm.ErrRecordNotFound
		}

		tx.Preload("Resource.AuthorMachine").
		Preload("VersionAuthor").
		Find(&vers, ver_ids)

		if err != nil {
			return err
		}
		tx.Model(&model.Subscription{}).Where("id in ?", sub_ids).Update("LastSync", time.Now())
		if err != nil {
			return err
		}

		return nil
	})
	return vers, err
}






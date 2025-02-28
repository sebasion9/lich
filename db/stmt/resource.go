package stmt

import (
	"lich/db/model"
	"time"

	"gorm.io/gorm"
)

type resourceService struct { 
	*gorm.DB
}

func (rs *resourceService) GetAllResource() ([]model.Resource, error) {
	var resources []model.Resource
	err := rs.Preload("Machine").Find(&resources).Error
	return resources, err
}

func (rs *resourceService) GetById(id uint) (model.Resource, error) {
	var resource model.Resource
	resource.ID = id
	err := rs.Preload("Machine").First(&resource).Error
	return resource, err
}
func (rs *resourceService) GetVersionsById(id uint) ([]model.Version, error) {
	var versions []model.Version
	err := rs.Where("resource_id = ?", id).
	Preload("Resource").
	Preload("Resource.Machine").
	Find(&versions).Error
	return versions, err
}


func (rs *resourceService) Insert(res model.Resource, blob string) (model.Resource, error) {
	ver := model.Version {
		Num : 0,
		Blob: blob,
	}

	err := rs.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&res).Error
		if err != nil {
			return err
		}
		ver.ResourceID = res.ID
		err = tx.Create(&ver).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Resource{}).Where("name = ?", res.Name).Update("CurrentVersionID", ver.ID).Error
		if err != nil {
			return err
		}

		err = tx.Preload("Machine").First(&res).Error
		if err != nil {
			return err
		}
		return nil
	})

	return res, err
}

func (rs *resourceService) DeleteById(id uint) (int, error) {
	var resource model.Resource
	resource.ID = id

	var rowsAffected int = 0
	err := rs.Transaction(func (tx *gorm.DB) error {
		if err := tx.Delete(&resource).Error; err != nil {
			return err
		}
		var versions []model.Version
		res := tx.Where("resource_id = ?", id).Delete(&versions)
		if res.Error != nil {
			return res.Error
		}
		rowsAffected = int(res.RowsAffected)
		return nil
	})
	return rowsAffected, err
}

func (rs *resourceService) NewVersion(id uint, blob string) (model.Version, error) {
	var version model.Version
	err := rs.Transaction(func(tx *gorm.DB) error {
		version.Blob = blob
		version.ResourceID = id

		var count int64 = 0
		err := tx.Model(&model.Version{}).Where("resource_id = ?", id).Count(&count).Error
		if err != nil {
			return err
		}

		version.Num = uint(count + 1)
		err = tx.Create(&version).Error
		if err != nil {
			return err
		}

		res := tx.Model(&model.Resource{}).
			Where("id = ?", id).
			Update("CurrentVersionID", version.ID).
			Update("LastChangeAt", time.Now())
		err = res.Error
		if err != nil {
			return err
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		err = tx.Preload("Resource").Preload("Resource.Machine").Find(&version).Error
		if err != nil {
			return err
		}

		return nil
	})

	return version, err
}




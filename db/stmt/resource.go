package stmt

import (
	"fmt"
	"lich/db/model"
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


func (rs *resourceService) Insert(res model.Resource) (model.Resource, error) {
	ver := model.Version {
		Num : 0,
		Url: fmt.Sprintf("%s@%d", res.Name, 0),
	}


	err := rs.Create(&res).Error
	if err != nil {
		return res, err
	}

	ver.ResourceID = res.ID
	ver.Current = true
	err = rs.Create(&ver).Error
	if err != nil {
		return res, err
	}


	err = rs.Preload("Machine").First(&res).Error

	return res, err
}



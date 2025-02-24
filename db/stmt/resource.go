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
	err := rs.Preload("Machine").Preload("Version").Find(&resources).Error
	return resources, err
}

func (rs *resourceService) Insert(res model.Resource) (model.Resource, error) {
	ver := model.Version {
		Num : 0,
		Url: fmt.Sprintf("%s@%d", res.Name, 0),
	}

	err := rs.Create(&ver).Error
	if err != nil {
		return res, err
	}

	res.VersionID = ver.ID
	err = rs.Create(&res).Error
	if err != nil {
		return res, err
	}

	err = rs.Preload("Machine").Preload("Version").First(&res).Error

	return res, err
}



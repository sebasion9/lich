package stmt

import (
	"lich/db/model"
	"gorm.io/gorm"
)

type resourceService struct { 
	*gorm.DB 
}

func (rs *resourceService) GetAllResource() ([]model.Resource, error) {
	var resources []model.Resource
	err := rs.Find(&resources).Error
	return resources, err
}

func (rs *resourceService) Insert() () {

}



package stmt

import "gorm.io/gorm"

type DbService struct {
	db *gorm.DB
	Machine machineService
	Resource resourceService
}
func NewDb(db *gorm.DB) DbService {
	return DbService {
		db : db,
		Machine : machineService{},
		Resource : resourceService{},
	}
}

package stmt

import "gorm.io/gorm"

type DbService struct {
	db *gorm.DB
	Machine machineService
	Resource resourceService
	Subscribe subService
}
func NewDb(db *gorm.DB) DbService {
	return DbService {
		db : db,
		Machine : machineService {db},
		Resource : resourceService {db},
		Subscribe : subService {db},
	}
}

package stmt

import (
	"fmt"
	"lich/db/model"

	"gorm.io/gorm"
)



type syncService struct { 
	*gorm.DB
}



// get machine and resource using sub
// 1. no version_id
// if machine.last_sync < resource.last_change
// fetch resource.version
// set machine.last_sync
// 2. if version_id
// fetch resource->version_id
// set machine.last_sync

// if machine and resource exist and version ahhgggg
func (sy *syncService) SyncOneVer(sub model.Subscription, version_id string) (model.Version, error) {
	err := sy.Transaction(func(tx *gorm.DB) error {
		err := tx.Preload("Machine").Preload("Resource.Machine").First(sub).Error
		if err != nil {
			return err 
		}

		if sub.Machine.LastSync.Unix() < sub.Resource.LastChangeAt.Unix() {
			// fetch R

		}
		return nil
	})
	fmt.Println(sub)
	return model.Version{}, err
}

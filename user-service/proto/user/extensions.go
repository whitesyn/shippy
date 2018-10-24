package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

// BeforeCreate assigns UUId as user Id
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.New()
	return scope.SetColumn("Id", id)
}

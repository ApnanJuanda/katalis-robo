package repository

import (
	"gorm.io/gorm"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
)

type GroupRepositoryImpl struct {
	DB *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &GroupRepositoryImpl{
		DB: db,
	}
}

func (repository GroupRepositoryImpl) Save(group *model.Group) {
	err := repository.DB.Create(group).Error
	helper.PanicIfError(err)
}

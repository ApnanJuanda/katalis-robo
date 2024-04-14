package repository

import (
	"gorm.io/gorm"
	"katalisRobo/component-store/helper"
	"katalisRobo/component-store/model"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repository CategoryRepositoryImpl) FindById(id string) *model.Category {
	var category model.Category
	err := repository.DB.First(&category, "id = ?", id).Error
	helper.PanicIfError(err)

	return &category
}

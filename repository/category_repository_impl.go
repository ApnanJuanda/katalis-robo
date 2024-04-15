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

func (repository CategoryRepositoryImpl) Save(category *model.Category) {
	err := repository.DB.Create(category).Error
	helper.PanicIfError(err)
}

func (repository CategoryRepositoryImpl) FindById(categoryId string) []*model.Group {
	var group []*model.Group
	err := repository.DB.Model(&group).Joins("Product").Joins("Category").First(&group, "groups.category_id = ?", categoryId).Error
	helper.PanicIfError(err)
	return group
}

func (repository CategoryRepositoryImpl) FindByCategoryId(categoryId string) *model.Category {
	var category model.Category
	err := repository.DB.First(&category, "id = ?", categoryId).Error
	if err != nil {
		return nil
	}
	return &category
}

func (repository CategoryRepositoryImpl) Update(category *model.Category) {
	err := repository.DB.Save(category).Error
	helper.PanicIfError(err)
}

func (repository CategoryRepositoryImpl) Delete(categoryId string) {
	err := repository.DB.Where("id = ?", categoryId).Delete(&model.Category{}).Error
	helper.PanicIfError(err)
}

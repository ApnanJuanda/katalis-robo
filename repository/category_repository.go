package repository

import "katalisRobo/component-store/model"

type CategoryRepository interface {
	Save(category *model.Category)
	FindById(categoryId string) []*model.Group
	FindByCategoryId(categoryId string) *model.Category
	Update(category *model.Category)
	Delete(categoryId string)
}

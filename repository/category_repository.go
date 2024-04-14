package repository

import "katalisRobo/component-store/model"

type CategoryRepository interface {
	FindById(id string) *model.Category
}

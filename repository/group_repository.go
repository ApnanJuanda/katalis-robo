package repository

import "katalisRobo/component-store/model"

type GroupRepository interface {
	Save(group *model.Group)
}

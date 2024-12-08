package repository

import "katalisRobo/component-store/model"

type ProductRepository interface {
	Save(product *model.Product)
	FindById(id string) *model.Product
	FindByMerchantId(merchantId string) []*model.Product
	Update(product *model.Product)
	Delete(id string)
}

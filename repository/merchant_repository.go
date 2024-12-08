package repository

import "katalisRobo/component-store/model"

type MerchantRepository interface {
	Save(merchant *model.Merchant)
	FindByEmail(email string) *model.Merchant
	FindAll() []*model.Merchant
	Update(merchant *model.Merchant)
	Delete(email string)
}

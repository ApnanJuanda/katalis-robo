package response

type ProductResponse struct {
	Name     string            `json:"name"`
	Merchant *MerchantResponse `json:"merchant"`
	Price    int               `json:"price"`
	Stock    int               `json:"stock"`
	ImageUrl string            `json:"imageUrl"`
}

package request

type ProductUpdateRequest struct {
	Name     string `validate:"max=200,min=1" json:"name"`
	Price    int    `validate:"gte=1" json:"price"`
	Stock    int    `validate:"gte=1" json:"stock"`
	ImageUrl string `json:"imageUrl"`
}

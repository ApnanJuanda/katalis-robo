package request

type ProductCreateRequest struct {
	Name       string `validate:"required,max=200,min=1" json:"name"`
	Price      int    `validate:"required,gte=1" json:"price"`
	Stock      int    `validate:"required,gte=1" json:"stock"`
	ImageUrl   string `validate:"required" json:"imageUrl"`
	CategoryId string `validate:"required" json:"categoryId"`
}

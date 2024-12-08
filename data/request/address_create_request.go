package request

type AddressCreateUpdateRequest struct {
	Detail string `validate:"required,max=200,min=1" json:"detail"`
}

package request

type CategoryCreateUpdateRequest struct {
	Name string `validate:"required,max=50,min=1" json:"name"`
}

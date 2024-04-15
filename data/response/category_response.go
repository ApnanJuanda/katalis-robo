package response

type CategoryResponse struct {
	Name     string             `json:"name"`
	Products []*ProductResponse `json:"products"`
}

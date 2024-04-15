package response

type AddressResponse struct {
	Customer *CustomerResponse `json:"customer"`
	Address  string            `json:"address"`
}

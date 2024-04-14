package request

type MerchantUpdateRequest struct {
	Name string `validate:"max=200,min=1" json:"name"`
	//Password      string `json:"password"`
	PhoneNumber   string `json:"phoneNumber"`
	AccountNumber string `json:"accountNumber"`
}

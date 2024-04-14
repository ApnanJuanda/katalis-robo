package request

type CustomerUpdateRequest struct {
	Name        string `validate:"max=200,min=1" json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

package request

type CustomerCreateRequest struct {
	Name        string `validate:"required,max=200,min=1" json:"name"`
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
	PhoneNumber string `validate:"required" json:"phoneNumber"`
}

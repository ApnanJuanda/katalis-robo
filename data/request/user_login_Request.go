package request

type UserLoginRequest struct {
	Username          string `validate:"required,max=200,min=1" json:"username"`
	EncryptedPassword string `validate:"required,max=200,min=1" json:"encryptedPassword"`
}

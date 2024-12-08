package response

type UserLoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiredIn   string `json:"expiredIn"`
	TokenType   string `json:"tokenType"`
}

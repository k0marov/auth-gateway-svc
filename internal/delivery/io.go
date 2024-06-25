package delivery

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokensResponse struct {
	AccessToken string `json:"access_token"`
}
